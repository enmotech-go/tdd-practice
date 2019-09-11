package class7_1

import "errors"

//Dictionary 自定义一个map类型
type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

//Search Search
func (d Dictionary) Search(word string) (string, error) {
	v, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
