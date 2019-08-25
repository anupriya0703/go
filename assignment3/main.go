package main

import (
	"log"
	"final_assignment/assignment3/router"
	"net/http"
)

func main() {
	log.Println("Starting......!!")
	router.Routers()
	http.ListenAndServe(":8080", router.Mux)
	log.Println("DONE!")
}
