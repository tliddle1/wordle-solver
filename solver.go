package main

import (
	"github.com/tliddle1/wordle"
	"github.com/tliddle1/wordle/data"
)

type MySolver struct {
	validTargets []string
	validGuesses []string
}

func NewMySolver() *MySolver {
	mySolver := MySolver{
		validTargets: make([]string, len(data.ValidTargets)),
		validGuesses: make([]string, len(data.ValidGuesses)+len(data.ValidTargets)),
	}
	copy(mySolver.validTargets, data.ValidTargets)
	mySolver.validGuesses = append(data.ValidGuesses, data.ValidTargets...)
	return &mySolver
}

func (this *MySolver) Debug() bool {
	return false
}

func (this *MySolver) Guess(turnHistory []wordle.Turn) string {
	//TODO implement me
	panic("implement me")
}

func (this *MySolver) Reset() {
	this.validTargets = data.ValidTargets
	this.validGuesses = data.ValidGuesses
}
