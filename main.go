package main

import (
	"fmt"

	simulatorLanding "github.com/Avigdor-Kolonimus/BereshitSimulation/internal/simulator"
)

func main() {
	simulator := simulatorLanding.NewSimulator()
	fmt.Println("Simulating Bereshit's Landing:")
	simulator.Run()
	fmt.Println("Finish")
}
