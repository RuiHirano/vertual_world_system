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
	time.Sleep(time.Second * 1)

	// Run Higashiyama Listener
	log.Printf("higashiyama runnning...")
	higashi := util.NewHigashiyama()
	go higashi.RunListener()
	time.Sleep(time.Second * 1)

	// PeopleSimulator
	ps := pep.NewPeopleSimulator(monitor)
	ps.AddAgents(util.GetMockAgents(100))

	// 時間によってエージェント数を増減させる
	go ps.ChangeAgent(higashi.Config)

	log.Printf("start simulation")
	for{
		t1 := time.Now()

		ps.Run(higashi.Routes)

		t2 := time.Now()
		duration := t2.Sub(t1).Milliseconds()
		interval := int64(1000) // 周期ms
		if duration > interval {
			log.Printf("time cycle delayed... Duration: %d", duration)
		} else {
			// 待機
			log.Printf("cycle finished Duration: %d ms, Wait: %d ms", duration, interval-duration)
			time.Sleep(time.Duration(interval-duration) * time.Millisecond)
		}
		
		if mas.EndDate.Before(time.Now()){
			log.Printf("finish simulation...")
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
