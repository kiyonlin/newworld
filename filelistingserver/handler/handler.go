package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println(request.URL.Path)
	if !strings.HasPrefix(request.URL.Path, prefix) {
		return userError(fmt.Sprintf("path %s must start with %s", request.URL.Path, prefix))
	}
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}

type userError string

func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}
