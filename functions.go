package main

import (
	"unicode/utf8"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func IsInRange(number int, min int, max int) bool {
	return (number >= min && number <= max)
}

func Clamp(number float32, min float32, max float32) float32 {
	if number < min {
		return min
	}
	if number > max {
		return max
	}
	return number
}

func Floor(number float32) int {
	return int(number)
}

func Ceil(number float32) int {
	return int(number) + 1
}

func HandleTextInput(text *string) {
	character := rl.GetCharPressed()
	if character != 0 {
		*text += string(character)
	}
	if rl.GetKeyPressed() == rl.KeyBackspace {
		_, lastRuneLength := utf8.DecodeLastRuneInString(*text)
		*text = (*text)[:len(*text)-lastRuneLength]
	}
}
