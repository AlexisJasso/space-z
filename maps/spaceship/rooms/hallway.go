package rooms

import "github.com/ghostsquad/space-z/maps"

const Hallway = "Hallway"

func NewHallwayRoom() *maps.Room {
	return maps.NewRoom(Hallway)
}
