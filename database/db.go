package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

)

func ConnectDB() *pgx.Conn {

	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

	connectionString := os.Getenv("DB_CONNECTION_STRING")

	ctx := context.Background()
	conn ,err:= pgx.Connect(ctx,connectionString)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return conn

}

func CloseDB(db *pgx.Conn){
	ctx := context.Background()
	err := db.Close(ctx)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("closing database connection ...")
}

