package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find word you were looking for")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}
