package pieceType

import (
	"github.com/stretchr/testify/assert"
	"go-chess/color"
	"go-chess/position"
	"sort"
	"testing"
)

func TestQueen_PossibleTargets(t *testing.T) {
	targets := Queen().PossibleTargets(position.FromString("D5"), color.White, Normal)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, []string{"A2", "A5", "A8", "B3", "B5", "B7", "C4", "C5", "C6", "D1", "D2", "D3", "D4", "D6", "D7", "D8", "E4", "E5", "E6", "F3", "F5", "F7", "G2", "G5", "G8", "H1", "H5"}, targetStrings)
}
