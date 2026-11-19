package main

import (
	"fmt"
	"os"
)

var appVersion = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(appVersion)
		return
	}

	services := map[string]int{
		"ssh":      22,
		"http":     80,
		"postgres": 5432,
		"app":      8080,
	}

	var userServices []string
	var systemServices []string

	for name, port := range services {
		if port < 1024 {
			systemServices = append(systemServices, fmt.Sprintf("Service %s (port %d): privileged\n", name, port))
		} else {
			userServices = append(userServices, fmt.Sprintf("Service %s (port %d): user-space\n", name, port))
		}

	}

	fmt.Println(systemServices)
	fmt.Println(userServices)
}
