package dcskellergeschwaderserverstatus

import (
	"encoding/json"
	"io/ioutil"
)

// DCSServerStatusPlayer Server hook exportet player status
type DCSServerStatusPlayer struct {
	Name       string  `json:"name"`
	Role       string  `json:"role"`
	OnlineTime float64 `json:"onlineTime"`
}

// DCSServerStatus Server hook exportet json status
type DCSServerStatus struct {
	Players       map[string]DCSServerStatusPlayer `json:"players"`
	MissionsNames []string                         `json:"missionsNames"`
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
