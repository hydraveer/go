package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hydraveer/workingwithmongo/routes"
)

func main() {
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening on port 40000")
}
