package main

import (
	"fmt"
	"log"
	"net/http"
)

//corsWrapper wraps handler enabling cross-origin calls
func corsWrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")

		http.HandlerFunc(f).ServeHTTP(w, r)

	}
}

// GetRealIPHandler gets real ip using X-Real-IP header provided by proxy
func GetRealIPHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "%s", req.Header.Get("X-Real-IP"))
}

func main() {
	http.Handle("/", corsWrapper(GetRealIPHandler))
	log.Println("GetRealIPHandler listening ...")
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		panic(err)
	}

}
