package world

type Vec2 struct {
	XPos int
	YPos int
}
type Entity struct {
	Name     string
	Position *Vec2
	Alive    bool
	Produces []ResourceEntry
	Needs    []NeedEntry
	//ShelterPrefs []string
	Home      *Vec2
	Aversions []AversionEntry
}
type ResourceEntry struct {
	Type      string
	Current   float64 //Current acoumt
	Max       float64 // Max amount
	RegenRate float64 //does this regen amount
}
type NeedEntry struct {
	Resource    string //this maps to type of Resource Entry
	Threshold   float64
	Capacity    float64
	ConsumeRate float64
	MinInterest float64
}
type AversionEntry struct {
	Resource  string  // flee from entities that produce this resource -> WOLF -> CARNIVOREMEAT the MEAT is descriptor of the production
	FleeRange float64 // within certain distance we want this to flee
}
