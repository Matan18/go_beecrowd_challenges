package main

import (
	"bufio"
	"fmt"
	"os"
)

func calc(reader *bufio.Reader) {
	var count, id int
	myGames := 0
	fmt.Fscanf(reader, "%d %d\n", &count, &id)
	for i := 0; i < count; i++ {
		var user, gameType int
		fmt.Fscanf(reader, "%d %d\n", &user, &gameType)
		if user == id && gameType == 0 {
			myGames++
		}
	}
	fmt.Printf("%v\n", myGames)
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	for i := true; i == true; i = reader.Buffered() > 0 {
		calc(reader)
	}
}
