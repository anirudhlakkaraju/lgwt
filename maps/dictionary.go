package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find word you were looking for")
var ErrWordExists = errors.New("word already exists in dictionary")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]
	if !ok {
		d[key] = value
		return nil
	}
	return ErrWordExists
}
