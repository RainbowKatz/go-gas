package main

import (
	"fmt"
	"sync"
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

	//Open gas station with a wait group that keeps station open
	var wg sync.WaitGroup
	mainStation.Open(&wg)

	//Wait for all station go routines to complete (i.e. station timer, pump inputs, etc.)
	wg.Wait()

	mainStation.LogMessage("Station is now CLOSED!")
}