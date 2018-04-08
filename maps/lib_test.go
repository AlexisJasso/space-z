package maps

// Basic imports
import (
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/assert"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type TestSuite struct {
	suite.Suite
	room *Room
}

const TestRoom RoomName = "Test Room"

// This is an actual test case:
func (this *TestSuite) TestNewRoomShouldSetName() {
	expectedName := TestRoom
	r := NewRoom(expectedName)

	assert.Equal(this.T(), r.Name, expectedName)
	assert.False(this.T(), r.Void)
}
