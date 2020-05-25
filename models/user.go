package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/benacook/environment-monitor-go/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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
		return SensorReading{},
			errors.New("New SensorReading must not include ID, or it must be zero")
	}
	u.ID = nextID
	nextID++
	SensorReadings = append(SensorReadings, &u)
	err := LogSensorReading(u)
	if err != nil {
		log.Fatal(err)
		return u, err
	}
	return u, nil
}

func LogSensorReading(u SensorReading) error {
	mdb := mongoDB.Mongodb{}
	err := mdb.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}

	mdb.InitDatabase()

	err = mdb.InsertElement(u)
	if err != nil {
		return err
	}
	mdb.Client.Disconnect(mdb.Context)
	return nil
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
	GetSensorReading()
	for _, u := range SensorReadings {
		if u.ID == (nextID - 1) {
			return *u, nil
		}
	}
	return SensorReading{}, fmt.Errorf("no readings yet")
}

func GetSensorReading() error {
	mdb := mongoDB.Mongodb{}
	err := mdb.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}

	mdb.InitDatabase()

	var result SensorReading
	err = mdb.Collection.FindOne(context.TODO(), bson.D{}).Decode(&result)
	if err != nil {
		return err
	}
	fmt.Println(result)
	mdb.Client.Disconnect(mdb.Context)
	return nil
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
