package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) { resp.Write([]byte("home page")) })
	fmt.Println("server start on port 8080")
	s := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	log.Fatal(s.ListenAndServe())
}
