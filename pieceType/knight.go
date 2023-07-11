package pieceType

import (
	"go-chess/color"
	"go-chess/position"
)

type knight struct {
}

func Knight() PieceType {
	return knight{}
}

func (k knight) PossibleTargets(pos position.Position, _ color.Color, _ MoveMode) []position.Position {
	return []position.Position{
		pos.North().North().West(),
		pos.North().North().East(),
		pos.West().West().North(),
		pos.West().West().South(),
		pos.South().South().West(),
		pos.South().South().East(),
		pos.East().East().North(),
		pos.East().East().South(),
	}
}
