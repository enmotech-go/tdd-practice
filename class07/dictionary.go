package class07

type Dictionary map[string]string

func (dic Dictionary) Search(keyword string) string {
	return dic[keyword]
}
