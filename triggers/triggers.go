package triggers

import (
	"github.com/ghostsquad/space-z/player"
	"github.com/ghostsquad/space-z/maps"
)

type TriggerFunction func(p *player.Player, m maps.GameMap) (bool, error)

type Event int

const (
	PLAYER_MOVED Event = iota
)

func FireTrigger(name string, function TriggerFunction) {

}
