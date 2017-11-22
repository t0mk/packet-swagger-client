package main

import (
	"fmt"
	"log"
	"os"

	"github.com/t0mk/packswgo"
)

func main() {
	c := packswgo.NewClient()
	projectID := os.Getenv("PACKET_PROJECT")

	devList, _, err := c.Devices.FindDevices(projectID, "", 0, 10)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range devList.Devices {
		fmt.Println(d)
	}

	facList, _, err := c.Facilities.FindFacilities()
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range facList.Facilities {
		fmt.Println(f)
	}

}
