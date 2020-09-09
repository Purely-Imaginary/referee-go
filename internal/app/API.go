package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func prepareData(data interface{}, w http.ResponseWriter, r *http.Request) string {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	outputMessage, err := json.Marshal(&data)
	Check(err)
	return string(outputMessage)
}

// GetLastMatches ..
func GetLastMatches(w http.ResponseWriter, r *http.Request) {
	outputMessage := prepareData(GetLastMatchesFromDB(200), w, r)
	fmt.Fprintf(w, outputMessage)
}

// GetPlayersTable ..
func GetPlayersTable(w http.ResponseWriter, r *http.Request) {
	outputMessage := prepareData(GetPlayersTableFromDB(), w, r)
	fmt.Fprintf(w, outputMessage)
}

// GetPlayersSnapshots ..
func GetPlayersSnapshots(w http.ResponseWriter, r *http.Request) {
	outputMessage := prepareData(GetPlayersSnapshotsFromDB(), w, r)
	fmt.Fprintf(w, outputMessage)
}