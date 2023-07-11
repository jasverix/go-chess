package pieceType

import (
	"go-chess/color"
	"go-chess/position"
)

type queen struct {
}

func Queen() PieceType {
	return queen{}
}

func (q queen) PossibleTargets(pos position.Position, _ color.Color, _ MoveMode) []position.Position {
	var positions []position.Position
	positions = append(positions, rookPaths(pos)...)
	positions = append(positions, bishopPaths(pos)...)
	return positions
}
