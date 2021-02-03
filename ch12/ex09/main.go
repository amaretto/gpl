package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

func Unmarshal(data []byte, out interface{}) (err error) {
	dec := NewDecoder(bytes.NewReader(data))
	return dec.Decode(out)
}

type Decoder struct {
	r   io.Reader
	buf []byte
	l   *lexer
}

func NewDecoder(r io.Reader) *Decoder {
	dec := &Decoder{r: r}
	dec.l = &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	dec.l.scan.Init(dec.r)
	dec.l.next()
	return dec
}

func (dec *Decoder) Decode(out interface{}) (err error) {
	read(dec.l, reflect.ValueOf(out).Elem())
	return nil
}

// Token API
type Token interface{}
type Symbol struct{ Name string }
type String string
type Int int
type StartList struct{} // (
type EndList struct{}   // )

func (dec *Decoder) Token() (Token, error) {
	switch dec.l.token {
	case scanner.EOF:
		return nil, io.EOF
	case scanner.Ident:
		name := dec.l.text()
		dec.l.next()
		return Symbol{Name: name}, nil
	case scanner.String:
		s, _ := strconv.Unquote(dec.l.text())
		dec.l.next()
		return String(s), nil
	case scanner.Int:
		i, _ := strconv.Atoi(dec.l.text())
		dec.l.next()
		return Int(i), nil
	case '(':
		dec.l.next()
		return StartList{}, nil
	case ')':
		dec.l.next()
		return EndList{}, nil
	}
	return nil, fmt.Errorf("unexpected token %q", dec.l.text())
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Satire          bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	Float           float64
	Complex         complex128
}

func main() {
	data := []byte(`
((Title "Dr. Strangelove")
 (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
 (Year 1964)
 (Actor (("Dr. Strangelove" "Peter Sellers")
         ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
         ("Pres. Merkin Muffley" "Peter Sellers")
         ("Gen. Buck Turgidson" "George C. Scott")
         ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
         ("Maj. T.J. \"King\" Kong" "Slim Pickens")))
 (Oscars ("Best Actor (Nomin.)"
          "Best Adapted Screenplay (Nomin.)"
          "Best Director (Nomin.)"
          "Best Picture (Nomin.)"))
`)
	dec := NewDecoder(bytes.NewReader(data))
	var stack []string
	var pos []int
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Tokenize failed : %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case StartList:
			pos = append(pos, len(stack))
		case EndList:
			if len(stack) > 1 {
				fmt.Printf("%s\n", strings.Join(stack, " "))
			}
			if pos[len(pos)-1] > 0 {
				stack = stack[:pos[len(pos)-1]]
			} else {
				stack = []string(nil)
			}
			pos = pos[:len(pos)-1]
		case Symbol:
			stack = append(stack, tok.Name)
		case String:
			stack = append(stack, string(tok))
		case Int:
			stack = append(stack, strconv.Itoa(int(tok)))
		}
	}

}
