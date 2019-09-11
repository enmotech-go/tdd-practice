package class07

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("cannot add word because it already exists")

type Dictionary map[string]string

func (dic Dictionary) Search(keyword string) (result string, err error) {
	result, exists := dic[keyword]
	if !exists {
		err = ErrNotFound
	}
	return
}

func (dic Dictionary) Add(word, definition string) (err error) {
	_, err = dic.Search(word)

	switch err {
	case ErrNotFound:
		dic[word] = definition
		err = nil
	case nil:
		return ErrWordExists
	}
	return
}
