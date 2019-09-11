package mapp

type Dictionary map[string]string

func (dictionary Dictionary) Search(key string) string {
	return dictionary[key]
}
