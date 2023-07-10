package pieceType

import (
	"go-chess/color"
	"go-chess/position"
)

type PieceType interface {
	PossibleTargets(position.Position, color.Color) []position.Position
}
