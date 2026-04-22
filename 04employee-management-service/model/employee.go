package model

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}
