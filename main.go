package main

import (
	"net/http"

	"github.com/lcantelli/golang-boilerplate-project/controllers"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
