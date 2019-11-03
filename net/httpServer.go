package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/",
		func(writer http.ResponseWriter,
			request *http.Request) {
			s := request.URL.Path[len("/list/"):]
			file, err := os.Open(s)
			if err != nil {
				panic(err)
			}
			data, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			_, err = writer.Write(data)
			if err != nil {
				panic(err)
			}
		})
	_ = http.ListenAndServe(":8888", nil)
}
