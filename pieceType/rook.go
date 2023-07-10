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

	p := pos.North()
	for p.Valid() {
		positions = append(positions, p)
		p = p.North()
	}

	p = pos.South()
	for p.Valid() {
		positions = append(positions, p)
		p = p.South()
	}

	p = pos.West()
	for p.Valid() {
		positions = append(positions, p)
		p = p.West()
	}

	p = pos.East()
	for p.Valid() {
		positions = append(positions, p)
		p = p.East()
	}

	return positions
}
