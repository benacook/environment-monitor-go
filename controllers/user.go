package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/benacook/environment-monitor-go/models"
)

type SensorReadingController struct {
	SensorReadingIDPattern *regexp.Regexp
}

func (uc SensorReadingController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/v1/environment" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if r.URL.Path == "/api/v1/environment/latest" {
		uc.getLatest(w)
	} else {
		matches := uc.SensorReadingIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *SensorReadingController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetSensorReadings(), w)
}

func (uc *SensorReadingController) get(id int, w http.ResponseWriter) {
	u, err := models.GetSensorReadingByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *SensorReadingController) getLatest(w http.ResponseWriter) {
	u, err := models.GetLatestSensorReading()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *SensorReadingController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not parse SensorReading object"))
		return
	}
	u, err = models.AddSensorReading(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *SensorReadingController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not parse SensorReading object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted SensorReading must match ID in URL"))
		return
	}
	u, err = models.UpdateSensorReading(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *SensorReadingController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveSensorReadingByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *SensorReadingController) parseRequest(r *http.Request) (models.SensorReading, error) {
	dec := json.NewDecoder(r.Body)
	var u models.SensorReading
	err := dec.Decode(&u)
	if err != nil {
		return models.SensorReading{}, err
	}
	return u, nil
}

func newSensorReadingController() *SensorReadingController {
	return &SensorReadingController{
		SensorReadingIDPattern: regexp.MustCompile(`^/api/v1/environment/(\d+)/?`),
	}
}
