package util

import (
	"log"
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

type Agent struct{
	ID string
	Position *Coord
	Direction float64
	Speed float64
	Destination *Coord
}

func GetMockAgents(num int) []*Agent{
	agents := make([]*Agent, 0)
	for i := 0; i < int(num); i++ {
		uid, _ := uuid.NewRandom()
		agents = append(agents, &Agent{
			ID:   strconv.Itoa(int(uid.ID())),
			Position: getRandomPosition(),
			Speed: 60,
			Direction: 30,
			Destination: getRandomPosition(),
		})
		log.Printf("agent: %v %v",uid.ID(), getRandomPosition())
	}
	return agents
}