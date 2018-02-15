package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/viant/toolbox"
	"log"
)

//simple echo app
func main() {

	port := os.Args[1];
	if toolbox.AsInt(port) == 0 {
		log.Fatal("echo port")
	}
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
	fmt.Printf("Starting echo app om port " + port + "\n")
	http.ListenAndServe(":" + port, nil)
}
