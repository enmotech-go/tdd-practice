package poker

import "os"

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	err = t.file.Truncate(0)
	if err != nil {
		return
	}
	_, err = t.file.Seek(0, 0)
	if err != nil {
		return
	}
	return t.file.Write(p)
}
