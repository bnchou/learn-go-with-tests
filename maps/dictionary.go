package maps

import (
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("definition already exist")

type Dictionary map[string]string

func (d Dictionary) Add(key, value string) error {
	if _, ok := d[key]; ok {
		return ErrWordExists
	}

	d[key] = value
	return nil
}

func (d Dictionary) Search(s string) (string, error) {
	res, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}
	return res, nil
}
