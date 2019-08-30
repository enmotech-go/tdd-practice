package iteration

func Repeat(v string)string{
	repeated:=""
	for i:=0; i<5; i++{
		repeated+= v
	}
	return repeated
}
