package main

import (
	"net/http"
	"fmt"
)

func main() {


	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("server workin\n")
		fmt.Fprintln(w,"at path '/'")
		
	})

	if err := http.ListenAndServe(":3000",nil);err!=nil{
		fmt.Println("server couldnt listen on port 3000")
	}

}