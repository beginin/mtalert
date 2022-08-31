package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("<h1>Hello, World</h1>"))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {

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
