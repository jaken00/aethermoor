package world

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell struct {
	CellType     string
	CellEntities []*Entity
}

type WorldMap struct {
	Grid  [][]Cell
	X_len int
	Y_len int
}

func randomInt(generationMax int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(generationMax)
}

/*
func entityCellTypeGeneration(cellType string) map[string]string {

	return
}
*/

func entityGenerationPerCellCount(cellType string) int {

	//switch statement + random gen bounds for entites that spawn iwthin the certain entity zones
	switch cellType {
	case "PLAINS":
		return randomInt(3)
	case "WOODS":
		return randomInt(1)
	case "MOUNTAIN":
		return randomInt(4)
	case "RIVER":
		return randomInt(0)
	case "CAVE":
		return randomInt(4)
	case "GRASSLAND":
		return randomInt(4)
	}

	return 0
}

func (cell *Cell) initEntities(position Vec2, templates map[string]EntityTemplate) {
	numEntities := entityGenerationPerCellCount(cell.CellType)
	cell.CellEntities = make([]*Entity, numEntities)

	for i := 0; i < numEntities; i++ {
		rabbitTemplate := templates["rabbit"]
		entityID := fmt.Sprintf("rabbit_%d_%d_%d", position.XPos, position.YPos, i)
		cell.CellEntities[i] = SpawnEntityFromTemplate(rabbitTemplate, position, entityID)
	}
}

func getRandomCell() string {
	celltypes := [...]string{"PLAINS", "WOODS", "MOUNTAIN", "RIVER", "CAVE", "GRASSLAND"}

	selection := celltypes[rand.Intn(len(celltypes))]

	return selection
}

func (cell *Cell) populateCellType() {
	if cell.CellType != "TO_INIT" {
		return
	}
	cell.CellType = getRandomCell()

}

func GenerateWorld(x_length int, y_length int) *WorldMap {

	var worldMap WorldMap
	var currentPosition Vec2
	worldMap.X_len = x_length
	worldMap.Y_len = y_length
	grid := make([][]Cell, y_length)
	templates, _ := LoadTemplates("template.json")

	for i := 0; i < x_length; i++ {
		grid[i] = make([]Cell, x_length)
		for j := 0; j < y_length; j++ {

			grid[i][j] = Cell{
				CellType:     "TO_INIT", //IF TYPE == TO_INIT GENERATE
				CellEntities: nil,
			}
			grid[i][j].populateCellType()

			currentPosition.XPos = i
			currentPosition.YPos = j
			grid[i][j].initEntities(currentPosition, templates)

		}
	}
	worldMap.Grid = grid
	return &worldMap
}

func (worldMap *WorldMap) PrintWorldMap() {
	for i := 0; i < len(worldMap.Grid); i++ {
		for j := 0; j < len(worldMap.Grid[i]); j++ {
			fmt.Printf("%-12s", worldMap.Grid[i][j].CellType)
		}
		fmt.Println()
	}
}
