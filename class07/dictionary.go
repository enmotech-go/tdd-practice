package class07

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (dic Dictionary) Search(keyword string) (result string, err error) {
	result, exists := dic[keyword]
	if !exists {
		err = ErrNotFound
	}
	return
}

func (dic Dictionary) Add(word, definition string) {
	dic[word] = definition
}
