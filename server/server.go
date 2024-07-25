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


func StartServer(server *http.Server,router *http.ServeMux){

	router.HandleFunc("GET /", Root)
	
	router.HandleFunc("GET /allproducts",AllProducts)
	router.HandleFunc("GET /product/{id}",GetProductByName)
	router.HandleFunc("POST /new",Create)

	router.HandleFunc("GET /randomcard",Random_CC)
	router.HandleFunc("GET /randombook",RandomBook)

	DB = database.ConnectDB()

	if err:= server.ListenAndServe(); err!= nil && !errors.Is(err,http.ErrServerClosed){
		fmt.Println("SERVER ERROR: ",err)
	}

}

func TerminateServer(server *http.Server){
	server.Close()
}
