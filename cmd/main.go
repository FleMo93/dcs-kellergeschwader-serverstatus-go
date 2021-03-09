package main

import (
	"log"
	"os"
	"strings"

	m "github.com/FleMo93/dcs-kellergeschwader-serverstatus-go"
)

func main() {
	arg := os.Args
	username := ""
	password := ""
	serverName := ""
	serverStatusFile := ""

	for _, ele := range arg {
		if strings.Index(ele, "--username ") == 0 {
			username = ele[11:]
		} else if strings.Index(ele, "--password ") == 0 {
			password = ele[11:]
		} else if strings.Index(ele, "--serverName ") == 0 {
			serverName = ele[13:]
		} else if strings.Index(ele, "--serverStatusFile ") == 0 {
			serverStatusFile = ele[19:]
		}
	}

	jsonStatus, err := m.ReadServerStatusFile(serverStatusFile)
	if err != nil {
		log.Panic(err)
	}
	edStatus, err := m.GetServerStatus(username, password, serverName)
	if err != nil {
		log.Panic(err)
	}

	log.Print(jsonStatus)
	log.Print(edStatus)
}
