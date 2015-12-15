package exercises

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/nottheeconomist/adventofcode_2015/puzzle"
)

func makeHash(secret string) []byte {
	hash := md5.New()
	io.WriteString(hash, secret)
	return hash.Sum(nil)
}

func checkHash(hash []byte, prefix string) bool {
	strHash := fmt.Sprintf("%x", hash)
	return strings.HasPrefix(strHash, prefix)
}

// DayFour runs the calculations for the Day Four puzzle
func DayFour(p *puzzle.Puzzle) {
	var solutionOne, solutionTwo string
	for i := 0; ; i++ {
		strI := strconv.Itoa(i)
		hash := makeHash(dayfourinput + strI)
		if solutionOne == "" && checkHash(hash, "00000") {
			solutionOne = strI
		}
		if solutionTwo == "" && checkHash(hash, "000000") {
			solutionTwo = strI
		}
		if solutionOne != "" && solutionTwo != "" {
			break
		}
	}
	p.AddSolution(solutionOne)
	p.AddSolution(solutionTwo)
}

const dayfourinput = `ckczppom`
