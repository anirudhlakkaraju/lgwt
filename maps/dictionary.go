package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", errors.New("could not find word you were looking for")
	}

	return value, nil
}
