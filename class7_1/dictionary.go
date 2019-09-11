package class7_1

import "errors"

//Dictionary 自定义一个map类型
type Dictionary map[string]string

//Search Search
func (d Dictionary) Search(word string) (string, error) {
	v, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return v, nil
}
