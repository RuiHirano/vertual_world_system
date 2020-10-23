package monitor

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	util "github.com/RuiHirano/vertual_world_system/src/util"
	gosocketio "github.com/mtfelian/golang-socketio"
)

var (
	ioserv         *gosocketio.Server
	mu             sync.Mutex
	assetsDir      http.FileSystem
	monitoraddr       = flag.String("monitoraddr", "127.0.0.1:5000", "Monitor Listening Address")
)

// Monitor : 
type Monitor struct {
}

func NewMonitor() *Monitor{
	return &Monitor{}
}

type MapMarker struct {
	mtype int32   `json:"mtype"`
	id    int32   `json:"id"`
	lat   float32 `json:"lat"`
	lon   float32 `json:"lon"`
	angle float32 `json:"angle"`
	speed int32   `json:"speed"`
	area  int32   `json:"area"`
}

// GetJson: json化する関数
func (m *MapMarker) GetJson() string {
	s := fmt.Sprintf("{\"mtype\":%d,\"id\":%d,\"lat\":%f,\"lon\":%f,\"angle\":%f,\"speed\":%d,\"area\":%d}",
		m.mtype, m.id, m.lat, m.lon, m.angle, m.speed, m.area)
	return s
}

// AddAgents : 
func (ps *Monitor)SendAgents(agents []*util.Agent) {
	log.Printf("send agents %d", len(agents))
	jsonAgents := make([]string, 0)
	for _, agent := range agents {
		id, _ := strconv.Atoi(agent.ID)
		//log.Printf("id: %v", agent)
		mm := &MapMarker{
			mtype: int32(0),
			id:    int32(id),
			lat:   float32(agent.Position.Latitude),
			lon:   float32(agent.Position.Longitude),
			angle: float32(agent.Direction),
			speed: int32(agent.Speed),
		}
		jsonAgents = append(jsonAgents, mm.GetJson())
	}

	// jsonRealAgents: そのままMarshalしたもの
	jsonRealAgents, err := json.Marshal(agents)
    if err != nil {
        fmt.Println(err)
        return
    }

	mu.Lock()
	if ioserv != nil{
		ioserv.BroadcastToAll("agents", jsonAgents)
		ioserv.BroadcastToAll("realAgents", jsonRealAgents)
	}else{
		log.Printf("ioserv is nil. can't send agents")
	}

	mu.Unlock()
}

func (ps *Monitor)Run() {
	log.Printf("run monitor")
	// Run HarmowareVis Monitor
	ioserv = runServer()
	log.Printf("Running Sio Server..\n")
	if ioserv == nil {
		os.Exit(1)
	}
	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", ioserv)
	serveMux.HandleFunc("/", assetsFileHandler)
	log.Printf("Starting Harmoware VIS  Provider on %s", *monitoraddr)
	err := http.ListenAndServe(*monitoraddr, serveMux)
	if err != nil {
		log.Fatal(err)
	}
}

func runServer() *gosocketio.Server {

	currentRoot, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	d := filepath.Join(currentRoot, "monitor", "build")

	assetsDir = http.Dir(d)
	log.Println("AssetDir:", assetsDir)

	server := gosocketio.NewServer()

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Printf("Connected from %s as %s", c.IP(), c.Id())
	})

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Printf("Disconnected from %s as %s", c.IP(), c.Id())
	})

	return server
}

// assetsFileHandler for static Data
func assetsFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return
	}

	file := r.URL.Path
	//	log.Printf("Open File '%s'",file)
	if file == "/" {
		file = "/index.html"
	}
	f, err := assetsDir.Open(file)
	if err != nil {
		log.Printf("can't open file %s: %v\n", file, err)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Printf("can't open file %s: %v\n", file, err)
		return
	}
	http.ServeContent(w, r, file, fi.ModTime(), f)
}