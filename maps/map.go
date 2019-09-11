package maps

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrExist
	default:
		return err
	}
	return nil
}
