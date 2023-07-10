package pieceType

import (
	"github.com/stretchr/testify/assert"
	"go-chess/color"
	"go-chess/position"
	"sort"
	"testing"
)

func TestPawn_PossibleTargets(t *testing.T) {
	targets := Pawn().PossibleTargets(position.New('C', 2), color.White)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"B3", "C3", "D3"})
}

func TestPawn_PossibleTargets2(t *testing.T) {
	targets := Pawn().PossibleTargets(position.New('C', 7), color.Black)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"B6", "C6", "D6"})
}
