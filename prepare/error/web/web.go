package main

import (
	"goCron/prepare/error/handlelist"
	"net/http"
	"os"
)

type appHandle func(writer http.ResponseWriter, request *http.Request) error

func errorHandle(handle appHandle) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := handle(writer, request); err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusBadRequest
			}
			http.Error(writer, http.StatusText(code), code)
		}

	}
}

func main() {
	http.HandleFunc("/list/", errorHandle(handlelist.HandelFileList))
	http.ListenAndServe(":8080", nil)
}
