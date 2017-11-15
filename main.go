package main

import (
	"fmt"
	"log"
	"os"

	"github.com/t0mk/packswgo"
)

const APITokenVar = "PACKET_AUTH_TOKEN"

func getCfg() *packswgo.Configuration {
	cfg := packswgo.NewConfiguration()
	cfg.APIKey["X-Auth-Token"] = os.Getenv(APITokenVar)
	return cfg
}

func main() {
	cfg := getCfg()

	devicesAPI := packswgo.DevicesApi{Configuration: cfg}
	projectID := os.Getenv("PACKET_PROJECT")

	devList, _, err := dAPI.FindDevices(projectID, "", 0, 10)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range devList.Devices {
		fmt.Println(d)
	}

}
