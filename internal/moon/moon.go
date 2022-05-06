package moon

import (
	"math"
)

const (
	// from: https://he.wikipedia.org/wiki/%D7%94%D7%99%D7%A8%D7%97
	RADIUS   float64 = 3475 * 1000 // meters
	ACC      float64 = 1.622       // m/s^2
	EQ_SPEED float64 = 1700        // m/s
)

func GetAcc(speed float64) float64 {
	n := math.Abs(speed) / EQ_SPEED
	ans := (1 - n) * ACC
	return ans
}
