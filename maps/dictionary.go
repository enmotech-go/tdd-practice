package maps

type Dictionary map[string]string

//search from dictionary
func (d Dictionary)Search( word string) string {

	return d[word]
}

