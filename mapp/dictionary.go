package mapp

import "errors"

var NotExistError = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (dictionary Dictionary) Search(key string) (string, error) {
	value, ok := dictionary[key]
	if !ok {
		return "", NotExistError
	}
	return value, nil
}
