package converter

var fromUnit, toUnit string
var units = map[string][]string{
	"length":      {"meter", "kilometer", "mile", "centimeter", "millimeter", "inch", "foot", "yard"},
	"weight":      {"gram", "kilogram", "pound", "ounce", "milligram"},
	"temperature": {"celsius", "fahrenheit", "kelvin"},
}

// convert the amount from unit to another unit in the same section
func Convert(section, fUnit, tUnit string, amount float64) float64 {
	return handleConvert(amount, fUnit, tUnit, section)
}

func handleConvert(a float64, from, to, section string) float64 {
	if from == to {
		return a
	}

	switch section {
	case "length":
		return convertLength(a, from, to)
	case "weight":
		return convertWeight(a, from, to)
	case "temperature":
		return convertTemperature(a, from, to)
	default:
		return 0
	}
}

func convertLength(a float64, from, to string) float64 {
	conversions := map[string]float64{
		"meter":      1,
		"kilometer":  1000,
		"mile":       1609.34,
		"centimeter": 0.01,
		"millimeter": 0.001,
		"inch":      0.0254,
		"yard":       0.9144,
		"foot":       0.3048,
	}

	meters := a * conversions[from]
	return meters / conversions[to]
}

func convertWeight(a float64, from, to string) float64 {
	conversions := map[string]float64{
		"gram":      1,
		"kilogram":  1000,
		"pound":     453.592,
		"ounce":     28.3495,
		"milligram": 0.001,
	}

	grams := a * conversions[from]
	return grams / conversions[to]
}

func convertTemperature(a float64, from, to string) float64 {
	switch from {
	case "celsius":
		if to == "fahrenheit" {
			return (a * 9 / 5) + 32
		} else if to == "kelvin" {
			return a + 273.15
		}
	case "fahrenheit":
		if to == "celsius" {
			return (a - 32) * 5 / 9
		} else if to == "kelvin" {
			return (a-32)*5/9 + 273.15
		}
	case "kelvin":
		if to == "celsius" {
			return a - 273.15
		} else if to == "fahrenheit" {
			return (a-273.15)*9/5 + 32
		}
	}
	return 0
}

func getSectionUniteFrom(sec string) []string {
	return units[sec]
}

func getSectionUniteTo(sec string) []string {
	return units[sec]
}
