package pieceType

import (
	"github.com/stretchr/testify/assert"
	"go-chess/color"
	"go-chess/position"
	"sort"
	"testing"
)

func TestRook_PossibleTargets(t *testing.T) {
	targets := Rook().PossibleTargets(position.FromString("D5"), color.White, 0)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, []string{"A5", "B5", "C5", "D1", "D2", "D3", "D4", "D6", "D7", "D8", "E5", "F5", "G5", "H5"}, targetStrings)
}
