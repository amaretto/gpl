package tempconv

//convert Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//convert Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//convert Celsius to Kervin
func CToK(c Celsius) Kervin { return Kervin(c + 273.15) }

//convert Fahrenheit to Kervin
func FtoK(f Fahrenheit) Kervin { return Kervin(FToC(f) + 273.15) }

//convert Kervin to Celsius
func KToC(k Kervin) Celsius { return Celsius(k - 273.15) }

//convert Kervin to Celsius
func KToF(k Kervin) Fahrenheit { return Fahrenheit(CToF(Celsius(k - 273.15))) }
