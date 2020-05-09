package models

import (
	"errors"
	"fmt"
)

//SensorReading data type
type SensorReading struct {
	ID          int
	Temperature float64
	Humidity    float64
}

var (
	SensorReadings []*SensorReading
	nextID         = 1
)

//GetSensorReadings returns the array of SensorReadings
func GetSensorReadings() []*SensorReading {
	return SensorReadings
}

//AddSensorReading blah
func AddSensorReading(u SensorReading) (SensorReading, error) {
	if u.ID != 0 {
		return SensorReading{}, errors.New("New SensorReading must not include ID, or it must be zero")
	}
	u.ID = nextID
	nextID++
	SensorReadings = append(SensorReadings, &u)
	return u, nil
}

//GetSensorReadingByID blah
func GetSensorReadingByID(id int) (SensorReading, error) {
	for _, u := range SensorReadings {
		if u.ID == id {
			return *u, nil
		}
	}
	return SensorReading{}, fmt.Errorf("SensorReading with ID %v not found", id)
}

//GetLatestSensorReading blah
func GetLatestSensorReading() (SensorReading, error) {
	for _, u := range SensorReadings {
		if u.ID == (nextID - 1) {
			return *u, nil
		}
	}
	return SensorReading{}, fmt.Errorf("no readings yet")
}

//UpdateSensorReading blah
func UpdateSensorReading(u SensorReading) (SensorReading, error) {
	for i, candidate := range SensorReadings {
		if candidate.ID == u.ID {
			SensorReadings[i] = &u
			return u, nil
		}
	}
	return SensorReading{}, fmt.Errorf("SensorReading with ID %v not found", u.ID)
}

//RemoveSensorReadingByID blah
func RemoveSensorReadingByID(id int) error {
	for i, u := range SensorReadings {
		if u.ID == id {
			SensorReadings = append(SensorReadings[:i], SensorReadings[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("SensorReading with ID %v not found", id)
}
