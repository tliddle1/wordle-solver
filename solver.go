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
	solver := MySolver{}
	solver.setData()
	return &solver
}

func (this *MySolver) Debug() bool {
	return false
}

func (this *MySolver) Guess(turnHistory []wordle.Turn) string {
	// TODO hard code first guess (once you figure out the best one)
	//len(turnHistory) == 0 {
	//	return ""
	//}

	this.trimValidTargets(turnHistory)
	return this.calculateWordWithBestExpectedInfo()
}

func (this *MySolver) Reset() {
	this.setData()
}

func (this *MySolver) setData() {
	this.validTargets = data.ValidTargets
	this.validGuesses = append(data.ValidGuesses, data.ValidTargets...)
}

func (this *MySolver) trimValidTargets(turnHistory []wordle.Turn) {
	if len(turnHistory) == 0 {
		return
	}
	lastTurn := turnHistory[len(turnHistory)-1]
	var newTargets []string
	for _, target := range this.validTargets {
		if this.isPossibleTarget(target, lastTurn.Guess, lastTurn.Pattern) {
			newTargets = append(newTargets, target)
		}
	}
	this.validTargets = newTargets
}

func (this *MySolver) isPossibleTarget(target string, guess string, pattern wordle.Pattern) bool {
	return wordle.CheckGuess(target, guess) == pattern
}

func (this *MySolver) calculateWordWithBestExpectedInfo() string {
	if len(this.validTargets) == 1 {
		return this.validTargets[0]
	}

	var mostExpectedInfo float64
	var bestWord string

	for _, possibleGuess := range this.validGuesses {
		expectedInfo := this.expectedInfo(possibleGuess)
		if expectedInfo > mostExpectedInfo {
			mostExpectedInfo = expectedInfo
			bestWord = possibleGuess
		}
	}
	return bestWord
}

func (this *MySolver) expectedInfo(possibleGuess string) float64 {
	// TODO implement me!
	panic("implement expectedInfo")
}
