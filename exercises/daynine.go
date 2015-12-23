package exercises

// import (
// 	"strconv"
// 	"strings"

// 	"github.com/nottheeconomist/adventofcode_2015/puzzle"
// )

// func buildInput(s string) map[string]map[string]int {
// 	distances := make(map[string]map[string]int)
// 	for _, line := range strings.Split(s, "\n") {
// 		fields := strings.Fields(line)
// 		from, to, strDistance := fields[0], fields[2], fields[4]
// 		distance, _ := strconv.Atoi(strDistance) // TODO: error handling?
// 		// map from -> to
// 		submap, ok := distances[from]
// 		if !ok {
// 			distances[from] = make(map[string]int)
// 			submap = distances[from]
// 		}
// 		submap[to] = distance
// 		// map to -> from
// 		submap, ok = distances[to]
// 		if !ok {
// 			distances[to] = make(map[string]int)
// 			submap = distances[to]
// 		}
// 		submap[from] = distance
// 	}
// 	return distances
// }

// func DayNine(p *puzzle.Puzzle) {
// 	p.AddSolution("Not working")
// 	p.AddSolution("yet.....")
// 	return
// 	distances := buildInput(daynineinput)
// 	numCities := len(distances)
// 	cities := make([]string, 0, numCities)
// 	for city := range distances {
// 		cities = append(cities, city)
// 	}
// 	leastDistance := 10e100
// 	for start := range distances {
// 		routeDistances := 0
// 		toTravel := make(chan string, numCities)
// 		for _, city := range cities {
// 			if city != start {
// 				// we've already been to our first city
// 				toTravel <- city
// 			}
// 		}
// 	}
// 	leastDistance
// 	routeDistance
// }

// const daynineinput = `AlphaCentauri to Snowdin = 66
// AlphaCentauri to Tambi = 28
// AlphaCentauri to Faerun = 60
// AlphaCentauri to Norrath = 34
// AlphaCentauri to Straylight = 34
// AlphaCentauri to Tristram = 3
// AlphaCentauri to Arbre = 108
// Snowdin to Tambi = 22
// Snowdin to Faerun = 12
// Snowdin to Norrath = 91
// Snowdin to Straylight = 121
// Snowdin to Tristram = 111
// Snowdin to Arbre = 71
// Tambi to Faerun = 39
// Tambi to Norrath = 113
// Tambi to Straylight = 130
// Tambi to Tristram = 35
// Tambi to Arbre = 40
// Faerun to Norrath = 63
// Faerun to Straylight = 21
// Faerun to Tristram = 57
// Faerun to Arbre = 83
// Norrath to Straylight = 9
// Norrath to Tristram = 50
// Norrath to Arbre = 60
// Straylight to Tristram = 27
// Straylight to Arbre = 81
// Tristram to Arbre = 90`
