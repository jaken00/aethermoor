package world

import (
	"fmt"
	"math/rand"
	"strings"
)

const TerrainUnknown TerrainType = "TO_INIT"

type Cell struct {
	CellType     TerrainType
	CellEntities []*Entity
}

// Might need to do a global counter to increment entities or do a UUID
type World struct {
	Grid               [][]Cell
	X_len, Y_len       int
	ResouceTerrainDict *ResourceTerrainMapping
	Entities           map[string]*Entity
	CellEntities       map[Vec2][]string //maybe instead of returning a string we return the actual entity? or at the least e.name?
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
	// Print separator and header
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("=== WORLD MAP ===")
	fmt.Println("Terrain: P=Plains, W=Woods, M=Mountain, R=River, C=Cave, G=Grassland")
	fmt.Println("Entities: r=rabbit, w=wolf, g=grass (numbers show count if >1)")
	fmt.Println()

	// Print column numbers header
	fmt.Print("   ")
	for j := 0; j < worldMap.Y_len; j++ {
		if j < 10 {
			fmt.Printf(" %d", j)
		} else {
			fmt.Printf("%d", j)
		}
	}
	fmt.Println()

	// Print grid with row numbers
	for i := 0; i < worldMap.X_len; i++ {
		// Row number
		if i < 10 {
			fmt.Printf("%d  ", i)
		} else {
			fmt.Printf("%d ", i)
		}

		for j := 0; j < worldMap.Y_len; j++ {
			cell := worldMap.Grid[i][j]

			// Get terrain symbol
			terrainSym := getTerrainSymbol(cell.CellType)

			// Build entity string
			entityStr := ""
			if len(cell.CellEntities) > 0 {
				// Count entities by type
				entityCounts := make(map[string]int)
				for _, e := range cell.CellEntities {
					if e != nil {
						// Extract entity type from name (e.g., "rabbit_0_0_1" -> "rabbit")
						parts := strings.Split(e.Name, "_")
						if len(parts) > 0 {
							entityType := parts[0]
							entityCounts[entityType]++
						}
					}
				}

				// Build compact entity representation
				var entityParts []string
				if count := entityCounts["rabbit"]; count > 0 {
					if count == 1 {
						entityParts = append(entityParts, "r")
					} else {
						entityParts = append(entityParts, fmt.Sprintf("r%d", count))
					}
				}
				if count := entityCounts["wolf"]; count > 0 {
					if count == 1 {
						entityParts = append(entityParts, "w")
					} else {
						entityParts = append(entityParts, fmt.Sprintf("w%d", count))
					}
				}
				if count := entityCounts["grass"]; count > 0 {
					if count == 1 {
						entityParts = append(entityParts, "g")
					} else {
						entityParts = append(entityParts, fmt.Sprintf("g%d", count))
					}
				}
				entityStr = strings.Join(entityParts, "")
			}

			// Print cell: [TerrainSymbol][Entities] (padded to 4 chars)
			if entityStr != "" {
				fmt.Printf("%s%s", terrainSym, entityStr)
				// Pad to 4 characters for alignment
				totalLen := len(terrainSym) + len(entityStr)
				for totalLen < 4 {
					fmt.Print(" ")
					totalLen++
				}
			} else {
				fmt.Printf("%s   ", terrainSym)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func getTerrainSymbol(terrain TerrainType) string {
	switch terrain {
	case TerrainPlains:
		return "P"
	case TerrainWoods:
		return "W"
	case TerrainMountain:
		return "M"
	case TerrainRiver:
		return "R"
	case TerrainCave:
		return "C"
	case TerrainGrassland:
		return "G"
	default:
		return "?"
	}
}

// TEMP PRINT FUNCTION
func (worldMap *World) PrintEntityStatus() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("=== ENTITY STATUS ===")
	fmt.Println(strings.Repeat("=", 80))

	for i := 0; i < worldMap.X_len; i++ {
		for j := 0; j < worldMap.Y_len; j++ {
			cell := worldMap.Grid[i][j]

			if len(cell.CellEntities) == 0 {
				continue
			}

			for _, entity := range cell.CellEntities {
				fmt.Printf("\nEntity: %s | Position: (%d, %d) | Alive: %t\n",
					entity.Name, entity.Position.XPos, entity.Position.YPos, entity.Alive)

				fmt.Println("  Produces:")
				for _, prod := range entity.Produces {
					fmt.Printf("    %s: %.2f/%.2f (regen: %.2f)\n",
						prod.Type, prod.Current, prod.Max, prod.RegenRate)
				}

				fmt.Println("  Needs:")
				for needType, need := range entity.Needs {
					fmt.Printf("    %s (%s): %.2f/%.2f (threshold: %.2f, consume: %.2f)\n",
						needType, need.Resource, need.Current, need.Max, need.Threshold, need.ConsumeRate)
				}
			}
		}
	}
	fmt.Println(strings.Repeat("=", 80))
}
