package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lourencogabe/api-go/controllers"
)

func HandleResquest() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/user/list", controllers.GetUsers)
	mux.HandleFunc("/user/get/{id}", controllers.GetUserId)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
