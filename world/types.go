package world

type TerrainType string
type ResourceType string
type NeedType string
type CurrentActivity string
type EntityType string

const (
	TerrainPlains    TerrainType = "PLAINS"
	TerrainWoods     TerrainType = "WOODS"
	TerrainMountain  TerrainType = "MOUNTAIN"
	TerrainRiver     TerrainType = "RIVER"
	TerrainCave      TerrainType = "CAVE"
	TerrainGrassland TerrainType = "GRASSLAND"
)
const (
	HuntingActivity CurrentActivity = "FOOD SEARCH"
	ShelterActivity CurrentActivity = "HOME SEARCH"
	NullActivity    CurrentActivity = "null"
)

const (
	ResourceGrass         ResourceType = "GRASS"
	ResourceMeat          ResourceType = "MEAT"
	ResourceCarnivoreMeat ResourceType = "CARNIVOREMEAT"
)

const (
	NeedFood    NeedType = "FOOD"
	NeedShelter NeedType = "SHELTER"
)

const (
	WolfEntity   EntityType = "WOLF"
	GrassEntity  EntityType = "GRASS"
	RabbitEntity EntityType = "RABBIT"
)

type Vec2 struct {
	XPos int
	YPos int
}
type Entity struct {
	Name           string
	Type           EntityType
	Position       *Vec2
	Alive          bool
	Produces       []ResourceEntry
	Needs          map[NeedType]*NeedEntry
	ShelterPrefs   []string
	Home           *Vec2
	Aversions      []AversionEntry
	EntitySettings *EntitySettingsEntry
}
type ResourceEntry struct {
	Type      ResourceType
	Current   float64 //Current acoumt
	Max       float64 // Max amount
	RegenRate float64 // Regeneration (Useful for Grass)
}

type ResourceTerrainMapping struct {
	ResourceDictionary map[string][]string
}
type EntitySettingsEntry struct {
	Health       int
	Attack       int
	Activity     CurrentActivity
	ActionPoints int
}

type Corpse struct {
	ResourceCorpse ResourceType
	Current        float64
}

type NeedEntry struct {
	Resource    ResourceType
	Kind        NeedType
	Current     float64 // current satisfaction level
	Max         float64 // maximum satisfaction level
	Threshold   float64 // when we start looking for this resource
	ConsumeRate float64 // how fast this need depletes
}
type AversionEntry struct {
	Resource  ResourceType // flee from entities that produce this resource -> WOLF -> CARNIVOREMEAT the MEAT is descriptor of the production
	FleeRange float64      // within certain distance we want this to flee
}
