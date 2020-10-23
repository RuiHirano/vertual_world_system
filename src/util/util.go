package util

import (
	//"log"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
)

type Coord struct {
	Latitude float64
	Longitude float64
}

func getRandomPosition() *Coord {
	// higashiyama
	lat1 := 35.160716
	lon1 := 136.973168
	lat2 := 35.152609
	lon2 := 136.989097
	coord := &Coord{
		Latitude:  lat1 + (lat2-lat1)*rand.Float64(),
		Longitude: lon1 + (lon2-lon1)*rand.Float64(),
	}
	return coord
}

func getHigashiyamaRouteRandomPosition(routes []*RoutePoint) (*Coord, *Coord) {
	// higashiyama
	//routes := GetRoutes()
	route1 := routes[rand.Intn(len(routes))]
	point1 := route1.Point
	point2 := route1.NeighborPoints[rand.Intn(int(len(route1.NeighborPoints)))].Point
	lat1 := point1.Latitude
	lon1 := point1.Longitude
	lat2 := point2.Latitude
	lon2 := point2.Longitude
	x := lat1 + (lat2-lat1)*rand.Float64()
	y := ((lon2-lon1) / (lat2 - lat1)) * (x - lat1) + lon1
	position := &Coord{
		Latitude:  x,
		Longitude: y,
	}

	destination := point2
	return position, destination
}

type Agent struct{
	ID string
	Position *Coord
	Direction float64
	Speed float64
	Destination *Coord
}

func GetMockAgents(num int) []*Agent{

	// higashiyama route
	routes := GetRoutes()

	agents := make([]*Agent, 0)
	for i := 0; i < int(num); i++ {
		position, destination := getHigashiyamaRouteRandomPosition(routes)
		uid, _ := uuid.NewRandom()
		agents = append(agents, &Agent{
			ID:   strconv.Itoa(int(uid.ID())),
			Position: position,
			Speed: 60,
			Direction: 30,
			Destination: destination,
		})
		//log.Printf("agent: %v %v",uid.ID(), getRandomPosition())
	}
	return agents
}