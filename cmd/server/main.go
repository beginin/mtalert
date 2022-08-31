package main

import (
	"fmt"
	"net/http"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("<h1>Hello, World</h1>"))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {

	spliturl := strings.Split(r.RequestURI, "/")

	fmt.Println(spliturl[2], spliturl[3], spliturl[4])

}

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/update/", updateHandler)
	mux.HandleFunc("/", homeHandler)
	//server := &http.Server{"127.0.0.1:8080", mux}
	http.ListenAndServe(":8080", mux)
}
