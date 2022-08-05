package weightconv

import "fmt"

type Pond float64
type Kilogram float64

const (
	// 1Pond = 0.45359237kg
	PondKilogram = 0.45359237
)

func (p Pond) String() string     { return fmt.Sprintf("%gPond", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gKilogram", k) }
