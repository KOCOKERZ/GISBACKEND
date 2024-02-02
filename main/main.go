package main

import (
	"fmt"
	"net/http"

	akbar "github.com/KOCOKERZ/GISBACKEND"
)

func main() {
	http.HandleFunc("/", HelloHTTP)
	http.ListenAndServe(":3000", nil)
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, akbar.PostNear("mongoenv", "GIS", "GIS", r))
}

func GetToken(r *http.Request) string {
	return r.Header.Get("Authorization")
}
