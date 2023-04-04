package endpoint

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	UserId int    `json:"userId"`
	PostId int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Make an HTTP GET request to the JSON endpoint
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Printf("Failed to fetch JSON data: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body and unmarshall into struct
	var data []Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Write the struct to ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetSinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Make an HTTP GET request to the JSON endpoint
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + vars["id"])
	if err != nil {
		fmt.Printf("Failed to fetch JSON data: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Write the struct to ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HttpExample() {
	port := ":3001"
	r := mux.NewRouter()
	r.HandleFunc("/posts", GetPosts)
	r.HandleFunc("/post/{id}", GetSinglePost)
	http.Handle("/", r)
	log.Printf("Server starting on %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
