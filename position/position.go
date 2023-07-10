package position

import (
	"errors"
	"fmt"
)

type position struct {
	lat int8 // A-H
	lng int8 // 1-8
}

type Position interface {
	String() string
}

func intToLat(lat int8) int8 {
	if lat <= 'H' && lat >= 'A' {
		return lat - 'A'
	}
	return -1
}

func intToLng(lng int8) int8 {
	if lng < 8 {
		return lng - 1
	}
	return -1
}

func New(lat int8, lng int8) (Position, error) {
	realLat := intToLat(lat)
	realLng := intToLng(lng)

	if realLat == -1 {
		return nil, errors.New("invalid value for lat")
	}
	if realLng == -1 {
		return nil, errors.New("invalid value for lng")
	}

	return position{
		lat: realLat,
		lng: realLng,
	}, nil
}

func (p position) String() string {
	return fmt.Sprintf("%c%d", p.lat+'A', p.lng+1)
}
