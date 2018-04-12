package spaceship

import (
	. "github.com/ghostsquad/space-z/maps"
	"github.com/ghostsquad/space-z/maps/spaceship/rooms"
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
|Cafe  +---+Showers+--+Lockers   |
+------+   +-------+  +----+-----+
                           |
                      +----+----+
                      |CargoBay |
                      +---------+


*/

type SmallSpaceShip struct {
	rooms map[RoomName]*Room
	startingRoom *Room
}

func (s *SmallSpaceShip) StartingRoom() *Room {
	return s.rooms[rooms.Airlock]
}

func (s *SmallSpaceShip) AvailableRooms() map[RoomName]*Room {
	return s.rooms
}

func NewSmallSpaceShip() *SmallSpaceShip {
	s := SmallSpaceShip{}
	s.rooms = map[RoomName]*Room{}

	roomsList := []RoomName{
		rooms.Airlock,
		rooms.Hallway,
		rooms.Bridge,
		rooms.Engineering,
		rooms.Lockers,
		rooms.Showers,
		rooms.PlanetSurface,
		rooms.Cafeteria,
		rooms.CargoBay,
		rooms.Space,
	}

	for _, r := range roomsList {
		s.rooms[r] = NewRoom(r)
	}

	s.rooms[rooms.Airlock].SetRoomNeighbors(map[CardinalDirection]*Room{
		West: s.rooms[rooms.PlanetSurface],
		East: s.rooms[rooms.Hallway],
	})

	s.rooms[rooms.Hallway].SetRoomNeighbors(map[CardinalDirection]*Room{
		North: s.rooms[rooms.Bridge],
		South: s.rooms[rooms.Lockers],
		West: s.rooms[rooms.Airlock],
		East: s.rooms[rooms.Engineering],
	})

	s.rooms[rooms.Bridge].SetRoomNeighbors(map[CardinalDirection]*Room{
		North: s.rooms[rooms.Space],
		South: s.rooms[rooms.Hallway],
		West: s.rooms[rooms.Space],
		East: s.rooms[rooms.Space],
	})

	s.rooms[rooms.Engineering].SetRoomNeighbors(map[CardinalDirection]*Room{
		West: s.rooms[rooms.Hallway],
	})

	s.rooms[rooms.Lockers].SetRoomNeighbors(map[CardinalDirection]*Room{
		North: s.rooms[rooms.Hallway],
		South: s.rooms[rooms.CargoBay],
		West: s.rooms[rooms.Showers],
	})

	s.rooms[rooms.Showers].SetRoomNeighbors(map[CardinalDirection]*Room{
		West: s.rooms[rooms.Cafeteria],
		East: s.rooms[rooms.Lockers],
	})

	s.rooms[rooms.CargoBay].SetRoomNeighbors(map[CardinalDirection]*Room{
		North: s.rooms[rooms.Lockers],
	})

	s.startingRoom = s.rooms[rooms.Airlock]

	return &s
}
