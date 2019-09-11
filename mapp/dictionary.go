package mapp

import "errors"

type Dictionary map[string]string

func (dictionary Dictionary) Search(key string) (string, error) {
	value, ok := dictionary[key]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return value, nil
}
