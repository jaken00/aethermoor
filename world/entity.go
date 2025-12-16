package world

import "fmt"

func Die(e *Entity, worldMap *World) {
	fmt.Printf(">>> ENTITY DIED: %s at position (%d, %d)\n", e.Name, e.Position.XPos, e.Position.YPos)
	fmt.Println("    Final needs status:")

	for needType, need := range e.Needs {

		fmt.Printf("      %s (%s): %.2f/%.2f (threshold: %.2f)\n",
			needType, need.Resource, need.Current, need.Max, need.Threshold)
	}
	fmt.Printf("Distance From Home: %d, %d\n", e.Position.XPos-e.Home.XPos, e.Position.YPos-e.Home.YPos)
	fmt.Printf("Entity Current Activity: %s\n", e.EntitySettings.Activity)
	fmt.Println("---------------------------")

	entity_pos := *e.Position

	oldCell := &worldMap.Grid[entity_pos.XPos][entity_pos.YPos]

	for i, entity := range oldCell.CellEntities {
		if entity.Name == e.Name {
			oldCell.CellEntities = append(oldCell.CellEntities[:i], oldCell.CellEntities[i+1:]...)
			break
		}
	}
	oldList := worldMap.CellEntities[entity_pos]

	for i, name := range oldList {
		if name == e.Name {
			worldMap.CellEntities[entity_pos] = append(oldList[:i], oldList[i+1:]...)
			break
		}
	}
}
