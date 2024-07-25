package main

import (
	"fmt"
	"go-sqlc/database"
	"go-sqlc/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	closeSignal := make(chan os.Signal,1)
	signal.Notify(closeSignal,syscall.SIGINT,syscall.SIGTERM)

	router := http.NewServeMux()
	s := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	go server.StartServer(s,router)
	
	<-closeSignal
	fmt.Println("Terminating Server and closing Database Connection")
	server.TerminateServer(s)
	database.CloseDB(server.DB)
}

