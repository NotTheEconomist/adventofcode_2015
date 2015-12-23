package exercises

import (
	"strconv"
	"strings"

	"github.com/nottheeconomist/adventofcode_2015/puzzle"
)

func lookAndSay(s string) string {
	var result []string
	count := 1
	num := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == num {
			count++
			continue
		}
		result = append(result, strconv.Itoa(count*10+int(num-'0')))
		count = 1
		num = s[i]
	}
	result = append(result, strconv.Itoa(count*10+int(num-'0'))) // grab the last one
	return strings.Join(result, "")
}

func DayTen(p *puzzle.Puzzle) {
	input := dayteninput
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}
	p.AddSolution(strconv.Itoa(len(input)))
	input2 := dayteninput
	for i := 0; i < 50; i++ {
		input2 = lookAndSay(input2)
	}
	p.AddSolution(strconv.Itoa(len(input2)))
}

const dayteninput = "1113222112"
