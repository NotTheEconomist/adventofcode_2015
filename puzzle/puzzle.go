package puzzle

import (
	"fmt"
	"strings"
)

// Puzzle represents a day of the advent of code calendar.
type Puzzle struct {
	Day       int
	solutions []string
}

// String stringifies the Puzzle results
func (p *Puzzle) String() string {
	return fmt.Sprintf("Day %v: %s", p.Day, strings.Join(p.solutions, ", "))
}

// AddSolution adds a solution to the Puzzle results
func (p *Puzzle) AddSolution(sol string) {
	p.solutions = append(p.solutions, sol)
}
