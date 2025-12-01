package world

type TerrainType string
type ResourceType string
type NeedType string

const (
	TerrainPlains    TerrainType = "PLAINS"
	TerrainWoods     TerrainType = "WOODS"
	TerrainMountain  TerrainType = "MOUNTAIN"
	TerrainRiver     TerrainType = "RIVER"
	TerrainCave      TerrainType = "CAVE"
	TerrainGrassland TerrainType = "GRASSLAND"
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

type Vec2 struct {
	XPos int
	YPos int
}
type Entity struct {
	Name         string
	Position     *Vec2
	Alive        bool
	Produces     []ResourceEntry
	Needs        map[NeedType]*NeedEntry
	ShelterPrefs []string
	Home         *Vec2
	Aversions    []AversionEntry
}
type ResourceEntry struct {
	Type      ResourceType
	Current   float64 //Current acoumt
	Max       float64 // Max amount
	RegenRate float64 //does this regen amount
}

type ResourceTerrainMapping struct {
	ResourceDictionary map[string][]string
}

type NeedEntry struct {
	Resource    ResourceType
	Kind        NeedType
	Type        string
	Priority    int
	Threshold   float64 // when we start looking for this resource
	Current     float64
	Capacity    float64
	ConsumeRate float64
	MinInterest float64 //ignores patches with certain amoutn of grass
}
type AversionEntry struct {
	Resource  string  // flee from entities that produce this resource -> WOLF -> CARNIVOREMEAT the MEAT is descriptor of the production
	FleeRange float64 // within certain distance we want this to flee
}
