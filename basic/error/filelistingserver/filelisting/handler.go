package filelisting

import (
	"io"
	"net/http"
	"os"
)

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
