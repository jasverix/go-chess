package pieceType

import (
	"github.com/stretchr/testify/assert"
	"go-chess/color"
	"go-chess/position"
	"sort"
	"testing"
)

func TestKnight_PossibleTargets(t *testing.T) {
	targets := Knight().PossibleTargets(position.FromString("D5"), color.White, 0)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, []string{"B4", "B6", "C3", "C7", "E3", "E7", "F4", "F6"}, targetStrings)
}
