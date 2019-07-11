package main

import (
	"log"
	"net/http"
	"os"

	"./objects"
)

//绑定接口并运行
func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
