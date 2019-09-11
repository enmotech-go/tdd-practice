package mapp

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	NotExistError = DictionaryErr("could not find the word you were looking for")
	ExistError    = DictionaryErr("word already exist")
)

type Dictionary map[string]string

func (dictionary Dictionary) Search(key string) (string, error) {
	value, ok := dictionary[key]
	if !ok {
		return "", NotExistError
	}
	return value, nil
}
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (dictionary Dictionary) Update(word, value string) error {
	_, err := dictionary.Search(word)
	switch err {
	case NotExistError:
		return NotExistError
	case nil:
		dictionary[word] = value
	default:
		return err
	}
	return nil
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
