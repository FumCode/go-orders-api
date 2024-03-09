package main

import (
	"fmt"
	"net/http"

	
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// init server
	server := &http.Server{
		Addr: ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}
	// check for errors
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("faided", err)
	}
}
// basic handler function
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Server"))
}