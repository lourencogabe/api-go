package routes

import (
	"log"
	"net/http"

	"github.com/lourencogabe/api-go/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/list/users", controllers.GetUsers)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
