package pieceType

import (
	"github.com/stretchr/testify/assert"
	"go-chess/color"
	"go-chess/position"
	"sort"
	"testing"
)

func TestKing_PossibleTargets(t *testing.T) {
	targets := King().PossibleTargets(position.FromString("D5"), color.White, 0)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, []string{"C4", "C5", "C6", "D4", "D6", "E4", "E5", "E6"}, targetStrings)
}
