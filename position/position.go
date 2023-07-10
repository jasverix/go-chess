package position

import (
	"fmt"
)

type position struct {
	lat int8 // A-H
	lng int8 // 1-8
}

type Position interface {
	North() Position
	South() Position
	West() Position
	East() Position

	Valid() bool
	String() string
}

func intToLat(lat uint8) int8 {
	if lat <= 'H' && lat >= 'A' {
		return int8(lat - 'A')
	}
	return -1
}

func intToLng(lng int8) int8 {
	if lng < 8 {
		return lng - 1
	}
	return -1
}

var invalidPosition = position{-1, -1}

func New(lat uint8, lng int8) Position {
	realLat := intToLat(lat)
	realLng := intToLng(lng)

	return position{lat: realLat, lng: realLng}
}

func FromString(pos string) Position {
	if len(pos) != 2 {
		return invalidPosition
	}

	lat := pos[0]
	var lng int8
	_, err := fmt.Sscanf(string(pos[1]), "%d", &lng)
	if err != nil {
		return invalidPosition
	}
	return New(lat, lng)
}

func (p position) North() Position {
	if !p.Valid() {
		return position{-1, -1}
	}
	return position{p.lat, p.lng + 1}
}

func (p position) South() Position {
	if !p.Valid() {
		return position{-1, -1}
	}
	return position{p.lat, p.lng - 1}
}

func (p position) East() Position {
	if !p.Valid() {
		return position{-1, -1}
	}
	return position{p.lat + 1, p.lng}
}

func (p position) West() Position {
	if !p.Valid() {
		return position{-1, -1}
	}
	return position{p.lat - 1, p.lng}
}

func (p position) Valid() bool {
	return p.lat >= 0 && p.lng >= 0 && p.lat <= 7 && p.lng <= 7
}

func (p position) String() string {
	if !p.Valid() {
		return "XX"
	}
	return fmt.Sprintf("%c%d", p.lat+'A', p.lng+1)
}
