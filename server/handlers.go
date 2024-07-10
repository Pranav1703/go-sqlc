package server

import (
	"HttpServer/database/produtcs"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"At path '/' .")
}

func AllProducts(w http.ResponseWriter, r *http.Request ) {
	ctx := context.Background()

	queries := products.New(DB)

	productList , err := queries.GetAllProducts(ctx) 
	if err!=nil{
		fmt.Println(err)
		json.NewEncoder(w).Encode(fmt.Sprintf("%v",err))
	}else{

		json.NewEncoder(w).Encode(productList)
	}

}

func GetProductByName(w http.ResponseWriter, r *http.Request ){

	idValue := r.PathValue("id")

	id ,err := strconv.Atoi(idValue)

	if err!= nil{
		fmt.Println(err)
	}

	ctx := context.Background()
	queries := products.New(DB)

	product, err := queries.GetProduct(ctx,int32(id))

	if err!= nil{
		fmt.Println(err)
		json.NewEncoder(w).Encode(fmt.Sprintf("%v",err))
	}else{
		json.NewEncoder(w).Encode(product)
	}

}