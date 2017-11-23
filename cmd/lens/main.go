package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"linuxlens"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

func seedRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/processes", GetProcessList).Methods("GET")
	return router
}

func GetProcessList(w http.ResponseWriter, r *http.Request) {

	log.Infof("[%s] %s", r.Method, r.URL)
	processes := linuxlens.GetProcesses()
	err := json.NewEncoder(w).Encode(processes)
	if err != nil {
		log.Error(err)
	}
}

func init() {
	formatter := log.TextFormatter{
		FullTimestamp: true,
	}

	log.SetFormatter(&formatter)
}

func main() {
	log.Info("booting up server")
	router := seedRoutes()
	log.Fatal(http.ListenAndServe(":8000", router))
}
