package world

import "fmt"

func tickNeeds(e *Entity) {
	//e.tickNeed("food")
	//e.tickNeed("shelter")
	//tick more as needed
	fmt.Print("TICK!")

}

func (worldMap *World) TickWorld() {

	for i := 0; i < worldMap.X_len; i++ {
		for j := 0; j < worldMap.Y_len; j++ {
			currentGrid := worldMap.Grid[i][j]

			if len(currentGrid.CellEntities) == 0 {
				continue
			} else {
				for _, entity := range currentGrid.CellEntities {
					tickNeeds(entity) //already a pointer
					//entity.ActOrchestrator()

				}
			}
		}
	}

}
