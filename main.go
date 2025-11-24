package main

import (
	"aethermoor/world"
	"fmt"
)

func main() {
	worldmap := world.GenerateWorld(10, 10) //10 by 10 for now, later we will make this dynamic
	worldmap.PrintWorldMap()

	templates, err := world.LoadTemplates("template.json")
	if err != nil {
		panic(err)
	}

	rabbit, ok := templates["rabbit"]
	if !ok {
		panic("rabbit template not found")
	}

	fmt.Printf("Loaded rabbit template: %+v\n", rabbit)

}
