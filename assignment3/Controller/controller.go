package controller

import (
	"fmt"
	"net/http"
	
	"github.com/anupriya0703/go/assignment3/domain"
	"github.com/anupriya0703/go/assignment3/mapstore"
	"github.com/gorilla/mux"
)

type CustomerController struct {
	store domain.CustomerStore
}

var controller = CustomerController{
	store: mapstore.NewMapStore(),
}

func Create(w http.ResponseWriter, r *http.Request) {

	    cust := domain.Customer{
		ID:   "1605344",
		Name: "Anupriya Gupta",
		Email:"anupriya.gupta@mayadata.io",
	}

	err := controller.store.Create(cust)
	if err != nil {
		fmt.Println("Could not create the customer!!")
		}else {
		fmt.Println("Customer  created successfuly!!")
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	    updateCustomer := domain.Customer{
		ID:   "160",
		Name: "Anupriya Gupta",
		Email:"anu.gupta@mayadata.io",
	}
	err := controller.store.Update("1605344", updateCustomer)
	if err != nil {
		fmt.Println("Update Failed!Try again")
	}else {
		fmt.Println("Deatils Updated", updateCustomer)
	}

}
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	del := vars["1605344"]
	err := controller.store.Delete(del)
	if err != nil {
		fmt.Println("Deletion Failed")

	}else {
		fmt.Println("DELETED!!")
	}

}

func GetById(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	k := vars["1605344"]
	//fmt.Fprintf(w,"GetbyID ")
	value, err := controller.store.GetById(k)
	if err != nil {
		fmt.Println("Incorrect!check Id ")

	}else {
		fmt.Println(value)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Getall executed")
	  value, err := controller.store.GetAll()
	if err != nil {
		fmt.Println("ERROR!TRY AGAIN")

	}else {
		fmt.Println( value)
	}

}
