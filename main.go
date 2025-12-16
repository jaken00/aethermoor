package main

import (
	"aethermoor/world"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	worldmap := world.GenerateWorld(10, 10)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Input: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "q" {
			break
		} else if input == "n" {
			worldmap.TickWorld() //ticks world forward
		}
	}

}
