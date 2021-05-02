package expnet

import (
	"log"
	"net/http"
	"testing"
)

func TestMiddleware(t *testing.T) {
	//http.HandleFunc("/", mainHandler)
	http.HandleFunc("/", firstMiddleware(mainHandler))

	http.ListenAndServe(":10101", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("main Handler")
	w.WriteHeader(http.StatusOK)
}

func firstMiddleware(next http.HandlerFunc)  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("first Middleware")
		next(w, r)
	}
}


