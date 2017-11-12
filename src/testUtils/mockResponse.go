package testUtils

import (
	"io/ioutil"
	"runtime"

	"path"

	"github.com/jarcoal/httpmock"
)

var _, currentFilePath, _, _ = runtime.Caller(0)
var currentDir = path.Dir(currentFilePath)

func NewMockResponder(status int, template string) httpmock.Responder {
	filepath := path.Join(currentDir, "responses", template+".json")
	jsonByte, _ := ioutil.ReadFile(filepath)

	return httpmock.NewStringResponder(status, string(jsonByte))
}
