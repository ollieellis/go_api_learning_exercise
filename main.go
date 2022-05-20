package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dumy_database map[string]int

type fruit_entry struct { //used for the put (currently think this is bad but idk)
	Name   string
	Number int
}

func add_fruit_and_quantity(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var new_entry fruit_entry
	json.Unmarshal(reqBody, &new_entry)
	dumy_database[new_entry.Name] = new_entry.Number
}

func get_num_of_fruit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	pathParams := mux.Vars(r)
	fruit := pathParams["name"]
	json.NewEncoder(w).Encode(dumy_database[fruit])
	//you may print to terimnal here to inspect recieved paramaters ect
}

func handle_requests() {
	r := mux.NewRouter()
	r.HandleFunc("/fruit/{name}", get_num_of_fruit).Methods(http.MethodGet)
	r.HandleFunc("/new_fruit/", add_fruit_and_quantity).Methods(http.MethodPut)
	r.HandleFunc("/print_db", print_db).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe(":8080", r))
}

func main() {
	dumy_database = make(map[string]int)
	dumy_database["apples"] = 3
	dumy_database["bananas"] = 2
	fmt.Println(dumy_database)
	fmt.Println("running")

	handle_requests()
}

func print_db(w http.ResponseWriter, r *http.Request) { //this should not be here but for learning purposes only
	w.Header().Set("Content-Type", "application/json") 
	w.WriteHeader(http.StatusOK)
	fmt.Print(dumy_database) //prints to terminal!
}
