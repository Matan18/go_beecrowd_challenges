package main

import (
	"fmt"
)

func verifySpidersAlive(spiderArena []bool) int {
	count := 0
	for i := 0; i < len(spiderArena); i++ {
		if spiderArena[i] == true {
			count++
		}
	}
	return count
}

func createArena(size int) []bool {
	arena := make([]bool, size*2)
	for i := 0; i < int(size*2); i++ {
		arena[i] = true
	}

	return arena
}

func walkThroughArena(spiderArena []bool, spiderSteps int, spiderStart int) int {
	spidersAlive := verifySpidersAlive(spiderArena)
	return ((spiderStart - 1) + spiderSteps) % spidersAlive
}

func killspider(spiderArena []bool, spiderPosition int) []bool {
	return append(spiderArena[:spiderPosition], spiderArena[spiderPosition+1:]...)
}

func playTest(inputCount int, steps int) int {
	spiderArena := createArena(inputCount)
	position := 0

	for i := true; i; i = len(spiderArena) > inputCount {
		// print("arena 1", spiderArena)

		position = walkThroughArena(spiderArena, steps, position)
		// print("position", position)
		if position >= inputCount {
			spiderArena = killspider(spiderArena, position)
			// print("kill spider", position)
		}

		if len(spiderArena) == inputCount {
			return steps
		}

		if position < inputCount {
			return -1
		}
	}
	return -1
}

func fight() {
	var n int
	fmt.Scan(&n)

	for i := n + 1; i < 1000000000; i++ {
		// print("teste", i)
		res := playTest(n, i)
		if res != -1 {
			fmt.Printf("%v\n", res)
			break
		}
	}
}

func main() {
	fight()
}
