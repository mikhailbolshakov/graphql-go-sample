package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	personGql, err := NewPersonGql()
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}

	http.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query().Get("query")
		log.Printf("Handle request. Query: %s", query)

		result, err := personGql.ExecQuery(query)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			json.NewEncoder(w).Encode(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(result)
		}
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}