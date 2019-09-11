package maps

import "errors"

type Dictionary map[string]string
var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)
//search from dictionary
func (d Dictionary)Search( word string) (str string,err error) {

	definition, ok := d[word]
	if !ok {
		return "",ErrNotFound
	}

	return definition, nil
}


func (d Dictionary) Add(word, definition string) (err error){
	_, err = d.Search(word)

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