package main

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"github.com/tliddle1/wordle"
)

func TestMySolverFixture(t *testing.T) {
	gunit.Run(new(MySolverFixture), t)
}

type MySolverFixture struct {
	*gunit.Fixture
	mySolver  *MySolver
	evaluator *wordle.Evaluator
}

func (this *MySolverFixture) Setup() {
	this.mySolver = NewMySolver()
	this.evaluator = wordle.NewEvaluator()
}

func (this *MySolverFixture) TestSolver() {
	numGuesses, err := this.evaluator.PlayGame("smart", this.mySolver)
	this.So(err, should.BeNil)
	this.So(numGuesses, should.BeLessThanOrEqualTo, wordle.MaxNumGuesses)
}
