package utils

import "math"

const (
	// https://davidson.weizmann.ac.il/online/askexpert/%D7%90%D7%99%D7%9A-%D7%9E%D7%98%D7%99%D7%A1%D7%99%D7%9D-%D7%97%D7%9C%D7%9C%D7%99%D7%AA-%D7%9C%D7%99%D7%A8%D7%97
	MAIN_ENG_F   float64 = 430 // N
	SECOND_ENG_F float64 = 25  // N
)

func AccMax(weight float64) float64 {
	return acc(weight, true, 8)
}
func acc(weight float64, main bool, seconds int) float64 {
	var t float64 = 0

	if main {
		t += MAIN_ENG_F
	}

	t += float64(seconds) * SECOND_ENG_F
	ans := t / weight
	return ans
}

func ToRadians(degrees float64) float64 {
	return float64(degrees) * (math.Pi / 180.0)
}
