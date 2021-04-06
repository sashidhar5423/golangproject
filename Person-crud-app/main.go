package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
}

var persons []Person

func GetPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)

}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, data := range persons {
		if data.Id == params["id"] {
			json.NewEncoder(w).Encode(data)
			return
		}

	}
	json.NewEncoder(w).Encode(&Person{})

}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	persons = append(persons, person)
	json.NewEncoder(w).Encode(persons)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	//router.HandleFunc("/getPerson/{id}", GetPerson).Methods("GET")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, data := range persons {
		if data.Id == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			var person Person
			_ = json.NewDecoder(r.Body).Decode(&person)
			//person.Id = params["id"]
			persons = append(persons, person)
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(persons)

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, data := range persons {
		if data.Id == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)

}
func main() {
	fmt.Println("Main function")
	persons = append(persons, Person{Id: "353", Firstname: "John", Lastname: "Moxley", City: "Bangalore"})
	persons = append(persons, Person{Id: "490", Firstname: "Seth", Lastname: "Rollins", City: "Mumbai"})
	router := mux.NewRouter()
	router.HandleFunc("/getAllPersons", GetPersons).Methods("GET")
	router.HandleFunc("/getPerson/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/createPerson", CreatePerson).Methods("POST")
	router.HandleFunc("/updatePerson/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/deletePerson/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":7071", router))

}
