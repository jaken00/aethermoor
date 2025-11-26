package world

func (e *Entity) tickNeed(needType string) {
	need, exists := e.Needs[needType]
	if !exists {
		return
	}

	need.Current -= 1

}

func (e *Entity) needComparison(needType string) bool {
	need, exists := e.Needs[needType]
	if !exists {
		return false
	}
	return need.Current < need.Threshold
}

func (e *Entity) locateNeed(needType string, cell *Cell, worldMap *WorldMap) {
	need, exists := e.Needs[needType]
	if !exists {
		return
	}
	needDict := worldMap.ResouceTerrainDict.ResourceDictionary
	terrainTypeHolder := needDict[needType]
	/*
		1. get current postion of the entity
		2. determine the terrain types surrounding this position
		3. go to that terrain type depending on need
	*/
	currentPosX := e.Position.XPos
	currentPosY := e.Position.YPos

}

func (worldMap *WorldMap) GetNeighbors(xPos int, yPos int, terrainType bool) bool {

	possible_values := [8]Vec2{{0, 1}, {0, -1}, {1, 1}, {1, -1}, {1, 0}, {1, 1}, {-1, 0}, {-1, -1}} //for loop through these
	//if terrain type matches then return else continue
	//if no terrain type lets change this to bool we return false

}

func (e *Entity) ActOrchestrator() {
	foodBool := e.needComparison("food")
	shelterBool := e.needComparison("shelter")

	if foodBool && shelterBool {
		// go get food
	} else if foodBool {
		//go get food
	} else {
		//go to shelter -> if no shelter close just walk to a random 8 square
		//to do this we need to create a for loop and break when we have the first choice
		//need a tick food as well
	}

}
