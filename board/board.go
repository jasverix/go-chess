package board

import (
	"fmt"
	"go-chess/piece"
	"go-chess/position"
)

type board struct {
	fields map[position.Position]*field
}

type Board interface {
	IsEmpty(pos position.Position) bool
	AddPiece(piece piece.Piece) error
}

func emptyBoardFields() map[position.Position]*field {
	res := make(map[position.Position]*field)
	init := position.FromString("A1")
	res[init] = &field{}

	for _, col := range init.Walk(position.North) {
		res[col] = &field{}
		for _, row := range col.Walk(position.East) {
			res[row] = &field{}
		}
	}

	return res
}

func New() Board {
	return &board{emptyBoardFields()}
}

func (b *board) IsEmpty(pos position.Position) bool {
	if !pos.Valid() {
		return false
	}
	return b.fields[pos].piece == nil
}

func (b *board) AddPiece(p piece.Piece) error {
	if !p.Position().Valid() {
		return fmt.Errorf("position of piece is not valid")
	}
	if !b.IsEmpty(p.Position()) {
		return fmt.Errorf("position %s already occupied by %s piece of type %s", p.Position(), p.Color(), p.Type())
	}

	b.fields[p.Position()].piece = p

	return nil
}
