package weightconv

//convert Kilogram to Pond
func KToP(k Kilogram) Pond { return Pond(k / PondKilogram) }

//convert Pond to Kilogram
func PToK(p Pond) Kilogram { return Kilogram(p * PondKilogram) }
