package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreditCardInfo struct {
	Type 		string 
	Number 		string
	Expiration 	string
	Owner 		string
}	

type BookInfo struct {
	Title 		string
	Author 		string
	Genre   	string
	Description string
	Isbn        string
	Published 	string
	Publisher 	string
}

func Random_CC(w http.ResponseWriter, r *http.Request) {
	resp,err := http.Get("https://fakerapi.it/api/v1/credit_cards?_quantity=1")
	
	type apiResponse struct {
		Data []CreditCardInfo
	}

	apiResp := &apiResponse{}

	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w,"Couldnt fetch")
	}
	
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(apiResp)
	
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Print("Couldnt parse response")
	}

	cardDetails,err := json.Marshal(apiResp)
	
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Print("Couldnt parse response")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(cardDetails)
}

func RandomBook(w http.ResponseWriter, r *http.Request) {

	resp,err := http.Get("https://fakerapi.it/api/v1/books?_quantity=1")

	if err!= nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w,"Couldnt fetch")
	}

	defer resp.Body.Close()

	type apiResponse struct {
		Data []BookInfo
	}
	apiResp := &apiResponse{}

	err = json.NewDecoder(resp.Body).Decode(apiResp)

	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Print("Couldnt parse response")
	}
	
	Book,err := json.Marshal(apiResp.Data)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Print("Couldnt parse response")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(Book)
}