package main

import (
	"fmt"
	"net/http"

	"github.com/Alfred-Onuada/go-dropbox/internals/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	address := "localhost"
	port := 4001
	url := fmt.Sprintf("%v:%v", address, port)

	fmt.Printf("Go server running @ %v \n", url)

	err := http.ListenAndServe(url, mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
