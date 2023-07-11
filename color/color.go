package color

type Color int8

const (
	Black Color = iota
	White
)

func (c Color) String() string {
	switch c {
	case Black:
		return "Black"
	case White:
		return "White"
	}
	return "Not a color"
}
