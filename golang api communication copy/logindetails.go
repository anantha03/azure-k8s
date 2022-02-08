package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getdetails", getDetails).Methods("GET")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	handler := c.Handler(r)
	srv := &http.Server{Handler: handler, Addr: ":4000"}
	log.Fatal(srv.ListenAndServe())

}

type Data struct {
	Title string `json: "title"`
}

func getDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	title := Data{"login details"}
	//user := UserData{"d", "d", "d"}
	json.NewEncoder(w).Encode(&title)

}
