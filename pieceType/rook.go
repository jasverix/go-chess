package pieceType

import (
	"go-chess/color"
	"go-chess/position"
)

type rook struct {
}

func Rook() PieceType {
	return rook{}
}

func (r rook) PossibleTargets(pos position.Position, _ color.Color) []position.Position {
	var positions []position.Position
	positions = append(positions, pos.Walk(position.North)...)
	positions = append(positions, pos.Walk(position.West)...)
	positions = append(positions, pos.Walk(position.East)...)
	positions = append(positions, pos.Walk(position.South)...)
	return positions
}
