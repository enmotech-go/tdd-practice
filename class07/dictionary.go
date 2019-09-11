package class07

type Dictionary map[string]string

func (dic Dictionary) Search(keyword string) (result string, err error) {
	return dic[keyword], nil
}
