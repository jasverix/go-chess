package position

func (p position) North() Position {
	return p.Step(North)
}

func (p position) South() Position {
	return p.Step(South)
}

func (p position) East() Position {
	return p.Step(East)
}

func (p position) West() Position {
	return p.Step(West)
}

func (p position) Step(direction Direction) Position {
	if !p.Valid() {
		return position{-1, -1}
	}
	switch direction {
	case North:
		return position{p.lat, p.lng + 1}
	case East:
		return position{p.lat + 1, p.lng}
	case West:
		return position{p.lat - 1, p.lng}
	case South:
		return position{p.lat, p.lng - 1}
	case NorthWest:
		return position{p.lat - 1, p.lng + 1}
	case NorthEast:
		return position{p.lat + 1, p.lng + 1}
	case SouthWest:
		return position{p.lat - 1, p.lng - 1}
	case SouthEast:
		return position{p.lat + 1, p.lng - 1}
	}
	return invalidPosition
}

func (p position) Walk(direction Direction) []Position {
	p2 := p.Step(direction)
	var res []Position
	for p2.Valid() {
		res = append(res, p2)
		p2 = p2.Step(direction)
	}
	return res
}
