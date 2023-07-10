package position

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Position_to_string(t *testing.T) {
	p, err := New('A', 1)
	assert.NoError(t, err)
	assert.Equal(t, "A1", p.String())
}
