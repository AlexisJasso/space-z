package rooms

import "github.com/ghostsquad/space-z/maps"

const Airlock maps.RoomName = "Airlock"

func NewAirlockRoom() *maps.Room {
	return maps.NewRoom(Airlock)
}
