package routes

import (
	"log"
	"net/http"
)

func InitializeRoutes() {

	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)
	mux.HandleFunc("/exit", closeHandler)

	log.Printf("listening on localhost%v", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
