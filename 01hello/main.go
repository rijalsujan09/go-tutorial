package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("iwlist", "wlan0", "scan")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	networks := parseIWListOutput(string(output))
	for _, network := range networks {
		fmt.Printf("SSID: %s\n", network)
	}
}

func parseIWListOutput(output string) []string {
	var networks []string

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "ESSID:") {
			// Extract SSID from the line
			ssid := strings.Split(line, ":")[1]
			ssid = strings.Trim(ssid, ` "`)
			networks = append(networks, ssid)
		}
	}

	return networks
}
