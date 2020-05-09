package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

//RegisterControllers blah
func RegisterControllers(){
	uc := newSensorReadingController()

	http.Handle("/api/v1/environment/", *uc)
	http.Handle("/api/v1/environment", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer){
	enc := json.NewEncoder(w)
	enc.Encode(data)
}