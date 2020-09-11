package main

import (
	"net/http"
)

// ProcessReplay .
func ProcessReplay(request *http.Request) uint {
	downloadPath := ExtractURL(request)

	urlToCheck := DownloadedURL{URL: downloadPath}
	id := urlToCheck.DoesExistsInDB()
	if id != 0 {
		return id
	}
	replayName := DownloadReplay(downloadPath)
	id = ProcessReplayFromFile(replayName)

	urlToCheck.MatchID = id
	urlToCheck.InsertIntoDB()

	return id

}

// ProcessReplayFromFile .
func ProcessReplayFromFile(replayName string) uint {

	rawMatch := ReadMatchFromFile(replayName)

	calculatedMatch := calculateMatch(rawMatch)

	return calculatedMatch.ID
}
