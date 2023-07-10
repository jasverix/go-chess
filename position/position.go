package position

import (
	"fmt"
)

type position struct {
	lat int8 // A-H
	lng int8 // 1-8
}

type Direction int8

const (
	North Direction = iota
	West
	East
	South
	NorthEast
	NorthWest
	SouthEast
	SouthWest
)

type Position interface {
	Step(direction Direction) Position
	Walk(direction Direction) []Position
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

const stringFormat = "%c%d"

func New(lat uint8, lng int8) Position {
	realLat := intToLat(lat)
	realLng := intToLng(lng)

	return position{lat: realLat, lng: realLng}
}

func FromString(pos string) Position {
	if len(pos) != 2 {
		return invalidPosition
	}

	var lat uint8
	var lng int8
	_, err := fmt.Sscanf(pos, stringFormat, &lat, &lng)
	if err != nil {
		return invalidPosition
	}

	return New(lat, lng)
}

func (p position) Valid() bool {
	return p.lat >= 0 && p.lng >= 0 && p.lat <= 7 && p.lng <= 7
}

func (p position) String() string {
	if !p.Valid() {
		return "XX"
	}
	return fmt.Sprintf(stringFormat, p.lat+'A', p.lng+1)
}
