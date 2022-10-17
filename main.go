package main

import (
	"experiments/counter"
	"experiments/fizzbuzz"
	"experiments/gameofpig"
	"experiments/mcstat"
	"experiments/random"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected an experiment to run!")
		fmt.Println("Experiments: calculator, gameofpig, counter, fizzbuzz, mcstat")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "calculator":
		random.Calculator()
	case "gameofpig":
		gameofpig.Play()
	case "counter":
		counter.Counter()
	case "fizzbuzz":
		fizzbuzz.Fizzbuzz()
	case "mcstat":
		mcstat.Mcstat()
	default:
		fmt.Println("Invalid!")
	}
}
