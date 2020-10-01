package app

import (
	"net/http"
	"github.com/utieyin/golang-micro/tree/master/controllers"
)

// StartApp provides an entry point for the user endpoint
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser())
}
