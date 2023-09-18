package routes

import "net/http"

func InitializeRoutes() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)
	mux.HandleFunc("/exit", closeHandler)

	http.ListenAndServe(":8000", mux)
}
