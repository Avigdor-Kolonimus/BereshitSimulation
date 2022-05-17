package main

import (
	"flag"
	"fmt"

	simulatorLanding "github.com/Avigdor-Kolonimus/BereshitSimulation/internal/simulator"
)

func main() {
	// register Int flag
	numAlgo := flag.Int("algorithm", 1, "Number of algorithm (0-BoazLanding 1-Landing 2-TwoPIDLanding)")
	var stringNameAlgo string

	// parse the flag
	flag.Parse()
	switch *numAlgo {
	case 0:
		stringNameAlgo = "BoazLanding algorithm"
	case 1:
		stringNameAlgo = "Landing algorithm"
	case 2:
		stringNameAlgo = "TwoPIDLanding algorithm"
	default:
		fmt.Println("Wrong number!")
		return
	}

	fmt.Println("You choose: ", stringNameAlgo)

	simulator := simulatorLanding.NewSimulator()
	fmt.Println("Simulating Bereshit's Landing:")
	simulator.Run(*numAlgo)
	fmt.Println("Finish")
}
