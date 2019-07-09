package main

import (
	"bytes"
	"fmt"
	"github.com/ega-forever/otus_go/api"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCopy(t *testing.T) {

	text := bytes.NewBuffer([]byte("super test 123 test 222"))
	filename, _ := ioutil.TempFile("", "test")
	filenameCopy, _ := ioutil.TempFile("", "test")
	_ = ioutil.WriteFile(filename.Name(), text.Bytes(), 0777)
	api.Copy(filename.Name(), filenameCopy.Name(), 2, 12)

	textCopy, _ := ioutil.ReadFile(filenameCopy.Name())
	textCopyString := bytes.NewBuffer(textCopy).String()
	textString := bytes.NewBuffer(text.Bytes()[2:14]).String()

	fmt.Println(strings.Compare(textCopyString, textString))
	fmt.Println(textCopyString == textString)

	_ = os.Remove(filename.Name())
	_ = os.Remove(filenameCopy.Name())

	if textCopyString != textString {
		t.Error("copy util doesn't work")
	}
}
