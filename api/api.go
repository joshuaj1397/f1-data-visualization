package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joshuaj1397/f1-data-visualization/model"
)

func GetCircuit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	circuitId, err := strconv.Atoi(params["circuitId"])
	circuit, err := model.GetCircuit(circuitId)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(circuit)
}

func GetRace(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	year, err := strconv.Atoi(params["year"])
	circuitId, err := strconv.Atoi(params["circuitId"])
	race, err := model.GetRace(year, circuitId)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(race)
}

func GetRaceLapTimes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	raceId, err := strconv.ParseInt(params["raceId"], 10, 32)
	driverId, err := strconv.ParseInt(params["driverId"], 10, 32)
	laptimes := model.GetLapTimes(int(raceId), int(driverId))
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(laptimes)
}

func GetDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	driverId, err := strconv.ParseInt(params["year"], 10, 32)
	driver, err := model.GetDriver(int(driverId))
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(driver)
}

func GetConstructor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	constructorId, err := strconv.ParseInt(params["constructorId"], 10, 32)
	constructor, err := model.GetConstructor(int(constructorId))
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(constructor)
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	raceId, err := strconv.ParseInt(params["raceId"], 10, 32)
	results := model.GetResults(int(raceId))
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(results)
}
