package router

import(
	"github.com/anupriya0703/go/assignment3/Controller"
	"github.com/gorilla/mux"
)


var Mux = mux.NewRouter()


func Routers() {

	Mux.HandleFunc("/create", controller.Create).Methods("GET")
	Mux.HandleFunc("/update", controller.Update).Methods("GET")
	Mux.HandleFunc("/getbyid", controller.GetById).Methods("GET")
	Mux.HandleFunc("/getall", controller.GetAll).Methods("GET")
	Mux.HandleFunc("/delete", controller.Delete).Methods("GET")

}
