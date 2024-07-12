package server

import (
	"errors"
	"fmt"
	"go-sqlc/database"
	"net/http"

	"github.com/jackc/pgx/v5"
)

var (
	DB *pgx.Conn
)


func StartServer() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Root)
	
	router.HandleFunc("GET /allproducts",AllProducts)
	router.HandleFunc("GET /product/{id}",GetProductByName)
	router.HandleFunc("POST /new",Create)

	router.HandleFunc("GET /randomcard",Random_CC)
	router.HandleFunc("GET /randombook",RandomBook)


	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	DB = database.ConnectDB()

	if err:= server.ListenAndServe(); err!= nil && !errors.Is(err,http.ErrServerClosed){
		fmt.Println("SERVER ERROR: ",err)
	}

}
