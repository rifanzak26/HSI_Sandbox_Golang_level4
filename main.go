package main

import (
	"fmt"
	"level_4/config"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"level_4/controller"
)

func main() {
	db := config.Connect()
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/getEmployee", controller.AllEmployee(db)).Methods("GET")
	router.HandleFunc("/insertEmployee", controller.InsertEmployee(db)).Methods("POST")
	router.HandleFunc("/getDepartments", controller.HandleAllDepartments(db)).Methods("GET")
	router.HandleFunc("/insertDepartments", controller.InsertDepartments(db)).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
