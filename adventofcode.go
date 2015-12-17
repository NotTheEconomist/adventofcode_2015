package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/nottheeconomist/adventofcode_2015/exercises"
	"github.com/nottheeconomist/adventofcode_2015/puzzle"
)

const puzzleCount = 7 // how many puzzles do we have?
const barWidth = 60   // barWidth of the loading bar

var wg sync.WaitGroup // waits for all exercises to complete before going on

// initPuzzle delegates to NewPuzzle and additionally manages a list of puzzles.
func initPuzzle(day int, f func(*puzzle.Puzzle), puzzleList *[]*puzzle.Puzzle) *puzzle.Puzzle {
	wg.Add(1)
	p := puzzle.Puzzle{Day: day}
	*puzzleList = append(*puzzleList, &p)
	go func() {
		f(&p)
		wg.Done()
	}()
	return &p
}

// loading will print a loading bar until the quit channel is closed.
func loading(width int, quit chan bool) {
	ascending := true
	for i := 0; ; {
		bars := strings.Repeat("|", i)
		dashes := strings.Repeat("-", width-i)
		fmt.Print("\r", bars, dashes)
		if i >= width {
			ascending = false
		} else if i <= 0 {
			ascending = true
		}
		if ascending {
			i++
		} else {
			i--
		}
		time.Sleep(50 * time.Millisecond)
		select {
		case _, ok := <-quit:
			if ok == false {
				// channel closed, done loading!
				return
			}
			panic("Something was sent to quit")
		default:
			continue
		}
	}
}

func main() {
	puzzles := make([]*puzzle.Puzzle, 0, puzzleCount)

	// TODO: build puzzles
	initPuzzle(1, exercises.DayOne, &puzzles)
	initPuzzle(2, exercises.DayTwo, &puzzles)
	initPuzzle(3, exercises.DayThree, &puzzles)
	initPuzzle(4, exercises.DayFour, &puzzles)
	initPuzzle(5, exercises.DayFive, &puzzles)
	initPuzzle(6, exercises.DaySix, &puzzles)
	initPuzzle(7, exercises.DaySeven, &puzzles)

	quit := make(chan bool)
	//go loading(barWidth, quit)

	wg.Wait()
	close(quit)                                          // stop loading
	fmt.Print("\r", strings.Repeat(" ", barWidth), "\r") // blank out loading bar
	for _, p := range puzzles {
		fmt.Println(p)
	}
}
