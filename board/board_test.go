package board

import (
	"github.com/stretchr/testify/assert"
	"go-chess/color"
	"go-chess/piece"
	"go-chess/pieceType"
	"go-chess/position"
	"testing"
)

func TestNew(t *testing.T) {
	board := New()
	assert.True(t, board.IsEmpty(position.FromString("A1")))
}

func TestBoard_AddPiece(t *testing.T) {
	board := New()
	p1 := piece.New(pieceType.Pawn(), color.White, position.FromString("G4"))
	err := board.AddPiece(p1)
	assert.NoError(t, err)
	assert.False(t, board.IsEmpty(position.FromString("G4")))
	assert.True(t, board.IsEmpty(position.FromString("A1")))
}

func TestBoard_AddPiece_invalid_position(t *testing.T) {
	board := New()
	p1 := piece.New(pieceType.Pawn(), color.White, position.FromString("F9"))
	err := board.AddPiece(p1)
	assert.Error(t, err)
	assert.Equal(t, "position of piece is not valid", err.Error())
}

func TestBoard_AddPiece_duplicate_check(t *testing.T) {
	board := New()
	p1 := piece.New(pieceType.Pawn(), color.White, position.FromString("G4"))
	err := board.AddPiece(p1)
	assert.NoError(t, err)

	p2 := piece.New(pieceType.Rook(), color.White, position.FromString("G4"))
	err = board.AddPiece(p2)
	assert.Error(t, err)
}
