package station

import (
	"log"
	"sync"
	"time"
)

var (
	//pump delays (in seconds) for powering up/down (warmup/cooldown)
	warmup, cooldown = 3, 5

	opStagePumpsUp *sync.WaitGroup
	opStageStationHours *sync.WaitGroup
	opStagePumpsDown *sync.WaitGroup
)

func CreateStation(stationName string, pumpCount int, pumpRate, operatingTime time.Duration) *Station {
	return &Station{
		Name: stationName,
		Pumps: createPumps(pumpCount, pumpRate),
		OperatingTime: operatingTime,
		OperatingStages: map[string]*sync.WaitGroup{
			"PUMPS_UP": opStagePumpsUp,
			"STATION_HOURS": opStageStationHours,
			"PUMPS_DOWN": opStagePumpsDown,
		},
	}
}

type Station struct {
	Name string
	Pumps []*Pump
	OperatingTime time.Duration
	OperatingStages map[string]*sync.WaitGroup
	IsOpen bool
}

func(s *Station) Open(wg *sync.WaitGroup) {
	s.LogMessage("Station is now OPEN!")
	s.IsOpen = true

	//Add to wait group
	wg.Add(len(s.Pumps)+1)

	//Wait for end of station operating time
	stationTimer := time.NewTimer(s.OperatingTime)

	//Spawn go routine that exits after station operating time
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		defer s.Close()
		<-stationTimer.C
	}(wg)

	//Turn on pumps
	for _, pump := range s.Pumps {
		go pump.On(warmup, wg)

		//Ping pump to ensure on and listening
		*pump.Input<-"hello"
	}
}

func(s *Station) Close() {
	s.LogMessage("Station is closing.  No new transactions accepted.  Shutting down pumps now..")
	s.IsOpen = false

	//Turn off pumps
	for _, pump := range s.Pumps {
		go pump.Off(cooldown)
	}
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
