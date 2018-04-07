package spaceship

import (
	. "github.com/ghostsquad/space-z/maps"
)

/*

                      +-------+
                      |Space  |
                      +---+---+
                          |
            +------+  +---+---+  +------+
            |Space +--+Bridge +--+Space |
            +------+  +---+---+  +------+
                          |
+-----+    +-------+  +---+---+  +-----------+
|Mars +----+AirLock+--+Hallway+--+Engineering|
+-----+    +-------+  +---+---+  +-----------+
                          |
+------+   +-------+  +---+------+
|Cafe  +---+Showers+--+LockerRoom|
+------+   +-------+  +----+-----+
                           |
                      +----+----+
                      |CargoBay |
                      +---------+


*/

func NewMap() *Room {
	airlock := NewRoom("Airlock")
	space := NewRoom("Space")
	bridge := NewRoom("Bridge")
	hallway := NewRoom("Hallway")
	lockerRoom := NewRoom("CoEd Locker Room")
	showers := NewRoom("CoEd Showers")
	cargoBay := NewRoom("Cargobay")
	cafe := NewRoom("Cafe")
	marsSurface := NewRoom("Mars Planet Surface")
	engineering := NewRoom("Engineering")

	airlock.Neighbors[East] = hallway
	airlock.Neighbors[West] = marsSurface

	hallway.Neighbors[North] = bridge
	hallway.Neighbors[East] = engineering
	hallway.Neighbors[South] = lockerRoom
	hallway.Neighbors[West] = airlock

	bridge.Neighbors[North] = space
	bridge.Neighbors[East] = space
	bridge.Neighbors[South] = hallway
	bridge.Neighbors[West] = space

	engineering.Neighbors[West] = hallway

	lockerRoom.Neighbors[North] = hallway
	lockerRoom.Neighbors[South] = cargoBay
	lockerRoom.Neighbors[West] = showers

	showers.Neighbors[East] = lockerRoom
	showers.Neighbors[West] = cafe

	return airlock
}
