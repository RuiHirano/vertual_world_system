package simulator

import (
	"log"

	"math"
	"math/rand"

	rvo "github.com/RuiHirano/rvo2-go/src/rvosimulator"
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

// SetAgents : 
func (ps *PeopleSimulator)SetAgents(agents []*util.Agent) {
	log.Printf("set agents")
	ps.Agents = agents
}

// Run : 
func (ps *PeopleSimulator)Run() {
	log.Printf("run")

	// Higashiyama Route
	higashi := util.NewHigashiyama()
	// 時間によってエージェント数を増減させる
	// 時間によってRouteの人気度を増減させる
	// サイネージによってPointの人気度を増減させる

	// RVO2
	rvo2 := NewRVO2(higashi.Routes)
	agents := rvo2.ForwardStep(ps.Agents)
	ps.SetAgents(agents)
	
	// Send Monitor
	ps.Monitor.SendAgents(ps.Agents)

}

type RVO2 struct {
	Sim *rvo.RVOSimulator
	Routes []*util.RoutePoint
}

func NewRVO2(routes []*util.RoutePoint) *RVO2 {
	timeStep := 0.1
	neighborDist := 0.00003 // どのくらいの距離の相手をNeighborと認識するか?Neighborとの距離をどのくらいに保つか？ぶつかったと認識する距離？
	maxneighbors := 3       // 周り何体を計算対象とするか
	timeHorizon := 1.0
	timeHorizonObst := 1.0
	radius := 0.00001  // エージェントの半径
	maxSpeed := 0.0004 // エージェントの最大スピード
	return &RVO2{
		Sim: rvo.NewRVOSimulator(timeStep, neighborDist, maxneighbors, timeHorizon, timeHorizonObst, radius, maxSpeed, &rvo.Vector2{X: 0, Y: 0}),
		Routes: routes,
	}
}

// ForwardStep :　次の時刻のエージェントを計算する関数
func (rvo2 *RVO2) ForwardStep(agents []*util.Agent) []*util.Agent {
	nextAgents := rvo2.CalcNextAgents(agents)

	return nextAgents
}

// ForwardStep :　次の時刻のエージェントを計算する関数
func (rvo2 *RVO2) CalcNextAgents(agents []*util.Agent) []*util.Agent {

	rvo2.SetupScenario(agents)
	rvo2.Sim.DoStep()

	// rvoエージェントからutil.Agentに変換
	nextAgents := make([]*util.Agent, 0)
	for rvoId, agent := range agents {

		// rvoの位置情報を緯度経度に変換する (rvoId: indexがrvoのidとマッチしている)
		rvoPosition := rvo2.Sim.GetAgentPosition(int(rvoId))
		nextCoord := &util.Coord{
			Latitude:  rvoPosition.Y,
			Longitude: rvoPosition.X,
		}
		goalVector := rvo2.Sim.GetAgentGoalVector(int(rvoId))
		direction := math.Atan2(goalVector.Y, goalVector.X)
		speed := agent.Speed

		// destinationにたどり着いたら次の行先を決定する
		destination := agent.Destination
		_, distance := rvo2.CalcDirectionAndDistance(nextCoord, destination)
		if distance < 10{
			destination = rvo2.GetNextDestination(destination)
		}

		nextAgents = append(nextAgents, &util.Agent{
			ID:    agent.ID,
			Position:      nextCoord,
			Direction:     direction,
			Speed:         speed,
			Destination:   destination,
		})
	}
	return nextAgents
	
}

// SetupScenario: Scenarioを設定する関数
func (rvo2 *RVO2) SetupScenario(agents []*util.Agent) {
	// Set Agent
	for _, agent := range agents {

		position := &rvo.Vector2{X: agent.Position.Longitude, Y: agent.Position.Latitude}
		goal := &rvo.Vector2{X: agent.Destination.Longitude, Y: agent.Destination.Latitude}

		// Agentを追加
		id, _ := rvo2.Sim.AddDefaultAgent(position)

		// 目的地を設定
		rvo2.Sim.SetAgentGoal(id, goal)

		// エージェントの速度方向ベクトルを設定
		goalVector := rvo2.Sim.GetAgentGoalVector(id)
		rvo2.Sim.SetAgentPrefVelocity(id, goalVector)
		//sim.SetAgentMaxSpeed(id, float64(util.MaxSpeed))
	}
}

// CalcDirectionAndDistance: 目的地までの距離と角度を計算する関数
func (rvo2 *RVO2) CalcDirectionAndDistance(startCoord *util.Coord, goalCoord *util.Coord) (float64, float64) {

	r := 6378137 // equatorial radius
	sLat := startCoord.Latitude * math.Pi / 180
	sLon := startCoord.Longitude * math.Pi / 180
	gLat := goalCoord.Latitude * math.Pi / 180
	gLon := goalCoord.Longitude * math.Pi / 180
	dLon := gLon - sLon
	dLat := gLat - sLat
	cLat := (sLat + gLat) / 2
	dx := float64(r) * float64(dLon) * math.Cos(float64(cLat))
	dy := float64(r) * float64(dLat)

	distance := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	direction := float64(0)
	if dx != 0 && dy != 0 {
		direction = math.Atan2(dy, dx) * 180 / math.Pi
	}

	return direction, distance
}

// GetNextDestination: 次の目的地を求める関数
func (rvo2 *RVO2) GetNextDestination(destination *util.Coord) *util.Coord {
	newDestination := destination
	for _, route := range rvo2.Routes {
		if route.Point.Longitude == destination.Longitude && route.Point.Latitude == destination.Latitude {
			index := rand.Intn(len(route.NeighborPoints))
			nextRoute := route.NeighborPoints[index]
			newDestination = nextRoute.Point
			break
		}
	}
	return newDestination
}