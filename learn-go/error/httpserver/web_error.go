package main

import (
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"strings"
)

/*
	服务器统一错误处理
 */

type userError interface{
	error
	Message() string
}

type UserError string

func (e UserError) Error() string {
	return e.Message()
}

func (e UserError) Message() string {
	return string(e)
}

const req_prefix = "/list/"
const file_prefix = "./data/"

func HandlerFileList(writer http.ResponseWriter, request *http.Request) (err error) {
	if strings.Index(request.URL.Path, req_prefix) != 0 {
		return UserError("path must start with" + req_prefix)
	}
	
	path := request.URL.Path[len(req_prefix):]
	file, err := os.Open(file_prefix + path)
	if err != nil {
		return
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	
	writer.Write(all)
	return
}

type Handler func(writer http.ResponseWriter, request *http.Request) (err error)

func errWrapper(handler Handler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occurred handling request: %s", err.Error())

			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
				case os.IsNotExist(err):
					code = http.StatusNotFound
				case os.IsPermission(err):
					code = http.StatusForbidden
				default:
					code = http.StatusInternalServerError 
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(HandlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}