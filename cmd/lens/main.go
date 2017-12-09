package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"linuxlens"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

var (
	cfg *linuxlens.ServerConfig
)

func seedRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/processes", GetProcessList).Methods("GET")
	router.HandleFunc("/cpu", GetCpuInfo).Methods("GET")
	router.HandleFunc("/memory", GetMemInfo).Methods("GET")
	return router
}

func GetMemInfo(w http.ResponseWriter, r *http.Request) {
	log.Infof("[%s] %s", r.Method, r.URL)

	memStats, err := linuxlens.ParseMemInfo()
	if err != nil {
		http.Error(w, "failed to fetch mem info", 500)
	}

	err = json.NewEncoder(w).Encode(memStats)
	if err != nil {
		log.Error(err)
	}
}

func GetCpuInfo(w http.ResponseWriter, r *http.Request) {
	log.Infof("[%s] %s", r.Method, r.URL)

	cpuInfo, err := linuxlens.ParseCpuInfo()
	if err != nil {
		http.Error(w, "failed to fetch cpu info", 500)
	}

	err = json.NewEncoder(w).Encode(cpuInfo)
	if err != nil {
		log.Error(err)
	}
}

func GetProcessList(w http.ResponseWriter, r *http.Request) {

	log.Infof("[%s] %s", r.Method, r.URL)
	processes := linuxlens.GetProcesses(cfg)

	err := json.NewEncoder(w).Encode(processes)
	if err != nil {
		log.Error(err)
	}
}

func init() {
	formatter := log.TextFormatter{FullTimestamp: true}
	log.SetFormatter(&formatter)
}

func main() {

	var err error
	cfg, err = linuxlens.LoadConfig("server.cfg")
	if err != nil {
		log.Fatal(err)
	}

	log.Info("booting up server at ", cfg.ServerAddress)
	router := seedRoutes()
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}
