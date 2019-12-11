package poker

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)
	got := string(newFileContents)
	want := "abc"
	assert.Equal(t, want, got)
}
