package maps

import (
	"errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(s string) (string, error) {
	res, ok := d[s]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return res, nil
}
