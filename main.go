package main

import (
	"log"
	"net/http"
	"os"

	"github.com/benacook/environment-monitor-go/controllers"
)

func main() {
		
	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
			port = "3000"
			log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal(err)
	}
	// [END setting_port]

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
