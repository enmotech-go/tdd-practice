package iteration

/**
class_03
demand: 编写重复字符5次的函数，单元测试、基准测试
*/

func Repeat(character string) string  {
	var repeated string
	count := 5
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}