package rooms

import "github.com/ghostsquad/space-z/maps"

const PlanetSurface = "Planet Surface"

func NewPlanetSurfaceRoom() *maps.Room {
	return maps.NewRoom(PlanetSurface)
}
