package main

import (
	"fmt"
	"time"

	"go-gas/station"
)

const (
	version  = "v0.0.1"
	overview = "This app simulates a gas station's operations for the day as they occur concurrently:\nopen, close, customer cars arriving/pumping/leaving, etc."

	//station args
	stationName   = "Gogas"
	pumpCount     = 4
	pumpRate      = time.Second * 5
	operatingTime = time.Second * 10
)

var (
	mainStation *station.Station
)

func main() {
	fmt.Printf("\n*************\nGo Gas %s\n*************\n%s\n\n", version, overview)

	//Create gas station
	mainStation = station.CreateStation(stationName, pumpCount, pumpRate, operatingTime)

	//Open gas station
	mainStation.Open()

	//Start timer that expires at the end of operating hours
	stationTimer := time.NewTimer(mainStation.OperatingTime)

	//Wait for end of operating hours
	<-stationTimer.C

	mainStation.LogMessage(">>ANNOUNCEMENT<<\n\n***\n***\n\nOperating hours are over!\n\n***\n***")

	//Begin station closing, all routines must complete before station is officially closed
	mainStation.Close()
}
