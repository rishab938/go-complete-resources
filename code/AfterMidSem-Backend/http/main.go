package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// struct tags

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}


func main(){

	// Server start 
	fmt.Println("Starting server on port 8090")

	mux:= http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Hello name
		// receive name from Query Params
		// r.Method
		// http.MethodGet
		// w.write
		// w.WriteHeader()

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Only Get methods are allowed")
			return
		}

		name:= r.URL.Query().Get("name")
		if name == "" {
			name = "Stranger"
		}

		fmt.Fprintf(w, "Hello %s!!", name)

		// fmt.Fprintf(w, "Hello world!!")
	})

	// /api/user
	mux.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		// return user object
		// json
		w.Header().Set("Content-Type","application/json")
		user:= User{Id: 1, Name: "Polaris"}
		json.NewEncoder(w).Encode(user)

	})

	// start server
	http.ListenAndServe(":8090", mux)
}