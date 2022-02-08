package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var userData Data

type accountName struct {
	Name string `json: "AccountName"`
}
type Data struct {
	Title string `json: "title"`
}

func main() {
	fmt.Println("connected")
	//	printName()
	r := mux.NewRouter()
	r.HandleFunc("/getdetails", getDetails).Methods("GET")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	handler := c.Handler(r)
	srv := &http.Server{Handler: handler, Addr: ":3000"}
	log.Fatal(srv.ListenAndServe())

}
func printName() {
	response, err := http.Get("http://backend-deploy-service.default.svc.cluster.local/getdetails")
	//response, err := http.Get("http://localhost:4000/getdetails")
	if err != nil {
		fmt.Println(err)
	}

	_ = json.NewDecoder(response.Body).Decode(&userData)
	fmt.Println(userData.Title)
}
func getDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get details")
	w.Header().Set("Content-Type", "application/json")

	response, err := http.Get("http://backend-deploy-service.default.svc.cluster.local/getdetails")
	if err != nil {
		fmt.Println(err)
	}

	_ = json.NewDecoder(response.Body).Decode(&userData)
	//user := UserData{"d", "d", "d"}
	fmt.Println(userData.Title)
	json.NewEncoder(w).Encode(&userData)

}
