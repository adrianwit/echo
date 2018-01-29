package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

//simple echo app
func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		if request.ContentLength > 0 {
			if content, err := ioutil.ReadAll(request.Body);err == nil {
				if len(content) > 0 {
					writer.Write(content)
				}
			}
		}
	})
	fmt.Printf("Starting echo app")
	http.ListenAndServe(":8080", nil)
}
