package rooms

import "github.com/ghostsquad/space-z/maps"

const Showers = "Showers"

func NewShowersRoom() *maps.Room {
	return maps.NewRoom(Showers)
}
