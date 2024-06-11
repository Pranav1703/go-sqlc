package api

import (
	"net/http"
)

func StartServer() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Root)
	router.HandleFunc("GET /randomcard",Random_CC)
	router.HandleFunc("GET /randomBook",RandomBook)


	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	
	server.ListenAndServe()
}
