package pid

import (
	"math"

	"github.com/Avigdor-Kolonimus/BereshitSimulation/internal/utils"
)

type PID struct {
	saveVerticalSpeed float64
	deltaTime         int
	pparam            float64
	iparam            float64
	dparam            float64
	previousError     float64
	integral          float64
}

func NewPID(p, i, d, svs float64, dt int) *PID {
	pid := PID{saveVerticalSpeed: svs, deltaTime: dt, pparam: p, iparam: i, dparam: d, previousError: 0, integral: 0}
	return &pid
}
func (pid *PID) Compute(angle, altitude, verticalSpeed, weight float64) float64 {
	if altitude > 2000 {
		return pid.firstCompute(angle, verticalSpeed, weight)
	}
	if altitude < 5 { // no need to stop
		return 0.4
	}
	if altitude < 750 { // very close to the ground!
		return pid.thirdCompute(verticalSpeed)
	}

	return 1.0 - pid.secondCompute(altitude) // brake slowly, a proper PID controller here is needed!
}

func (pid *PID) thirdCompute(verticalSpeed float64) float64 {
	nn := 1.0              // maximum braking!
	if verticalSpeed < 5 { // if it is slow enough - go easy on the brakes
		nn = 0.7
	}
	return nn
}

func (pid *PID) secondCompute(altitude float64) float64 {
	pid.integral += altitude * float64(pid.deltaTime)
	nn := pid.pparam*altitude + pid.iparam*pid.integral + pid.dparam*(altitude-pid.previousError)/float64(pid.deltaTime)
	pid.previousError = altitude
	if nn > 1 {
		nn = 1
	}
	return nn
}
func (pid *PID) firstCompute(angle, verticalSpeed, weight float64) float64 {
	angleRad := utils.ToRadians(angle)
	nn := (verticalSpeed - pid.saveVerticalSpeed) / math.Cos(angleRad) * utils.AccMax(weight)
	if nn > 1 {
		return 1.0
	}
	if nn < 0 {
		return 0.0
	}
	return nn
}

// func (pid *PID) getIntegral() float64 {
// 	return pid.integral
// }

// func (pid *PID) setIntegral(integral float64) {
// 	pid.integral = integral
// }
// func (pid *PID) getSaveVerticalSpeed() float64 {
// 	return pid.saveVerticalSpeed
// }

// func (pid *PID) setSaveVerticalSpeed(verticalSpeed float64) {
// 	pid.saveVerticalSpeed = verticalSpeed
// }
