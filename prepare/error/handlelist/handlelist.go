package handlelist

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandelFileList(writer http.ResponseWriter, request *http.Request) error {
	//打开文件
	fileName := request.URL.Path[len("/list/"):]
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	read := bufio.NewReader(file)
	defer read.Buffered()

	data, err := ioutil.ReadAll(read)
	if err != nil {
		return err
	}
	//log.Fatal(data)
	fmt.Fprintln(writer, string(data))

	return nil
}
