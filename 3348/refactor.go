package main

import (
	"fmt"
)

func walkThroughArena(spiderArena int, spiderSteps int, spiderStart int) int {
	return ((spiderStart - 1) + spiderSteps) % spiderArena
}

func playTest(inputCount int, steps int) int {
	spiderArena := inputCount * 2
	position := 0

	for i := true; i; i = spiderArena > inputCount {
		position = walkThroughArena(spiderArena, steps, position)
		if position >= inputCount {
			spiderArena = spiderArena - 1
		}

		if spiderArena == inputCount {
			return steps
		}

		if position < inputCount {
			return -1
		}
	}
	return -1
}

func nextValidSteps(currentSteps int, nSpiders int) int {
	if (currentSteps+1)%nSpiders >= nSpiders {
		return currentSteps + 1
	}
	return currentSteps + nSpiders + 1
}

func fight() {
	var n int
	fmt.Scan(&n)
	max := 2147483646

	for i := n + 1; i < max; i = nextValidSteps(i, n) {
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
