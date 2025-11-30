# Refactor Overview:

World
  - Width, Height
  - Cells[y][x] with TerrainType and maybe static resources
  - Entities: map[EntityID]*Entity   (global registry)
  - CellEntities: map[Vec2][]EntityID (who is standing in each tile)

Entity
  - ID, Name
  - Position Vec2
  - Needs map[ResourceType]*NeedEntry
  - Produces []ResourceEntry
  - Behavior: derived from needs each tick

Tick()
  - For each entity:
      - tickNeeds()
      - decideAction()
      - maybe move to neighboring cell
      - maybe consume or gain resources
