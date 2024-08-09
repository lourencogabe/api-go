package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Document string `json:"document"`
	Address  string `json:"address"`
}

var Users []User
