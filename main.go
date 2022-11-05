package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox."))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with id %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet."))
}
