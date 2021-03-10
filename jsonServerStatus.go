package dcskellergeschwaderserverstatus

import (
	"encoding/json"
	"io/ioutil"
)

// DCSServerStatusSeason Server hook exported
type DCSServerStatusSeason struct {
	Temperature int `json:"temperature"`
}

// DCSServerStatusWind Server hook exported wind information
type DCSServerStatusWind struct {
	Speed int `json:"speed"`
	Dir   int `json:"dir"`
}

// DCSServerStatusWinds Server hook exported winds information
type DCSServerStatusWinds struct {
	At8000   DCSServerStatusWind `json:"at8000"`
	At2000   DCSServerStatusWind `json:"at2000"`
	AtGround DCSServerStatusWind `json:"atGround"`
}

// DCSServerStatusClouds Server hook exported cloud information
type DCSServerStatusClouds struct {
	Density   int `json:"density"`
	Base      int `json:"base"`
	Thickness int `json:"thickness"`
}

// DCSServerStatusWeather Server hook exported weather information
type DCSServerStatusWeather struct {
	Wind   DCSServerStatusWinds  `json:"wind"`
	Season DCSServerStatusSeason `json:"season"`
	Clouds DCSServerStatusClouds `json:"clouds"`
}

// DCSServerStatusPlayer Server hook exported player status
type DCSServerStatusPlayer struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Role       string  `json:"role"`
	OnlineTime float64 `json:"onlineTime"`
}

// DCSServerStatus Server hook exported json status
type DCSServerStatus struct {
	Players         []DCSServerStatusPlayer `json:"players"`
	MissionsNames   []string                `json:"missionsNames"`
	MissionTimeLeft int                     `json:"missionTimeLeft"`
	Time            int                     `json:"time"`
	Weather         DCSServerStatusWeather  `json:"weather"`
}

// ReadServerStatusFile Reads the server status file exported by the server hook
func ReadServerStatusFile(filePath string) (DCSServerStatus, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	status := DCSServerStatus{}
	if err != nil {
		return status, err
	}

	err = json.Unmarshal(fileBytes, &status)
	if err != nil {
		return status, err
	}

	return status, nil
}
