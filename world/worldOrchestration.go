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
					tickNeed(entity)
					entity.MoveEntity(worldMap)
				}
			}
		}
	}

}
