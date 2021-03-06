package tempconv

// FToC и CToF выполняют преобразование температур между двумя различными шкалами;
// они возвращают отличающиеся от исходных значения
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) // это преобразование типа Celsius в Fahrenheit
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC) // это преобразование типа Celsius в Kelvin
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC)) // это преобразование типа Kelvin в Celsius
}

func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit((k+Kelvin(AbsoluteZeroC))*9/5 + 32) // это преобразование типа Kelvin в Fahrenheit
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9) // это преобразование типа Fahrenheit в Celsius
}

func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 - Fahrenheit(AbsoluteZeroC)) // это преобразование типа Fahrenheit в Kelvin
}
