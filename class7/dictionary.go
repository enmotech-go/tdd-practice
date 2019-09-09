package class7

//Dictionary 自定义的map类型
type Dictionary map[string]string

//Search 自定义类型的方法
func (d Dictionary) Search(word string) (value string) {
	return d[word]
}
