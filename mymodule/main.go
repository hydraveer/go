package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":4000",r))
	
}
func greet() {
	fmt.Println("Hello user")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hi there</h1>"))
}
