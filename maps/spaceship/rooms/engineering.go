package rooms

import "github.com/ghostsquad/space-z/maps"

const Engineering = "Engineering"

func NewEngineeringRoom() *maps.Room {
	return maps.NewRoom(Engineering)
}
