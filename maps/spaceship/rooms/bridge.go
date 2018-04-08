package rooms

import "github.com/ghostsquad/space-z/maps"

const Bridge = "Bridge"

func NewBridgeRoom() *maps.Room {
	return maps.NewRoom(Bridge)
}
