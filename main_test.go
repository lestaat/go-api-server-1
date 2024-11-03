package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetNotes(t *testing.T) {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/notes", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Accessing API version 1"))
	})

	req, err := http.NewRequest("GET", "/api/v1/notes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Accessing API version 1"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetNoteByID(t *testing.T) {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/notes/{nid}", func(w http.ResponseWriter, r *http.Request) {
		isStr := r.URL.Path[len("/api/v1/notes/"):]
		nid, err := strconv.Atoi(isStr)
		if err != nil {
			http.Error(w, "Unable to parse Id", http.StatusBadRequest)
			return
		}
		w.Write([]byte("Accessing API version " + strconv.Itoa(nid)))
	})

	req, err := http.NewRequest("GET", "/api/v1/notes/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Accessing API version 123"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
