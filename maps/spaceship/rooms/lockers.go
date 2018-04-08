package rooms

import "github.com/ghostsquad/space-z/maps"

const Lockers = "Lockers"

func NewLockersRoom() *maps.Room {
	return maps.NewRoom(Lockers)
}
