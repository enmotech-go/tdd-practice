package class7_1

//Dictionary 自定义一个map类型
type Dictionary map[string]string

//Search Search
func (d Dictionary) Search(word string) string {
	return d[word]
}
