package main

import (
	"log"
	"time"

	mon "github.com/RuiHirano/vertual_world_system/src/monitor"
	pep "github.com/RuiHirano/vertual_world_system/src/simulator"
	util "github.com/RuiHirano/vertual_world_system/src/util"
)


type Master struct{
	EndDate time.Time
}

func NewMaster(endDate time.Time) *Master{
	return &Master{
		EndDate: endDate,
	}
}

func (mas *Master)Run(){
	// Run Monitor
	log.Printf("monitor runnning...")
	monitor := mon.NewMonitor()
	go monitor.Run()
	time.Sleep(time.Second * 2)

	// PeopleSimulator
	ps := pep.NewPeopleSimulator(monitor)
	ps.AddAgents(util.GetMockAgents(600))

	log.Printf("start simulation")
	for{
		time.Sleep(time.Second * 1)
		ps.Run()
		if mas.EndDate.Before(time.Now()){
			log.Printf("finish cycle")
			break
		}
	}
}

func main() {
	log.Printf("test")
	endDate := time.Date(2020, 10, 25, 23, 59, 59, 0, time.UTC)
	mas := NewMaster(endDate)
	mas.Run()
}
