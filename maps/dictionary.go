package maps

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("definition already exist")
)

type Dictionary map[string]string

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		return ErrWordExists
	case ErrNotFound:
		d[key] = value
		return nil
	default:
		return err
	}
}

func (d Dictionary) Search(s string) (string, error) {
	res, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}
	return res, nil
}
