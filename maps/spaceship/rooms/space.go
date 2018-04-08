package rooms

import "github.com/ghostsquad/space-z/maps"

const Space = "Space"

func NewSpaceRoom() *maps.Room {
	return maps.NewRoom(Space)
}
