package rooms

import "github.com/ghostsquad/space-z/maps"

const Cafeteria = "Cafeteria"

func NewCafeteriaRoom() *maps.Room {
	return maps.NewRoom(Cafeteria)
}
