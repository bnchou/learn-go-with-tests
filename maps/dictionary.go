package maps

import (
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Add(key, value string) {
	d[key] = value
}

func (d Dictionary) Search(s string) (string, error) {
	res, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}
	return res, nil
}
