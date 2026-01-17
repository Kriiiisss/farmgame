package main

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
