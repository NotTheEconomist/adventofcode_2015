package exercises

import (
	"strings"

	"github.com/nottheeconomist/adventofcode_2015/puzzle"
)

func wrapByte(b byte) (wrappedByte byte, didWrap bool) {
	wrappedByte = ((b - 'a') % ('z' + 1 - 'a')) + 'a'
	// fmt.Printf("b is %d, wrappedByte is %d\n", b, wrappedByte)
	if b > 'z' {
		didWrap = true
	}
	return
}

func filter(s string) bool {
	filterIOL := func(s string) bool {
		bad := "iol"
		return !strings.ContainsAny(s, bad)
	}
	filterPairs := func(s string) bool {
		count := 0
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				count++
				i++ // skip so we don't overlap
			}
		}
		return count >= 2
	}
	filterRunOfThree := func(s string) bool {
		for i := 0; i < len(s)-2; i++ {
			ba := []byte(s[i : i+3])

			if ba[0]+1 == ba[1] && ba[1]+1 == ba[2] {
				return true
			}
		}
		return false
	}
	funcs := []func(string) bool{filterIOL, filterPairs, filterRunOfThree}
	ok := true
	for _, f := range funcs {
		ok = ok && f(s)
	}
	return ok
}

func nextPass(s string) string {
	ba := []byte(s)
	for i := len(ba) - 1; i >= 0; i-- {
		b, didWrap := wrapByte(ba[i] + 1)
		ba[i] = b
		// fmt.Printf("ba has length %d, and ba[%d]==%c. New ba is %s\n", len(ba), i, ba[i], ba)
		if !didWrap {
			break
		}
	}
	return string(ba)
}

func DayEleven(p *puzzle.Puzzle) {
	newPass := dayeleveninput
	for {
		newPass = nextPass(newPass)
		// fmt.Printf("Testing \"%s\"\n", newPass)
		if filter(newPass) {
			break
		}
	}
	p.AddSolution(newPass)
	for {
		newPass = nextPass(newPass)
		// fmt.Printf("Testing \"%s\"\n", newPass)
		if filter(newPass) {
			break
		}
	}
	p.AddSolution(newPass)
}

const dayeleveninput = "hepxcrrq"
