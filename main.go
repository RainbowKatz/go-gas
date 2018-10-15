package main

import (
	"fmt"
	"time"

	"github.com/RainbowKatz/go-gas/station"
)

const (
	version = "v0.0.1"
	overview = "This app simulates a gas station's operations for the day as they occur concurrently:\nopen, close, customer cars arriving/pumping/leaving, etc."
	
	//station args
	stationName = "Gogas"
	pumpCount = 4
	pumpRate = time.Second * 5
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

	//Wait for last operating stage to complete before station if officially closed
	lastOpStageIdx := len(station.OperatingStageNames)-1
	mainStation.OperatingStages[station.OperatingStageNames[lastOpStageIdx]].Wait()

	mainStation.LogMessage("Station is now CLOSED!")
}