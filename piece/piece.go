package piece

import (
	"go-chess/color"
	"go-chess/pieceType"
	"go-chess/position"
)

type Piece interface {
	Color() color.Color
	Type() pieceType.PieceType
	Position() position.Position

	PossibleTargets(pieceType.MoveMode) []position.Position
}

type piece struct {
	position  position.Position
	color     color.Color
	pieceType pieceType.PieceType
}

func New(pieceType pieceType.PieceType, color color.Color, position position.Position) Piece {
	return piece{position, color, pieceType}
}

func (p piece) Color() color.Color {
	return p.color
}

func (p piece) Type() pieceType.PieceType {
	return p.pieceType
}

func (p piece) Position() position.Position {
	return p.position
}

func (p piece) PossibleTargets(moveMode pieceType.MoveMode) []position.Position {
	var targets []position.Position
	for _, pos := range p.pieceType.PossibleTargets(p.position, p.color, moveMode) {
		if pos.Valid() {
			targets = append(targets, pos)
		}
	}
	return targets
}
