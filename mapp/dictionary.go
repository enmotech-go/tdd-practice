package mapp

import "errors"

var NotExistError = errors.New("could not find the word you were looking for")
var ExistError = errors.New("word already exist")

type Dictionary map[string]string

func (dictionary Dictionary) Search(key string) (string, error) {
	value, ok := dictionary[key]
	if !ok {
		return "", NotExistError
	}
	return value, nil
}

func (dictionary Dictionary) Add(key, value string) error {
	_, err := dictionary.Search(key)
	switch err {
	case NotExistError:
		dictionary[key] = value
	case nil:
		return ExistError
	default:
		return err
	}
	return nil
}
