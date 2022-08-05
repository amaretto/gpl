package lengthconv

import "fmt"

type Metre float64
type Feet float64

const (
	FeetMetre = 0.3048
)

func (m Metre) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gfeet", f) }
