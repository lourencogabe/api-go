package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Document string `json:"Document"`
	Address  string `json:"Address"`
}

var Users []User
