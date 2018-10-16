package station

import (
	"log"
	"sync"
	"time"
)

var (
	//operatingStageNames is an ordered list of stages in the course of a day in station operation
	operatingStageNames = []string{"OPENING", "CLOSING"}

	//pump delays (in seconds) for powering up/down (warmup/cooldown)
	warmup, cooldown = 3, 5
)

func CreateStation(stationName string, pumpCount int, pumpRate, operatingTime time.Duration) *Station {
	gasStation := &Station{
		Name: stationName,
		Pumps: createPumps(pumpCount, pumpRate),
		OperatingTime: operatingTime,
		OperatingStages: map[string]*sync.WaitGroup{},
	}

	//populate OperatingStages WaitGroup's dynamically
	for _, stageName := range operatingStageNames {
		wg := sync.WaitGroup{}
		gasStation.OperatingStages[stageName] = &wg
	}

	return gasStation
}

type Station struct {
	Name string
	Pumps []*Pump
	OperatingTime time.Duration
	OperatingStages map[string]*sync.WaitGroup
	IsOpen bool
}

func(s *Station) Open() {
	s.LogMessage("Station is opening.  Transactions not yet accepted.  Starting up pumps now..")
	s.IsOpen = true

	//Add to wait group
	s.OperatingStages["OPENING"].Add(len(s.Pumps))

	//Turn on pumps
	for _, pump := range s.Pumps {
		go pump.On(warmup, s.OperatingStages["OPENING"])

		//Ping pump to ensure on and listening
		*pump.Input<-"hello"
	}

	s.OperatingStages["OPENING"].Wait()

	s.LogMessage("Station is now OPEN!")
}

func(s *Station) Close() {
	s.LogMessage("Station is closing.  No new transactions accepted.  Shutting down pumps now..")
	s.IsOpen = false

	//Shut off pumps
	s.OperatingStages["CLOSING"].Add(len(s.Pumps))
	for _, pump := range s.Pumps {
		go pump.Off(cooldown, s.OperatingStages["CLOSING"])
	}
	s.OperatingStages["CLOSING"].Wait()

	s.LogMessage("Station is now CLOSED!")
}

func(s *Station) LogMessage(msg string) {
	log.Printf("[STATION %s] %s\n\n", s.Name, msg)
}

func createPumps(pumpCount int, pumpRate time.Duration) []*Pump {
	out := []*Pump{}
	for i := 1; i <= pumpCount; i++ {
		inputChan := make(chan string, 10)
		newPump := &Pump{
			ID: i,
			RatePerGallon: time.Second * pumpRate,
			Input: &inputChan,
		}
		out = append(out, newPump)
	}
	return out
}
