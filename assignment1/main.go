package main

import (
	"fmt"
	)

func Add(m map[string]string){

	var key,value string
	//Taking input
	fmt.Println("Enter the key to be added")
	fmt.Scanln(&key)
	fmt.Println("Enter the value to be added")
	fmt.Scanln(&value)

	//

    m[key]=value

}


func Update(key string,m map[string]string)(string){
	 var value string
	fmt.Println("Enter the value ")
	fmt.Scanln(&value)
	m[key]=value
	return ("Updated Successfully!!!")

}

func Display(m map[string]string){
	fmt.Println(m)
}

func Get(m map[string]string){
	var key string
	fmt.Println("Enter the key")
	fmt.Scanln(&key)
	if value, ok := m[key]; ok {
		fmt.Println("value: ", value)
	} else {
		fmt.Println("key not found")
	}
}
func Delete(m map[string]string){
 var key string
 fmt.Println("Enter the entry to be deleted")
 fmt.Scanln(&key)
 delete(m,key)
 fmt.Print("Deleted entry!!")
}

func main(){

	var m map[string]string
	m=make(map[string]string)

	var key1 string
	//To add an entry
	Add(m)
	Add(m)
	//To Update
	fmt.Println("Enter the ey to be updated")
	fmt.Scanln(&key1)
	fmt.Println(Update(key1,m))
     //GET
     Get(m)
    Display(m)
	Delete(m)
}
