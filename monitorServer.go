package monitor_server

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"monitor-client"
	"log"
)

var snapshots = []monitor_client.MonitorSnapshot{}

func MonitorServer() {
	// REST Api
	// POST /monitor/push
	// Add a new MonitorSnapshot to the collection
	http.HandleFunc("/monitor/push", pushHandler)
	// GET /monitor/pull
	// Return collection of MonitorSnapshot
	http.HandleFunc("/monitor/pull", pullHandler)
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	log.Println("MonitorSnapshot received in bytes: ", b)

	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg monitor_client.MonitorSnapshot

	err = json.Unmarshal(b, &msg)
	log.Println("MonitorSnapshot in type received:", msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	snapshots = append(snapshots, msg);
}

func pullHandler(w http.ResponseWriter, r *http.Request) {
	//Response
	output, err := json.Marshal(snapshots)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
