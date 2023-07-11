package pieceType

import (
	"go-chess/color"
	"go-chess/position"
)

type pawn struct{}

func Pawn() PieceType {
	return pawn{}
}

func (p pawn) PossibleTargets(pos position.Position, col color.Color, mode MoveMode) []position.Position {
	switch mode {
	case Normal:
		switch col {
		case color.Black:
			return []position.Position{
				pos.South(),
			}
		case color.White:
			return []position.Position{
				pos.North(),
			}
		}
	case Attack:
		switch col {
		case color.Black:
			return []position.Position{
				pos.South().East(),
				pos.South().West(),
			}
		case color.White:
			return []position.Position{
				pos.North().East(),
				pos.North().West(),
			}
		}
	}
	return []position.Position{}
}
