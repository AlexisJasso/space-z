package player

import (
	"github.com/ghostsquad/space-z/maps"
	"github.com/pkg/errors"
)

type player struct {
	currentRoom *maps.Room
}

type Player interface {
	Move(direction maps.CardinalDirection) (bool, error)
	GetLocation() *maps.Room
	SetLocation(room *maps.Room) (bool, error)
}

func (p *player) Move(direction maps.CardinalDirection) (bool, error) {
	return false, errors.New("Not implemented")
}

func (p *player) GetLocation() *maps.Room {
	return maps.VoidRoom
}

func (p *player) SetLocation(room *maps.Room) (bool, error) {
	return false, errors.New("Not implemented")
}
