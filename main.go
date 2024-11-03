package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/notes", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Accessing API version 1"))
	})

	router.HandleFunc("GET /api/v1/notes/{nid}", func(w http.ResponseWriter, r *http.Request) {
		isStr := r.PathValue("nid")
		nid, err := strconv.Atoi(isStr)
		if err != nil {
			http.Error(w, "Unable to parse Id", http.StatusBadRequest)
			return
		}
		w.Write([]byte(fmt.Sprintf("Accessing API version %s", strconv.Itoa(nid))))
	})

	router.HandleFunc(("GET /healthcheck"), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is up and running"))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server running on port 8080")

	server.ListenAndServe()
}
