package util

import (
	"math/rand"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/////////////////////////////////////////////////////
//////// util for creating higashiyama route ////////
///////////////////////////////////////////////////////

type RoutePoint struct {
	Id             int
	Name           string
	Point          *Coord
	Popularity     float64
	NeighborPoints []*RoutePoint
}

type Signage struct {
	Id             int
	Name           string
	Point          *Coord
	NeighborPoints []*RoutePoint
}

type Higashiyama struct {
	Routes          []*RoutePoint
	Signages          []*Signage
	MaxPeople  int
	MinPeople int
}

func NewHigashiyama() *Higashiyama{
	return &Higashiyama{
		Routes: GetRoutes(),
		Signages: []*Signage{},
		MaxPeople: 3000,
		MinPeople: 300,
	} 
}

func (hig *Higashiyama) RunListener(){
	e := echo.New()

    // 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // ルーティング
    e.GET("/hello", func () echo.HandlerFunc{
		return func(c echo.Context) error {     //c をいじって Request, Responseを色々する 
			return c.String(http.StatusOK, "Hello World")
		}
	}())

    // サーバー起動
    e.Start(":8000")    //ポート番号指定してね
}

func GetRoutes() []*RoutePoint {
	routes := []*RoutePoint{
		{
			Id: 0, Name: "gate", Point: &Coord{Longitude: 136.974024, Latitude: 35.158995},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 1, Name: "enterance", Point: &Coord{Longitude: 136.974688, Latitude: 35.158228}},
			},
		},
		{
			Id: 1, Name: "enterance", Point: &Coord{Longitude: 136.974688, Latitude: 35.158228},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 0, Name: "gate", Point: &Coord{Longitude: 136.974024, Latitude: 35.158995}},
				{Id: 2, Name: "rightEnt", Point: &Coord{Longitude: 136.974645, Latitude: 35.157958}},
				{Id: 3, Name: "leftEnt", Point: &Coord{Longitude: 136.974938, Latitude: 35.158164}},
			},
		},
		{
			Id: 2, Name: "rightEnt", Point: &Coord{Longitude: 136.974645, Latitude: 35.157958},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 1, Name: "enterance", Point: &Coord{Longitude: 136.974688, Latitude: 35.158228}},
				{Id: 4, Name: "road1", Point: &Coord{Longitude: 136.974864, Latitude: 35.157823}},
			},
		},
		{
			Id: 3, Name: "leftEnt", Point: &Coord{Longitude: 136.974938, Latitude: 35.158164},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 1, Name: "enterance", Point: &Coord{Longitude: 136.974688, Latitude: 35.158228}},
				{Id: 5, Name: "road2", Point: &Coord{Longitude: 136.975054, Latitude: 35.158001}},
				{Id: 17, Name: "north1", Point: &Coord{Longitude: 136.976395, Latitude: 35.158410}},
			},
		},
		{
			Id: 4, Name: "road1", Point: &Coord{Longitude: 136.974864, Latitude: 35.157823},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 2, Name: "rightEnt", Point: &Coord{Longitude: 136.974645, Latitude: 35.157958}},
				{Id: 5, Name: "road2", Point: &Coord{Longitude: 136.975054, Latitude: 35.158001}},
				{Id: 6, Name: "road3", Point: &Coord{Longitude: 136.975517, Latitude: 35.157096}},
			},
		},
		{
			Id: 5, Name: "road2", Point: &Coord{Longitude: 136.975054, Latitude: 35.158001},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 3, Name: "leftEnt", Point: &Coord{Longitude: 136.974938, Latitude: 35.158164}},
				{Id: 4, Name: "road1", Point: &Coord{Longitude: 136.974864, Latitude: 35.157823}},
			},
		},
		{
			Id: 6, Name: "road3", Point: &Coord{Longitude: 136.975517, Latitude: 35.157096},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 7, Name: "road4", Point: &Coord{Longitude: 136.975872, Latitude: 35.156678}},
				{Id: 4, Name: "road1", Point: &Coord{Longitude: 136.974864, Latitude: 35.157823}},
			},
		},
		{
			Id: 7, Name: "road4", Point: &Coord{Longitude: 136.975872, Latitude: 35.156678},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 6, Name: "road3", Point: &Coord{Longitude: 136.975517, Latitude: 35.157096}},
				{Id: 8, Name: "road5", Point: &Coord{Longitude: 136.976314, Latitude: 35.156757}},
				{Id: 10, Name: "burger", Point: &Coord{Longitude: 136.976960, Latitude: 35.155697}},
			},
		},
		{
			Id: 8, Name: "road5", Point: &Coord{Longitude: 136.976314, Latitude: 35.156757},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 6, Name: "road3", Point: &Coord{Longitude: 136.975517, Latitude: 35.157096}},
				{Id: 9, Name: "toilet", Point: &Coord{Longitude: 136.977261, Latitude: 35.155951}},
			},
		},
		{
			Id: 9, Name: "toilet", Point: &Coord{Longitude: 136.977261, Latitude: 35.155951},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 8, Name: "road5", Point: &Coord{Longitude: 136.976314, Latitude: 35.156757}},
				{Id: 10, Name: "burger", Point: &Coord{Longitude: 136.976960, Latitude: 35.155697}},
			},
		},
		{
			Id: 10, Name: "burger", Point: &Coord{Longitude: 136.976960, Latitude: 35.155697},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 8, Name: "road5", Point: &Coord{Longitude: 136.976314, Latitude: 35.156757}},
				{Id: 7, Name: "road4", Point: &Coord{Longitude: 136.975872, Latitude: 35.156678}},
				{Id: 11, Name: "lake1", Point: &Coord{Longitude: 136.978217, Latitude: 35.155266}},
			},
		},
		{
			Id: 11, Name: "lake1", Point: &Coord{Longitude: 136.978217, Latitude: 35.155266},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 10, Name: "burger", Point: &Coord{Longitude: 136.976960, Latitude: 35.155697}},
				{Id: 12, Name: "lake2", Point: &Coord{Longitude: 136.978623, Latitude: 35.155855}},
				{Id: 16, Name: "lake6", Point: &Coord{Longitude: 136.978297, Latitude: 35.154755}},
			},
		},
		{
			Id: 12, Name: "lake2", Point: &Coord{Longitude: 136.978623, Latitude: 35.155855},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 11, Name: "lake1", Point: &Coord{Longitude: 136.978217, Latitude: 35.155266}},
				{Id: 13, Name: "lake3", Point: &Coord{Longitude: 136.979657, Latitude: 35.155659}},
			},
		},
		{
			Id: 13, Name: "lake3", Point: &Coord{Longitude: 136.979657, Latitude: 35.155659},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 12, Name: "lake2", Point: &Coord{Longitude: 136.978623, Latitude: 35.155855}},
				{Id: 14, Name: "lake4", Point: &Coord{Longitude: 136.980489, Latitude: 35.154484}},
				{Id: 26, Name: "east6", Point: &Coord{Longitude: 136.984100, Latitude: 35.153693}},
				{Id: 22, Name: "east1", Point: &Coord{Longitude: 136.981124, Latitude: 35.157283}},
				{Id: 27, Name: "east-in1", Point: &Coord{Longitude: 136.982804, Latitude: 35.154175}},
			},
		},
		{
			Id: 14, Name: "lake4", Point: &Coord{Longitude: 136.980489, Latitude: 35.154484},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 13, Name: "lake3", Point: &Coord{Longitude: 136.979657, Latitude: 35.155659}},
				{Id: 15, Name: "lake5", Point: &Coord{Longitude: 136.980143, Latitude: 35.153869}},
			},
		},
		{
			Id: 15, Name: "lake5", Point: &Coord{Longitude: 136.980143, Latitude: 35.153869},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 14, Name: "lake4", Point: &Coord{Longitude: 136.980489, Latitude: 35.154484}},
				{Id: 16, Name: "lake6", Point: &Coord{Longitude: 136.978297, Latitude: 35.154755}},
			},
		},
		{
			Id: 16, Name: "lake6", Point: &Coord{Longitude: 136.978297, Latitude: 35.154755},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 11, Name: "lake1", Point: &Coord{Longitude: 136.978217, Latitude: 35.155266}},
				{Id: 15, Name: "lake5", Point: &Coord{Longitude: 136.980143, Latitude: 35.153869}},
			},
		},
		{
			Id: 17, Name: "north1", Point: &Coord{Longitude: 136.976395, Latitude: 35.158410},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 3, Name: "leftEnt", Point: &Coord{Longitude: 136.974938, Latitude: 35.158164}},
				{Id: 5, Name: "road2", Point: &Coord{Longitude: 136.975054, Latitude: 35.158001}},
				{Id: 18, Name: "north2", Point: &Coord{Longitude: 136.977821, Latitude: 35.159220}},
			},
		},
		{
			Id: 18, Name: "north2", Point: &Coord{Longitude: 136.977821, Latitude: 35.159220},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 17, Name: "north1", Point: &Coord{Longitude: 136.976395, Latitude: 35.158410}},
				{Id: 19, Name: "medaka", Point: &Coord{Longitude: 136.979040, Latitude: 35.158147}},
			},
		},
		{
			Id: 19, Name: "medaka", Point: &Coord{Longitude: 136.979040, Latitude: 35.158147},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 18, Name: "north2", Point: &Coord{Longitude: 136.977821, Latitude: 35.159220}},
				{Id: 20, Name: "tower", Point: &Coord{Longitude: 136.978846, Latitude: 35.157108}},
			},
		},
		{
			Id: 20, Name: "tower", Point: &Coord{Longitude: 136.978846, Latitude: 35.157108},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 19, Name: "medaka", Point: &Coord{Longitude: 136.979040, Latitude: 35.158147}},
				{Id: 21, Name: "north-out", Point: &Coord{Longitude: 136.977890, Latitude: 35.156563}},
			},
		},
		{
			Id: 21, Name: "north-out", Point: &Coord{Longitude: 136.977890, Latitude: 35.156563},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 20, Name: "tower", Point: &Coord{Longitude: 136.978846, Latitude: 35.157108}},
				{Id: 17, Name: "north1", Point: &Coord{Longitude: 136.976395, Latitude: 35.158410}},
				{Id: 9, Name: "toilet", Point: &Coord{Longitude: 136.977261, Latitude: 35.155951}},
			},
		},
		{
			Id: 22, Name: "east1", Point: &Coord{Longitude: 136.981124, Latitude: 35.157283},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 13, Name: "lake3", Point: &Coord{Longitude: 136.979657, Latitude: 35.155659}},
				{Id: 23, Name: "east2", Point: &Coord{Longitude: 136.984350, Latitude: 35.157271}},
			},
		},
		{
			Id: 23, Name: "east2", Point: &Coord{Longitude: 136.984350, Latitude: 35.157271},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 22, Name: "east1", Point: &Coord{Longitude: 136.981124, Latitude: 35.157283}},
				{Id: 24, Name: "east3", Point: &Coord{Longitude: 136.987567, Latitude: 35.158233}},
			},
		},
		{
			Id: 24, Name: "east3", Point: &Coord{Longitude: 136.987567, Latitude: 35.158233},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 23, Name: "east2", Point: &Coord{Longitude: 136.984350, Latitude: 35.157271}},
				{Id: 25, Name: "east4", Point: &Coord{Longitude: 136.988522, Latitude: 35.157286}},
			},
		},
		{
			Id: 25, Name: "east4", Point: &Coord{Longitude: 136.988522, Latitude: 35.157286},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 24, Name: "east3", Point: &Coord{Longitude: 136.987567, Latitude: 35.158233}},
				{Id: 25, Name: "east5", Point: &Coord{Longitude: 136.988355, Latitude: 35.155838}},
			},
		},
		{
			Id: 25, Name: "east5", Point: &Coord{Longitude: 136.988355, Latitude: 35.155838},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 25, Name: "east4", Point: &Coord{Longitude: 136.988522, Latitude: 35.157286}},
				{Id: 26, Name: "east6", Point: &Coord{Longitude: 136.984100, Latitude: 35.153693}},
			},
		},
		{
			Id: 26, Name: "east6", Point: &Coord{Longitude: 136.984100, Latitude: 35.153693},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 25, Name: "east5", Point: &Coord{Longitude: 136.988355, Latitude: 35.155838}},
				{Id: 13, Name: "lake3", Point: &Coord{Longitude: 136.979657, Latitude: 35.155659}},
				{Id: 27, Name: "east-in1", Point: &Coord{Longitude: 136.982804, Latitude: 35.154175}},
			},
		},
		{
			Id: 27, Name: "east-in1", Point: &Coord{Longitude: 136.982804, Latitude: 35.154175},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 26, Name: "east6", Point: &Coord{Longitude: 136.984100, Latitude: 35.153693}},
				{Id: 13, Name: "lake3", Point: &Coord{Longitude: 136.979657, Latitude: 35.155659}},
				{Id: 28, Name: "east-in2", Point: &Coord{Longitude: 136.984244, Latitude: 35.156283}},
			},
		},
		{
			Id: 28, Name: "east-in2", Point: &Coord{Longitude: 136.984244, Latitude: 35.156283},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 29, Name: "east-in3", Point: &Coord{Longitude: 136.987627, Latitude: 35.157104}},
				{Id: 27, Name: "east-in1", Point: &Coord{Longitude: 136.982804, Latitude: 35.154175}},
			},
		},
		{
			Id: 29, Name: "east-in3", Point: &Coord{Longitude: 136.987627, Latitude: 35.157104},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 28, Name: "east-in2", Point: &Coord{Longitude: 136.984244, Latitude: 35.156283}},
				{Id: 30, Name: "east-in4", Point: &Coord{Longitude: 136.986063, Latitude: 35.155353}},
			},
		},
		{
			Id: 30, Name: "east-in4", Point: &Coord{Longitude: 136.986063, Latitude: 35.155353},
			Popularity: rand.Float64(),
			NeighborPoints: []*RoutePoint{
				{Id: 29, Name: "east-in3", Point: &Coord{Longitude: 136.987627, Latitude: 35.157104}},
				{Id: 26, Name: "east6", Point: &Coord{Longitude: 136.984100, Latitude: 35.153693}},
			},
		},
	}

	return routes
}