package world

type Vec2 struct {
	XPos int
	YPos int
}
type Entity struct {
	Name         string
	Position     *Vec2
	Alive        bool
	Produces     []ResourceEntry
	Needs        map[string]*NeedEntry
	ShelterPrefs []string
	Home         *Vec2
	Aversions    []AversionEntry
}
type ResourceEntry struct {
	Type      string
	Current   float64 //Current acoumt
	Max       float64 // Max amount
	RegenRate float64 //does this regen amount
}

/*
	Will need to add a type filter to NeedEntry. 1. food -> so get all of the food types 2. Shelter -> so they go back home after hunger is fulfilled 2. Pro? so they go to do
	food before shelter? This will later expand into havng need for renown / quests etc
*/

type NeedEntry struct {
	Resource    string //this maps to type of Resource Entry
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
