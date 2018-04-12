package maps

type RoomName string

type Room struct {
	Name 			RoomName
	Neighbors map[CardinalDirection]*Room
	Void 			bool
}

func NewRoom(name RoomName) *Room {
	return &Room{
		Name: name,
		Neighbors: map[CardinalDirection]*Room{},
	}
}

var VoidRoom = &Room{Void: true}

type CardinalDirection int

const (
	North CardinalDirection = iota
	South
	East
	West
)

var ValidDirections = []CardinalDirection{
	North,
	South,
	East,
	West,
}

func (r *Room) SetRoomNeighbors(directions map[CardinalDirection]*Room) {
	// assert that all valid directions are filled with void rooms
	for _, d := range ValidDirections {
		if _, ok := directions[d]; ok {
			r.Neighbors[d] = directions[d]
		} else {
			r.Neighbors[d] = VoidRoom
		}
	}
}

type GameMap interface {
	StartingRoom() *Room
	AvailableRooms() map[RoomName]*Room
}
