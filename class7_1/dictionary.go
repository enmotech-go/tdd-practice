package class7_1

import "errors"

//Dictionary 自定义一个map类型
type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("cannot add word because it already exists")

//Search Dictionary Search
func (d Dictionary) Search(word string) (string, error) {
	v, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

//Add   Dictionary Add
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
