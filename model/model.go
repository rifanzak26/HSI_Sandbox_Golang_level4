package model

type Employee struct {
	Id         string `form:"employee_id" json:"employee_id"`
	Name       string `form:"employee_name" json:"employee_name"`
	Department string `form:"department_id" json:"department_id"`
}

type Department struct {
	Id   string `form:"department_id" json:"department_id"`
	Name string `form:"department_name" json:"department_name"`
}

type ResponseEmployee struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Employee
}

type ResponseDepartements struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Department
}
