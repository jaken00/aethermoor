package world

//Now we need a function to call TickWorld in the Main function
func (worldMap *World) TickWorld() {
	// Collect all unique entities first to avoid processing them multiple times
	processedEntities := make(map[string]bool)
	entitiesToProcess := []*Entity{}

	for i := 0; i < worldMap.X_len; i++ {
		for j := 0; j < worldMap.Y_len; j++ {
			for _, entity := range worldMap.Grid[i][j].CellEntities {
				if !processedEntities[entity.Name] {
					processedEntities[entity.Name] = true
					entitiesToProcess = append(entitiesToProcess, entity)
				}
			}
		}
	}

	// Now process each entity exactly once
	for _, entity := range entitiesToProcess {
		if len(entity.Needs) == 0 {
			continue
			//Replenish Grass Function
		}

		entityDead := tickNeed(entity)

		if entityDead {
			Die(entity, worldMap)
			continue
		}

		if !entity.CheckCurrentCell(worldMap, getLowestNeedResource(entity)) {
			entity.MoveEntity(worldMap) //Move if we cant find out lowest need type at the current location
		}
	}

	worldMap.PrintWorldMap()

}
