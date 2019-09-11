package maps

import "errors"

type Dictionary map[string]string
var ErrNotFound = errors.New("could not find the word you were looking for")
//search from dictionary
func (d Dictionary)Search( word string) (str string,err error) {

	definition, ok := d[word]
	if !ok {
		return "",ErrNotFound
	}

	return definition, nil
}

