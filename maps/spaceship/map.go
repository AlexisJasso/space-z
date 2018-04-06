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
	airlock := Room{
		Name: "Airlock",
	}

	space := Room{
		Name: "Space",
	}

	bridge := Room{
		Name: "Bridge",
	}

	hallway := Room{
		Name: "Hallway",
	}

	lockerRoom := Room{
		Name: "CoEd LockerRoom",
	}

	showers := Room{
		Name: "CoEd Showers",
	}

	cargoBay := Room{
		Name: "CargoBay",
	}

	cafe := Room{
		Name: "Cafe",
	}

	mars := Room{
		Name: "Mars",
	}

	engineering := Room{
		Name: "Engineering",
	}

	airlock.Neighbors[East] = &hallway
	airlock.Neighbors[West] = &mars

	hallway.Neighbors[North] = &bridge
	hallway.Neighbors[East] = &engineering
	hallway.Neighbors[South] = &lockerRoom
	hallway.Neighbors[West] = &airlock

	bridge.Neighbors[North] = &space
	bridge.Neighbors[East] = &space
	bridge.Neighbors[South] = &hallway
	bridge.Neighbors[West] = &space

	engineering.Neighbors[West] = &hallway

	lockerRoom.Neighbors[North] = &hallway
	lockerRoom.Neighbors[South] = &cargoBay
	lockerRoom.Neighbors[West] = &showers

	showers.Neighbors[East] = &lockerRoom
	showers.Neighbors[West] = &cafe

	return &airlock
}
