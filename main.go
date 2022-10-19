package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/campeonatos", getCampeonatos).Methods("GET")
	r.HandleFunc("/", home).Methods("GET")

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func getCampeonatos(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://apijogos.herokuapp.com/campeonatos")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	json.NewEncoder(w).Encode(string(responseData))
}
