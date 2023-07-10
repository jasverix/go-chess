package pieceType

import (
	"github.com/stretchr/testify/assert"
	"go-chess/color"
	"go-chess/position"
	"sort"
	"testing"
)

func TestBishop_PossibleTargets(t *testing.T) {
	targets := Bishop().PossibleTargets(position.FromString("D5"), color.White)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, []string{"A2", "A8", "B3", "B7", "C4", "C6", "E4", "E6", "F3", "F7", "G2", "G8", "H1"}, targetStrings)
}
