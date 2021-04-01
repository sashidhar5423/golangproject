package main

import (
	//empDao "employee-app/dao"
	dbCon "employee-app/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("from main function")
	// init router
	r := mux.NewRouter()
	dbCon.GetDbconnection()

	r.HandleFunc("/getAll", dbCon.GetEmployees).Methods("GET")
	r.HandleFunc("/getEmployee/{id}", dbCon.GetEmployee).Methods("GET")
	r.HandleFunc("/create", dbCon.CreateEmployees).Methods("POST")
	r.HandleFunc("/update/{id}", dbCon.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/delete/{id}", dbCon.DeleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9095", r))

}
