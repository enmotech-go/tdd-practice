package maps

const (
	ErrNotFound = DictionaryErr("not found")
	ErrExist    = DictionaryErr("exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, def string) {
	d[word] = def
}

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
