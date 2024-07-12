package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {

	ctx := context.Background()
	conn ,err:= pgx.Connect(ctx,"postgres://postgres:popcat@localhost:5432/sqlctest")
	fmt.Println("Connected to database")
	if err!=nil{
		log.Fatal(err)
	}

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

