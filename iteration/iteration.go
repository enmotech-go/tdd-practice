package iteration

const repeatCount= 5

func Repeat(v string)string{
	repeated:=""
	for i:=0; i<repeatCount; i++{
		repeated+= v
	}
	return repeated
}
