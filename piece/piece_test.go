package piece

import (
	"github.com/stretchr/testify/assert"
	"go-chess/color"
	"go-chess/pieceType"
	"go-chess/position"
	"sort"
	"testing"
)

func TestPiece_PossibleTargets_As_Pawn_normal(t *testing.T) {
	p := New(pieceType.Pawn(), color.White, position.FromString("C2"))
	targets := p.PossibleTargets(pieceType.Normal)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"C3"})
}

func TestPiece_PossibleTargets_As_Pawn_attack(t *testing.T) {
	p := New(pieceType.Pawn(), color.White, position.FromString("C2"))
	targets := p.PossibleTargets(pieceType.Attack)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"B3", "D3"})
}

func TestPiece_PossibleTargets_As_Pawn_at_edge(t *testing.T) {
	p := New(pieceType.Pawn(), color.White, position.FromString("A2"))
	targets := p.PossibleTargets(pieceType.Attack)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"B3"})
}
