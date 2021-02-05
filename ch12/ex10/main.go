package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
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
		if lex.text() == "t" {
			v.SetBool(true)
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
	fmt.Println(reflect.TypeOf(lex.token))
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

func Unmarshal(r io.Reader, out interface{}) (err error) {
	dec := NewDecoder(r)
	return dec.Decode(out)
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
type Float float64
type Interface interface{}
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
	case scanner.Float:
		f, _ := strconv.ParseFloat(dec.l.text(), 64)
		dec.l.next()
		return Float(f), nil
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
}

func main() {
	var m Movie
	Unmarshal(os.Stdin, &m)
	fmt.Println(m)
}
