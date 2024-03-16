package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"level_4/model"
	"log"
	"net/http"
)

// AllEmployee = Select Employee API
func AllEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var employee model.Employee
		var response model.ResponseEmployee
		var arrEmployee []model.Employee

		rows, err := db.Query("SELECT employee_id, employee_name, department_id FROM employee")

		if err != nil {
			log.Print(err)
		}

		for rows.Next() {
			err = rows.Scan(&employee.Id, &employee.Name, &employee.Department)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				arrEmployee = append(arrEmployee, employee)
			}
		}

		response.Status = 200
		response.Message = "Success"
		response.Data = arrEmployee

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}

// InsertEmployee = Insert Employee API
func InsertEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.ResponseEmployee

		err := r.ParseMultipartForm(4096)
		if err != nil {
			panic(err)
		}
		employee_name := r.FormValue("employee_name")
		employee_id := r.FormValue("employee_id")
		department_id := r.FormValue("department_id")

		_, err = db.Exec("INSERT INTO employee(employee_name, employee_id, department_id) VALUES(?, ?, ?)", employee_name, employee_id, department_id)

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
