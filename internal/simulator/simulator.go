package simulator

import (
	"encoding/csv"
	"fmt"
	"os"

	lunarLander "github.com/Avigdor-Kolonimus/BereshitSimulation/internal/bereshit"
	pidLander "github.com/Avigdor-Kolonimus/BereshitSimulation/internal/pid"
)

type Simulator struct {
	Time      int
	DeltaTime int
}

func NewSimulator() *Simulator {
	simulator := Simulator{Time: 0, DeltaTime: 1}
	return &simulator
}

func (simulator *Simulator) Run(algoNum int) {
	bereshit := lunarLander.NewBereshit()
	var fileName string

	switch algoNum {
	case 0:
		fileName = "BoazBereshitLanding.csv"
	case 1:
		fileName = "BereshitLanding.csv"
	case 2:
		fileName = "BereshitTwoPIDLanding.csv"
	}

	csvFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to open file", err)
		return
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	fmt.Println("Time, Vertical Speed, Horizontal Speed, Distance, Altitude, Angle, Weight, Acceleration Rate, Fuel")
	if err := csvWriter.Write([]string{"Time", "Vertical Speed", "Horizontal Speed", "Distance",
		"Altitude", "Angle", "Weight", "Acceleration Rate", "Fuel"}); err != nil {
		fmt.Println("Error writing row to file", err)
		return
	}
	pid := pidLander.NewPID(0.00045, 0.0000004, 0.00000155, 25, 1)
	for bereshit.Altitude > 0 && simulator.Time < 700 {
		row := bereshit.ToString(simulator.Time)
		if simulator.Time%10 == 0 || bereshit.Altitude < 100 { // to console
			fmt.Println(row)
		}
		// to csv
		if err := csvWriter.Write(row); err != nil {
			fmt.Println("Error writing row to file", err)
			return
		}
		switch algoNum {
		case 0:
			bereshit.BoazLanding()
		case 1:
			bereshit.Landing()
		case 2:
			bereshit.TwoPIDLanding(pid)
		}

		simulator.Time += simulator.DeltaTime
	}
	row := bereshit.ToStringFinish(simulator.Time)
	fmt.Println(row)

	// to csv
	if err := csvWriter.Write(row); err != nil {
		fmt.Println("Error writing row to file", err)
	}
}
