package main

import (
	"log"
	"net/http"
	"os"

	"github.com/benacook/environment-monitor-go/controllers"
)

func main() {

	//http.HandleFunc("/", indexHandler)
	controllers.RegisterControllers()

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]
}
