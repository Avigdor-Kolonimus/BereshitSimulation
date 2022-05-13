package bereshit

import (
	"fmt"
	"math"

	"github.com/Avigdor-Kolonimus/BereshitSimulation/internal/moon"
)

const (
	WEIGHT_EMP  float64 = 165                      // kg
	WEIGHT_FULE float64 = 420                      // kg
	WEIGHT_FULL float64 = WEIGHT_EMP + WEIGHT_FULE // kg
	// https://davidson.weizmann.ac.il/online/askexpert/%D7%90%D7%99%D7%9A-%D7%9E%D7%98%D7%99%D7%A1%D7%99%D7%9D-%D7%97%D7%9C%D7%9C%D7%99%D7%AA-%D7%9C%D7%99%D7%A8%D7%97
	MAIN_ENG_F   float64 = 430   // N
	SECOND_ENG_F float64 = 25    // N
	MAIN_BURN    float64 = 0.15  //liter per sec, 12 liter per m'
	SECOND_BURN  float64 = 0.009 //liter per sec 0.6 liter per m'
	ALL_BURN     float64 = MAIN_BURN + 8*SECOND_BURN
	DeltaTime    float64 = 1 // sec
)

var flag = false

type Bereshit struct {
	VerticalSpeed    float64
	HorizontalSpeed  float64
	Distance         float64
	Angle            float64 // zero is vertical (as in landing)
	Altitude         float64
	AccelerationRate float64 // m/s^2
	Fuel             float64
	Weight           float64
	NN               float64 // rate[0,1]
}

func NewBereshit() Bereshit {
	bereshit := Bereshit{VerticalSpeed: 24.8, HorizontalSpeed: 932.0, Distance: 181 * 1000.0,
		Angle: 58.3, Altitude: 13748.0, AccelerationRate: 0.0,
		Fuel: 121.0, Weight: WEIGHT_EMP + 121, NN: 0.7}
	return bereshit
}

func (bereshit *Bereshit) ToString(time int) []string {
	row := []string{fmt.Sprintf("%d", time), fmt.Sprintf("%.3f", bereshit.VerticalSpeed), fmt.Sprintf("%.3f", bereshit.HorizontalSpeed),
		fmt.Sprintf("%.3f", bereshit.Distance), fmt.Sprintf("%.3f", bereshit.Altitude), fmt.Sprintf("%.3f", bereshit.Angle),
		fmt.Sprintf("%.3f", bereshit.Weight), fmt.Sprintf("%.3f", bereshit.AccelerationRate), fmt.Sprintf("%.3f", bereshit.Fuel)}
	return row
}

func (bereshit *Bereshit) ToStringFinish(time int) []string {
	row := []string{fmt.Sprintf("%d", time), "0.000", "0.000", fmt.Sprintf("%.3f", bereshit.Distance), "0.000",
		"0.000", fmt.Sprintf("%.3f", bereshit.Weight), "0.000", fmt.Sprintf("%.3f", bereshit.Fuel)}
	return row
}

func (bereshit *Bereshit) PID() (float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return bereshit.VerticalSpeed, bereshit.HorizontalSpeed, bereshit.Distance, bereshit.Altitude, bereshit.Angle, bereshit.Weight, bereshit.AccelerationRate, bereshit.Fuel, bereshit.NN
}

func (bereshit *Bereshit) BoazLanding() {
	// over 2 km above the ground
	if bereshit.Altitude > 2000 { // maintain a vertical speed of [20-25] m/s
		if bereshit.VerticalSpeed > 25 { // more power for braking
			bereshit.NN += 0.003 * DeltaTime
		}
		if bereshit.VerticalSpeed < 20 { // less power for braking
			bereshit.NN -= 0.003 * DeltaTime
		}
	} else { // lower than 2 km - horizontal speed should be close to zero
		if bereshit.Angle > 3 { // rotate to vertical position.
			bereshit.Angle -= 3
		} else {
			bereshit.Angle = 0
		}

		if bereshit.HorizontalSpeed < 2 {
			bereshit.HorizontalSpeed = 0
		}
		bereshit.NN = 0.5            // brake slowly, a proper PID controller here is needed!
		if bereshit.Altitude < 125 { // very close to the ground!
			bereshit.NN = 1.0               // maximum braking!
			if bereshit.VerticalSpeed < 5 { // if it is slow enough - go easy on the brakes
				bereshit.NN = 0.7
			}
		}
	}
	if bereshit.Altitude < 5 { // no need to stop
		bereshit.NN = 0.4
	}

	// main computations
	angle_rad := ToRadians(bereshit.Angle)
	horizontal_acc := math.Sin(angle_rad) * bereshit.AccelerationRate
	vertical_acc := math.Cos(angle_rad) * bereshit.AccelerationRate
	moon_acc := moon.GetAcc(bereshit.HorizontalSpeed)
	dw := DeltaTime * ALL_BURN * bereshit.NN

	if bereshit.Fuel > 0 {
		bereshit.Fuel -= dw
		bereshit.Weight = WEIGHT_EMP + bereshit.Fuel
		bereshit.AccelerationRate = bereshit.NN * accMax(bereshit.Weight)
	} else { // ran out of fuel
		bereshit.AccelerationRate = 0.0
	}

	vertical_acc -= moon_acc
	if bereshit.HorizontalSpeed > 0 {
		bereshit.HorizontalSpeed -= horizontal_acc * DeltaTime
	}
	bereshit.VerticalSpeed -= vertical_acc * DeltaTime
	bereshit.Distance -= bereshit.HorizontalSpeed * DeltaTime
	bereshit.Altitude -= bereshit.VerticalSpeed * DeltaTime
}

func (bereshit *Bereshit) Landing() {
	// over 2 km above the ground
	if bereshit.Altitude > 2000 {
		bereshit.NN = 1.0
		if bereshit.Altitude > 10000 && !flag {
			bereshit.Angle += 3
			if bereshit.Angle > 82 {
				bereshit.Angle = 82.3
				flag = true
			}
		} else {
			if bereshit.Altitude > 8000 {
				bereshit.Angle -= 2
				if bereshit.Angle < 58 {
					bereshit.Angle = 58.3
					flag = false
				}
			} else {
				bereshit.NN = 0.825
			}
		}
	} else { // lower than 2 km - horizontal speed should be close to zero
		if bereshit.Angle > 3 { // rotate to vertical position.
			bereshit.Angle -= 3
		} else {
			bereshit.Angle = 0
		}

		if bereshit.HorizontalSpeed < 2 {
			bereshit.HorizontalSpeed = 0
		}
		bereshit.NN = 0.68           // brake slowly, a proper PID controller here is needed!
		if bereshit.Altitude < 125 { // very close to the ground!
			bereshit.NN = 0.9                // maximum braking!
			if bereshit.VerticalSpeed < 10 { // if it is slow enough - go easy on the brakes
				bereshit.NN = 0.665
			}
		}
	}
	if bereshit.Altitude < 5 { // very close to the ground!
		bereshit.NN = 0.4
	}
	// main computations
	angle_rad := ToRadians(bereshit.Angle)
	horizontal_acc := math.Sin(angle_rad) * bereshit.AccelerationRate
	vertical_acc := math.Cos(angle_rad) * bereshit.AccelerationRate
	moon_acc := moon.GetAcc(bereshit.HorizontalSpeed)
	dw := DeltaTime * ALL_BURN * bereshit.NN

	if bereshit.Fuel > 0 {
		bereshit.Fuel -= dw
		bereshit.Weight = WEIGHT_EMP + bereshit.Fuel
		bereshit.AccelerationRate = bereshit.NN * accMax(bereshit.Weight)
	} else { // ran out of fuel
		bereshit.AccelerationRate = 0.0
	}

	vertical_acc -= moon_acc
	if bereshit.HorizontalSpeed > 0 {
		bereshit.HorizontalSpeed -= horizontal_acc * DeltaTime
	}
	bereshit.VerticalSpeed -= vertical_acc * DeltaTime
	bereshit.Distance -= bereshit.HorizontalSpeed * DeltaTime
	bereshit.Altitude -= bereshit.VerticalSpeed * DeltaTime
}

func accMax(weight float64) float64 {
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
