package pieceType

import (
	"go-chess/color"
	"go-chess/position"
)

type MoveMode int8

const (
	Normal MoveMode = iota
	Attack
)

type PieceType interface {
	PossibleTargets(position.Position, color.Color, MoveMode) []position.Position
}
