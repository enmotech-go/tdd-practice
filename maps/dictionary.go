package maps

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

//search from dictionary
func (d Dictionary)Search( word string) (str string,err error) {

	definition, ok := d[word]
	if !ok {
		return "",ErrNotFound
	}

	return definition, nil
}


func (d Dictionary) Add(word, definition string) (err error){
	_, err = d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}


func (d Dictionary) Update(word, definition string) (err error){
	_, err = d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}