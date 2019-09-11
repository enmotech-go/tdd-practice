package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if !exist {
		return "", errors.New("not found")
	}
	return value, nil
}
