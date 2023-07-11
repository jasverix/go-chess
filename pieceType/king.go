package pieceType

import (
	"go-chess/color"
	"go-chess/position"
)

type king struct {
}

func King() PieceType {
	return king{}
}

func (k king) PossibleTargets(pos position.Position, _ color.Color) []position.Position {
	return []position.Position{
		pos.Step(position.North),
		pos.Step(position.NorthEast),
		pos.Step(position.East),
		pos.Step(position.SouthEast),
		pos.Step(position.South),
		pos.Step(position.SouthWest),
		pos.Step(position.West),
		pos.Step(position.NorthWest),
	}
}
