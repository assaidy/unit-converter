package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
)

var section string
var unitFrom, unitTo string
var value string
var amount float64
var units = map[string][]string{
	"length":      {"meters", "kilometers", "miles", "centimeters", "millimeters", "inches"},
	"weight":      {"grams", "kilograms", "pounds", "ounces", "milligrams"},
	"temperature": {"celsius", "fahrenheit", "kelvin"},
}

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("length", "weight", "temperature")...).
				Value(&section).
				Title("Section"),

			huh.NewSelect[string]().
				Value(&unitFrom).
				Height(8).
				Title("Convert from").
				OptionsFunc(func() []huh.Option[string] {
					opts := getSectionUniteFrom(section)
					return huh.NewOptions(opts...)
				}, &section),

			huh.NewSelect[string]().
				Value(&unitTo).
				Height(8).
				Title("Convert to").
				OptionsFunc(func() []huh.Option[string] {
					opts := getSectionUniteTo(section)
					return huh.NewOptions(opts...)
				}, &section),

			huh.NewInput().
				Title("Amount").
				Prompt("? ").
				Validate(isFloat64).
				Value(&value),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	// handling conversion
	res := handleConvert(amount, unitFrom, unitTo)

	fmt.Println("section:", section)
	fmt.Printf("%.3f %s = %.3f %s\n",
		amount,
		unitFrom,
		res,
		unitTo,
	)
}

func handleConvert(a float64, from, to string) float64 {
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
		"meters":      1,
		"kilometers":  1000,
		"miles":       1609.34,
		"centimeters": 0.01,
		"millimeters": 0.001,
		"inches":      0.0254,
	}

	meters := a * conversions[from]
	return meters / conversions[to]
}

func convertWeight(a float64, from, to string) float64 {
	conversions := map[string]float64{
		"grams":      1,
		"kilograms":  1000,
		"pounds":     453.592,
		"ounces":     28.3495,
		"milligrams": 0.001,
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

func isFloat64(input string) error {
	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	amount = num
	if err != nil {
		return fmt.Errorf("'%s' is not a valid number.", input)
	}
	return nil
}
