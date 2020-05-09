package main

import (
	"net/http"

	"github.com/benacook/environment-monitor-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
