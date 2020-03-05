package main

import (
	"log"
	"github.com/anupriya0703/go/assignment3/router"
	"net/http"
)

func main() {"github.com/anupriya0703/go/assignment3/"
5
        "github.com/anupriya0703/go/assignment3/mapstore"
	log.Println("Starting......!!")
	router.Routers()
	http.ListenAndServe(":8080", router.Mux)
	log.Println("DONE!")
}
