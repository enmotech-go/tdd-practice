package maps

import "errors"

var ErrNotFound = errors.New("not found")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}
	return value, nil
}
