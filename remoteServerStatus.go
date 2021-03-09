package dcskellergeschwaderserverstatus

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//DCSServer Eagle Dynamics server information about the DCS server
type DCSServer struct {
	ID          string `json:"ID"`
	NAME        string `json:"NAME"`
	IPADDRESS   string `json:"IP_ADDRESS"`
	PORT        string `json:"PORT"`
	MISSIONNAME string `json:"MISSION_NAME"`
	MISSIONTIME string `json:"MISSION_TIME"`
	PLAYERS     string `json:"PLAYERS"`
	PLAYERSMAX  string `json:"PLAYERS_MAX"`
	PASSWORD    string `json:"PASSWORD"`
	URLTODETAIL string `json:"URL_TO_DETAIL"`
}

// DCSServerList Eagle Dynamics DCS server list
type DCSServerList struct {
	SERVERSMAXCOUNT int         `json:"SERVERS_MAX_COUNT"`
	SERVERSMAXDATE  string      `json:"SERVERS_MAX_DATE"`
	PLAYERSCOUNT    int         `json:"PLAYERS_COUNT"`
	MYSERVERS       []DCSServer `json:"MY_SERVERS"`
	SERVERS         []struct {
		NAME                 string `json:"NAME"`
		IPADDRESS            string `json:"IP_ADDRESS"`
		PORT                 string `json:"PORT"`
		MISSIONNAME          string `json:"MISSION_NAME"`
		MISSIONTIME          string `json:"MISSION_TIME"`
		PLAYERS              string `json:"PLAYERS"`
		PLAYERSMAX           string `json:"PLAYERS_MAX"`
		PASSWORD             string `json:"PASSWORD"`
		DESCRIPTION          string `json:"DESCRIPTION"`
		UALIAS0              string `json:"UALIAS_0"`
		MISSIONTIMEFORMATTED string `json:"MISSION_TIME_FORMATTED"`
	} `json:"SERVERS"`
}

// GetServerStatus Get server status from eagle dynamics server list
func GetServerStatus(username string, password string, serverName string) (DCSServer, error) {
	client := &http.Client{}

	url := "https://www.digitalcombatsimulator.com/en/personal/server/?ajax=y&_=" + strconv.FormatInt(time.Now().UTC().Unix(), 10)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DCSServer{}, err
	}

	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return DCSServer{}, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	jsonStart := strings.Index(body, "{")
	serverList := body[jsonStart:]
	serverStatus := &DCSServerList{}
	err = json.Unmarshal([]byte(serverList), serverStatus)
	if err != nil || body == "" {
		return DCSServer{}, err
	}

	for _, server := range serverStatus.MYSERVERS {
		if server.NAME == serverName {
			return server, nil
		}
	}
	return DCSServer{}, errors.New("Server not found")
}
