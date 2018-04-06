package maps

type Room struct {
	Name string
	Neighbors map[CardinalDirection]*Room
}

type CardinalDirection int

const (
	North CardinalDirection = iota
	South
	East
	West
)
