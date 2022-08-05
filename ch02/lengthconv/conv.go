package lengthconv

//convert Metre to Feet
func MToF(m Metre) Feet { return Feet(m / FeetMetre) }

//convert Feet to Metre
func FToM(f Feet) Metre { return Metre(f * FeetMetre) }
