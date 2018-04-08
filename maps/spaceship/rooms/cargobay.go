package rooms

import "github.com/ghostsquad/space-z/maps"

const CargoBay = "Cargo Bay"

func NewCargoBayRoom() *maps.Room {
	return maps.NewRoom(CargoBay)
}
