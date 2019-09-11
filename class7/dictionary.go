package class7

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

//Dictionary 自定义的map类型
type Dictionary map[string]string

//Search 自定义类型的方法
func (d Dictionary) Search(word string) (value string, err error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return
}
