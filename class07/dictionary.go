package class07

import "errors"

type Dictionary map[string]string

func (dic Dictionary) Search(keyword string) (result string, err error) {
	result, exists := dic[keyword]
	if !exists {
		err = errors.New("could not find the word you were looking for")
	}
	return
}
