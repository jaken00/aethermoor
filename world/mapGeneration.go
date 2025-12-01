package world

import (
	"fmt"
	"math/rand"
)

const TerrainUnknown TerrainType = "TO_INIT"

type Cell struct {
	CellType     TerrainType
	CellEntities []*Entity
}

type World struct {
	Grid               [][]Cell
	X_len, Y_len       int
	ResouceTerrainDict *ResourceTerrainMapping
	Entities           map[string]*Entity
	CellEntities       map[Vec2][]string
}

func randomInt(generationMax int) int {
	if generationMax <= 0 {
		return 0
	}
	return rand.Intn(generationMax)
}

func entityCellTypeGeneration(cellType TerrainType) string {
	switch cellType {
	case TerrainPlains:
		return "rabbit"
	case TerrainWoods:
		return "wolf"
	case TerrainMountain:
		return "wolf"
	case TerrainRiver:
		return "grass"
	case TerrainCave:
		return "wolf"
	case TerrainGrassland:
		return "grass"
	default:
		return "none"
	}
}

func setTerrainResourceDictionary() *ResourceTerrainMapping {
	var mapping ResourceTerrainMapping
	t_dict := map[string][]string{
		"GRASS": {"RIVER", "GRASSLAND"},
		"MEAT":  {"PLAINS"},
	}
	mapping.ResourceDictionary = t_dict
	return &mapping
}

func (worldMap *World) GetTerrainResource(resource ResourceType) []string {
	if worldMap == nil || worldMap.ResouceTerrainDict == nil {
		return nil
	}
	return worldMap.ResouceTerrainDict.ResourceDictionary[string(resource)]
}

func entityGenerationPerCellCount(cellType TerrainType) int {

	switch cellType {
	case TerrainPlains:
		return randomInt(3)
	case TerrainWoods:
		return randomInt(1)
	case TerrainMountain:
		return randomInt(4)
	case TerrainRiver:
		return randomInt(4)
	case TerrainCave:
		return randomInt(4)
	case TerrainGrassland:
		return randomInt(4)
	default:
		return 0
	}
}

func (cell *Cell) initEntities(position Vec2, templates map[string]EntityTemplate) {
	numEntities := entityGenerationPerCellCount(cell.CellType)
	cell.CellEntities = make([]*Entity, 0, numEntities)

	entityTypeKey := entityCellTypeGeneration(cell.CellType)
	if entityTypeKey == "none" {
		return
	}

	for i := 0; i < numEntities; i++ {
		tmpl, ok := templates[entityTypeKey]
		if !ok {
			continue
		}
		entityID := fmt.Sprintf("%s_%d_%d_%d", entityTypeKey, position.XPos, position.YPos, i)
		newEntity := SpawnEntityFromTemplate(tmpl, position, entityID)
		cell.CellEntities = append(cell.CellEntities, newEntity)
	}
}

func getRandomCell() TerrainType {
	celltypes := []TerrainType{
		TerrainPlains, TerrainWoods, TerrainMountain,
		TerrainRiver, TerrainCave, TerrainGrassland,
	}
	selection := celltypes[rand.Intn(len(celltypes))]
	return selection
}

func (cell *Cell) populateCellType() {
	if cell.CellType != TerrainUnknown {
		return
	}
	cell.CellType = getRandomCell()

}

func GenerateWorld(x_length int, y_length int) *World {

	// seed RNG once here (deterministic runs can seed with a fixed value)
	rand.Seed(time.Now().UnixNano())

	var worldMap World
	var currentPosition Vec2
	worldMap.X_len = x_length
	worldMap.Y_len = y_length

	// allocate rows = x_length, columns = y_length (so grid[x][y])
	grid := make([][]Cell, x_length)

	templates, _ := LoadTemplates("template.json")
	worldMap.Entities = make(map[string]*Entity)
	worldMap.CellEntities = make(map[Vec2][]string)

	for i := 0; i < x_length; i++ {
		grid[i] = make([]Cell, y_length)
		for j := 0; j < y_length; j++ {
			// initialize default cell
			grid[i][j] = Cell{
				CellType:     TerrainUnknown,
				CellEntities: nil,
			}
			grid[i][j].populateCellType()

			currentPosition.XPos = i
			currentPosition.YPos = j
			grid[i][j].initEntities(currentPosition, templates)

			// register entities into world-level registries
			for _, e := range grid[i][j].CellEntities {
				if e == nil {
					continue
				}
				// prefer to keep the entity pointer in the global entity map
				worldMap.Entities[e.Name] = e
				pos := Vec2{XPos: i, YPos: j}
				worldMap.CellEntities[pos] = append(worldMap.CellEntities[pos], e.Name)
			}
		}
	}
	worldMap.Grid = grid
	worldMap.ResouceTerrainDict = setTerrainResourceDictionary()
	return &worldMap
}

func (worldMap *World) PrintWorldMap() {
	for i := 0; i < len(worldMap.Grid); i++ {
		for j := 0; j < len(worldMap.Grid[i]); j++ {
			fmt.Println("")
			fmt.Printf("%-12s", string(worldMap.Grid[i][j].CellType))
			entities := worldMap.Grid[i][j].CellEntities
			if len(entities) == 0 {
				continue
			}
			for k, e := range entities {
				if k > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%s", e.Name)
			}
		}
		fmt.Println()
	}
}
