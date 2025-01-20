package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hydraveer/restapi/internal/config"
)

func main() {
	//load conf
	cfg := config.MustLoad()
	//database setup
	//setup route
	r := http.NewServeMux()
	r.HandleFunc("GET /", ServerHome)
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: r,
	}

	fmt.Println("Server started")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}


	
	//setup server
}
func ServerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there"))
	r.Body.Close() 
}
