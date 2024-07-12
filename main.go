package main

import (
	"go-sqlc/database"
	"go-sqlc/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	closeSignal := make(chan os.Signal,1)

	signal.Notify(closeSignal,syscall.SIGINT,syscall.SIGTERM)


	go server.StartServer()
	
	<-closeSignal

	database.CloseDB(server.DB)
}

