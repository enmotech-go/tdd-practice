package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()
	tape := &tape{file}
	_, err := tape.Write([]byte("abc"))
	assert.NoError(t, err)
	_, err = file.Seek(0, 0)
	assert.NoError(t, err)
	newFileContents, _ := ioutil.ReadAll(file)
	got := string(newFileContents)
	want := "abc"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
