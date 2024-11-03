package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lestaat/go-api-server-1/resources"
)

func main() {
	// Initialize the resources package
	if err := resources.Init(); err != nil {
		log.Fatalf("failed to initialize resources package: %v", err)
	}
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

	router.HandleFunc("GET /api/v1/pods", func(w http.ResponseWriter, r *http.Request) {
		namespace := r.URL.Query().Get("namespace")
		if namespace == "" {
			http.Error(w, "Namespace is required", http.StatusBadRequest)
			return
		}

		pods, err := resources.ListPods(namespace)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to list pods: %v", err), http.StatusInternalServerError)
			return
		}

		for _, pod := range pods {
			fmt.Fprintf(w, "Pod Name: %s\n", pod)
		}
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
