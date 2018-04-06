package maps

type Room struct {
	Name string
	Neighbors map[CardinalDirection]*Room
}

func NewRoom(name string) *Room {
	return &Room{
		Name: name,
	}
}

type CardinalDirection int

const (
	North CardinalDirection = iota
	South
	East
	West
)
