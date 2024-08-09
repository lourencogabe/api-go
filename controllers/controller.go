package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lourencogabe/api-go/database"
	"github.com/lourencogabe/api-go/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	database.DB.First(&user, id)

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	json.NewDecoder(r.Body).Decode(&newUser)

	database.DB.Create(&newUser)
	log.Print("Usuário Criado com sucesso!")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	database.DB.Delete(&user, id)

	json.NewEncoder(w).Encode(user)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	database.DB.First(&user, id)
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
	log.Print("Usuário Editado com sucesso!")
}
