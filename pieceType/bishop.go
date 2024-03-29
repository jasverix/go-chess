package pieceType

import (
	"go-chess/color"
	"go-chess/position"
)

type bishop struct {
}

func Bishop() PieceType {
	return bishop{}
}

func bishopPaths(pos position.Position) []position.Position {
	var positions []position.Position
	positions = append(positions, pos.Walk(position.NorthWest)...)
	positions = append(positions, pos.Walk(position.NorthEast)...)
	positions = append(positions, pos.Walk(position.SouthWest)...)
	positions = append(positions, pos.Walk(position.SouthEast)...)
	return positions
}

func (r bishop) PossibleTargets(pos position.Position, _ color.Color, _ MoveMode) []position.Position {
	return bishopPaths(pos)
}
