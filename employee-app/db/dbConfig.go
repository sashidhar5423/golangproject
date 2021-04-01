package db

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	//"gorm.io/gorm"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//var Db *sql.DB
var dbc *gorm.DB
var err error

type Employee struct {
	gorm.Model
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Age    uint   `json:"age"`
	Salary uint   `json:"salary"`
}

func GetDbconnection() {
	fmt.Println("db connection")
	dbName := "root:raju5423a@tcp(localhost:3306)/schema1?charset=utf8mb4&p"

	dbc, err = gorm.Open("mysql", dbName)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}

	dbc.AutoMigrate(&Employee{})
	fmt.Println("Database connected")
	//defer Db.Close()

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all employees")
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	dbc.Find(&employees)
	json.NewEncoder(w).Encode(employees)

}
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get employee")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	empId := params["id"]
	var emp Employee
	dbc.Find(&emp, empId)
	json.NewEncoder(w).Encode(emp)

}
func CreateEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create employee")
	w.Header().Set("Content-Type", "application/json")
	var createEmp Employee
	json.NewDecoder(r.Body).Decode(&createEmp)
	dbc.Create(&createEmp)
	json.NewEncoder(w).Encode(createEmp)
}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update employee")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	empId := params["id"]
	var updateEmp Employee
	dbc.First(&updateEmp, empId)
	json.NewDecoder(r.Body).Decode(&updateEmp)
	dbc.Save(&updateEmp)
	json.NewEncoder(w).Encode(updateEmp)
}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete employee")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var deleteEmp Employee
	dbc.Delete(&deleteEmp, params["id"])
	json.NewEncoder(w).Encode(deleteEmp)

}
