package main

import (
	"aethermoor/world"
)

func main() {
	worldmap := world.GenerateWorld(10, 10)
	worldmap.PrintWorldMap()

}
