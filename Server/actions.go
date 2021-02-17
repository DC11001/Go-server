package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Model struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func Response(w http.ResponseWriter, status int, results Model) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func PostFunction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var info Model
	err := decoder.Decode(&info)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	Response(w, 200, info)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var_id := params["id"]
	fmt.Fprintf(w, "You asked for the id: "+var_id)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data Model
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	log.Println("Removed:\n", data)
	fmt.Fprintf(w, "Item removed")
}
