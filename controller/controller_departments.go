package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"level_4/model"
	"log"
	"net/http"
)

func HandleAllDepartments(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var department model.Department
		var response model.ResponseDepartements
		var arrdepartment []model.Department

		rows, err := db.Query("SELECT department_id, department_name FROM departments")

		if err != nil {
			log.Print(err)
		}

		for rows.Next() {
			err = rows.Scan(&department.Id, &department.Name)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				arrdepartment = append(arrdepartment, department)
			}
		}

		response.Status = 200
		response.Message = "Success"
		response.Data = arrdepartment

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}

// InsertEmployee = Insert Employee API
func InsertDepartments(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.ResponseDepartements
		err := r.ParseMultipartForm(4096)
		if err != nil {
			panic(err)
		}
		department_id := r.FormValue("department_id")
		department_name := r.FormValue("department_name")

		_, err = db.Exec("INSERT INTO departments (department_id, department_name) VALUES(?, ?)", department_id, department_name)

		if err != nil {
			log.Print(err)
			return
		}
		response.Status = 200
		response.Message = "Insert data successfully"
		fmt.Print("Insert data to database")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}
