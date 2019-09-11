package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrWordExists = errors.New("cannot add word because it already exists")
	ErrNotFound   = errors.New("could not find the word you were looking for")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
