package pieceType

import (
	"github.com/stretchr/testify/assert"
	"go-chess/color"
	"go-chess/position"
	"sort"
	"testing"
)

func TestPawn_PossibleTargets_white_normal(t *testing.T) {
	targets := Pawn().PossibleTargets(position.New('C', 2), color.White, Normal)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"C3"})
}

func TestPawn_PossibleTargets_white_attack(t *testing.T) {
	targets := Pawn().PossibleTargets(position.New('C', 2), color.White, Attack)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"B3", "D3"})
}

func TestPawn_PossibleTargets_black_normal(t *testing.T) {
	targets := Pawn().PossibleTargets(position.New('C', 7), color.Black, Normal)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"C6"})
}

func TestPawn_PossibleTargets_black_attack(t *testing.T) {
	targets := Pawn().PossibleTargets(position.New('C', 7), color.Black, Attack)
	targetStrings := make([]string, len(targets))
	for i := range targets {
		targetStrings[i] = targets[i].String()
	}
	sort.Strings(targetStrings)
	assert.EqualValues(t, targetStrings, []string{"B6", "D6"})
}
