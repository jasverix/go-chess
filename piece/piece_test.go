package piece

import (
	"github.com/stretchr/testify/assert"
	"go-chess/color"
	"go-chess/pieceType"
	"go-chess/position"
	"sort"
	"testing"
)

func TestPiece_PossibleTargets_As_Pawn(t *testing.T) {
	p := New(pieceType.Pawn(), color.White, position.FromString("C2"))
	targets := p.PossibleTargets()
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"B3", "C3", "D3"})
}

func TestPiece_PossibleTargets_As_Pawn_at_edge(t *testing.T) {
	p := New(pieceType.Pawn(), color.White, position.FromString("A2"))
	targets := p.PossibleTargets()
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"A3", "B3"})
}
