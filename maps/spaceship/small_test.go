package spaceship

// Basic imports
import (
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/assert"
	"github.com/ghostsquad/space-z/maps"
	"github.com/ghostsquad/space-z/maps/spaceship/rooms"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSmallSpaceShipSuite(t *testing.T) {
	suite.Run(t, new(SmallSpaceShipSuite))
}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type SmallSpaceShipSuite struct {
	suite.Suite
	gameMap maps.GameMap
}

func (this *SmallSpaceShipSuite) AssertNeighbors(r *maps.Room, directions map[maps.CardinalDirection]*maps.Room) {
	for _, d := range maps.ValidDirections {
		if _, ok := directions[d]; ok {
			assert.Equal(this.T(), r.Neighbors[d], directions[d])
		} else {
			// assert that all other valid directions are filled with void rooms
			assert.Equal(this.T(), r.Neighbors[d], maps.VoidRoom)
		}
	}
}

func (this *SmallSpaceShipSuite) AssertNeighbor(r *maps.Room, d maps.CardinalDirection, e *maps.Room){
	assert.Equal(this.T(), r.Neighbors[d], e)
	assert.False(this.T(), r.Neighbors[d].Void)
}

func (this *SmallSpaceShipSuite) AssertNeighborIsVoid(r *maps.Room, d maps.CardinalDirection) {
	assert.Equal(this.T(), r.Neighbors[d], maps.VoidRoom)
}

func (this *SmallSpaceShipSuite) SetupTest() {
	this.gameMap = NewSmallSpaceShip()
}

// This is an actual test case:
func (this *SmallSpaceShipSuite) TestShouldStartInAirlock() {
	assert.Equal(this.T(), this.gameMap.StartingRoom(), this.gameMap.AvailableRooms()[rooms.Airlock])
}

func (this *SmallSpaceShipSuite) TestAirLockConnections() {
	a := this.gameMap.AvailableRooms()
	l := a[rooms.Airlock]

	this.AssertNeighbors(l, map[maps.CardinalDirection]*maps.Room{
		maps.West: a[rooms.PlanetSurface],
		maps.East: a[rooms.Hallway],
	})
}

func (this *SmallSpaceShipSuite) TestHallwayConnections() {
	a := this.gameMap.AvailableRooms()
	l := a[rooms.Hallway]

	this.AssertNeighbors(l, map[maps.CardinalDirection]*maps.Room{
		maps.North: a[rooms.Bridge],
		maps.South: a[rooms.Lockers],
		maps.West: a[rooms.Airlock],
		maps.East: a[rooms.Engineering],
	})
}

func (this *SmallSpaceShipSuite) TestEngineeringConnections() {
	a := this.gameMap.AvailableRooms()
	l := a[rooms.Engineering]

	this.AssertNeighbors(l, map[maps.CardinalDirection]*maps.Room{
		maps.West: a[rooms.Hallway],
	})
}

func (this *SmallSpaceShipSuite) TestBridgeConnections() {
	a := this.gameMap.AvailableRooms()
	l := a[rooms.Bridge]

	this.AssertNeighbors(l, map[maps.CardinalDirection]*maps.Room{
		maps.North: a[rooms.Space],
		maps.South: a[rooms.Hallway],
		maps.East: a[rooms.Space],
		maps.West: a[rooms.Space],
	})
}

func (this *SmallSpaceShipSuite) TestLockersConnections() {
	a := this.gameMap.AvailableRooms()
	l := a[rooms.Lockers]

	this.AssertNeighbors(l, map[maps.CardinalDirection]*maps.Room{
		maps.North: a[rooms.Hallway],
		maps.South: a[rooms.CargoBay],
		maps.West: a[rooms.Showers],
	})
}

func (this *SmallSpaceShipSuite) TestShowersConnections() {
	a := this.gameMap.AvailableRooms()
	l := a[rooms.Showers]

	this.AssertNeighbors(l, map[maps.CardinalDirection]*maps.Room{
		maps.East: a[rooms.Lockers],
		maps.West: a[rooms.Cafeteria],
	})
}

func (this *SmallSpaceShipSuite) TestCargoBayConnections() {
	a := this.gameMap.AvailableRooms()
	l := a[rooms.CargoBay]

	this.AssertNeighbors(l, map[maps.CardinalDirection]*maps.Room{
		maps.North: a[rooms.Lockers],
	})
}
