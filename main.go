package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	details "github.com/itisaby/go-microservice/details"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello, %s!", name)
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("healthHandler")

	response := map[string]string{
		"status": "UP",
		"time":   time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving Home PAge")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World, App is up and running")

}
func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving details Page")
	hostname, err := details.GetHostName()
	if err != nil {
		panic(err)
	}
	IP, err := details.GetIP()

	fmt.Println(hostname, IP)
	response := map[string]string{
		"hostname": hostname,
		"ip":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", handler).Methods("GET")
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	log.Println("Server is up and running on port 80")

	log.Fatal(http.ListenAndServe(":80", r))
}
