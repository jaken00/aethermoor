package world

//Now we need a function to call TickWorld in the Main function
func (worldMap *World) TickWorld() {

	for i := 0; i < worldMap.X_len; i++ {
		for j := 0; j < worldMap.Y_len; j++ {
			currentGrid := worldMap.Grid[i][j]

			if len(currentGrid.CellEntities) == 0 {
				continue
			} else {
				for _, entity := range currentGrid.CellEntities { // loop through entities
					// Skip entities with no needs (like Grass)
					if len(entity.Needs) == 0 {
						continue
						//Add in tick grass here
					}

					tickNeed(entity)
					if !entity.CheckCurrentCell(worldMap, ResourceType(getLowestNeedtype(entity))) {
						entity.MoveEntity(worldMap) //Move if we cant find out lowest need type at the current location
					}
				}
			}
		}
	}
	worldMap.PrintWorldMap()

}
