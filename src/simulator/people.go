package simulator

import (
	"log"

	mon "github.com/RuiHirano/vertual_world_system/src/monitor"
	util "github.com/RuiHirano/vertual_world_system/src/util"
)



// PeopleSimulator : 
type PeopleSimulator struct {
	Agents []*util.Agent
	Monitor *mon.Monitor
}

func NewPeopleSimulator(monitor *mon.Monitor) *PeopleSimulator{
	return &PeopleSimulator{
		Agents: []*util.Agent{},
		Monitor: monitor,
	}
}

// AddAgents : 
func (ps *PeopleSimulator)AddAgents(agents []*util.Agent) {
	log.Printf("add agents")
	ps.Agents = append(ps.Agents, agents...)
}

// Run : 
func (ps *PeopleSimulator)Run() {
	log.Printf("run")
	
	// Send Monitor
	ps.Monitor.SendAgents(ps.Agents)

}


