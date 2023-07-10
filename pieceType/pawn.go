package pieceType

import (
	"go-chess/color"
	"go-chess/position"
)

type pawn struct{}

func Pawn() PieceType {
	return pawn{}
}

func (p pawn) PossibleTargets(pos position.Position, col color.Color) []position.Position {
	switch col {
	case color.Black:
		return []position.Position{
			pos.South(),
			pos.South().East(),
			pos.South().West(),
		}
	case color.White:
		return []position.Position{
			pos.North(),
			pos.North().East(),
			pos.North().West(),
		}
	}
	return []position.Position{}
}
