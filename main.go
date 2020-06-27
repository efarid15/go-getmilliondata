package main

import (
	"goglobalrank/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/domainrank", handlers.GetDomainrank)

	println("Listening Port :9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func home(w http.ResponseWriter, r *http.Request)  {
	println("Get Million Data")
}

