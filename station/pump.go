package station

import (
	"log"
	"sync"
	"time"
)

type Pump struct {
	ID int
	RatePerGallon time.Duration
	Input *chan string
}

func(p *Pump) On(warmup int, wg *sync.WaitGroup) {
	p.LogMessage("Powering up..")
	
	// warm up period
	time.Sleep(time.Second * time.Duration(warmup))
	
	// start listening to Input
	go p.pollInput(wg)
}

func(p *Pump) Off(cooldown int) {
	// cooldown period
	p.LogMessage("Powering down..")

	time.Sleep(time.Second * time.Duration(cooldown))
	
	*p.Input<-"off"
}

func(p *Pump) pollInput(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		//wait for input message
		message := <-*p.Input
		switch message {
		case "hello":
			p.LogMessage("Powered up!")
		case "off":
			p.LogMessage("Powered down.")
			return
		}
	}
}

func(p *Pump) LogMessage(msg string) {
	log.Printf("[PUMP %d] %s\n\n", p.ID, msg)
}