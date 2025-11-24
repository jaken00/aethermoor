package world

type Cell struct {
	CellType     string
	CellEntities []*Entity
}

type WorldMap struct {
	Grid [][]Cell
}

func GenerateWorld(x_length int, y_length int) *WorldMap {

	var worldMap WorldMap
	grid := make([][]Cell, y_length)

	for i := 0; i < x_length; i++ {
		grid[i] = make([]Cell, x_length)
		for j := 0; j < y_length; j++ {

			grid[i][j] = Cell{
				CellType:     "TO_INIT",
				CellEntities: nil,
			}

		}
	}
	worldMap.Grid = grid

	return &worldMap
}
