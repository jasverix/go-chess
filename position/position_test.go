package position

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Position_A1_to_string(t *testing.T) {
	p := New('A', 1)
	assert.Equal(t, "A1", p.String())
	assert.True(t, p.Valid())
}

func TestPosition_Valid(t *testing.T) {
	assert.False(t, New('I', 2).Valid())
	assert.False(t, New('B', 9).Valid())
	assert.False(t, New('B', -2).Valid())
	assert.False(t, New('A'-1, 2).Valid())
}

func TestPosition_North(t *testing.T) {
	assert.Equal(t, "B2", New('B', 1).North().String())
	assert.Equal(t, "C8", New('C', 7).North().String())
	assert.Equal(t, "XX", New('C', 8).North().String())
	assert.Equal(t, "XX", New('F', -1).North().String())
}

func TestPosition_East(t *testing.T) {
	assert.Equal(t, "E1", New('D', 1).East().String())
	assert.Equal(t, "G7", New('F', 7).East().String())
}

func TestPosition_West(t *testing.T) {
	assert.Equal(t, "C1", New('D', 1).West().String())
}

func TestPosition_South(t *testing.T) {
	assert.Equal(t, "F4", New('F', 5).South().String())
}

func TestFromString(t *testing.T) {
	assert.Equal(t, "A1", FromString("A1").String())
	assert.Equal(t, "B5", FromString("B5").String())
	assert.Equal(t, "XX", FromString("I7").String())
	assert.Equal(t, "XX", FromString("C9").String())
}
