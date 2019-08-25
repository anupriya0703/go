package main

import "final_assignment/assignment_2/domain"
import "final_assignment/assignment_2/mapstore"
//import "errors"
import"fmt"

type CustomerController struct {

	store domain.CustomerStore
}


func (c1 CustomerController ) Add (c domain.Customer) {
	err:= c1.store.Create(c)
	if err!=nil {
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (c1 CustomerController)Update(s string,c domain.Customer){
	e:=c1.store.Update(s,c)
	if e!=nil{
		fmt.Println(e)
	}else{
		fmt.Println("Updated !!")
	}
}


func(c1 CustomerController)Getall(){
	value,err:=c1.store.GetAll()
	if err!=nil{
	fmt.Println("ERROR")
	}else {
		fmt.Println(value)
	}
}


func(c1 CustomerController)GetBy(){
	var id string
	fmt.Println("Enter the key")
	fmt.Scanln(&id)
	value,err:=c1.store.GetById(id)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(value)
	}
}

func (c1 CustomerController)Delete(){
	err:=c1.store.Delete("A")
	if err!=nil{
		fmt.Println("DELETION UNSUCCESSFUL,error")
	}else {
		fmt.Println("Successful")
	}
}
func main(){

	controller := CustomerController{
		store : mapstore.NewMapStore(),
		}

	customer := domain.Customer {
	     	      ID : "Anu",
	     	      Name: "Anupriya",
	     	      Email :"anupriya.gupta@mayadata.io",
	                     }
	controller.Add(customer)
	controller.Update("A",customer)
	controller.Getall()
	controller.GetBy()
	controller.Delete()
}

