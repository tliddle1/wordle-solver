package main

import (
	"fmt"

	"github.com/tliddle1/wordle"
)

func main() {
	evaluator := wordle.NewEvaluator()
	mySolver := NewMySolver()
	avgNumGuesses, err := evaluator.EvaluateSolver(mySolver)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(avgNumGuesses)
}
