//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func info(list map[string]int) {
	countOnline := 0
	countOffline := 0
	countMaintenance := 0
	countRetired := 0
	for key, value := range list {
		switch value {
		case 0:
			fmt.Println("The status of", key, "is online")
			countOnline += 1
		case 1:
			fmt.Println("The status of", key, "is offline")
			countOffline += 1
		case 2:
			fmt.Println("The status of", key, "is maintance")
			countMaintenance += 1
		case 3:
			fmt.Println("The status of", key, "is retired")
			countRetired += 1
		}
	}
	fmt.Println("The count servers online is", countOnline)
	fmt.Println("The count servers offline is", countOffline)
	fmt.Println("The count servers in maintenance is", countMaintenance)
	fmt.Println("The count servers retired is", countRetired)
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}
	fmt.Println("--------Default Status--------")
	info(serverStatus)
	serverStatus["darkstar"] = 3
	serverStatus["aiur"] = 1
	fmt.Println("--------Change some Status--------")
	info(serverStatus)
	for _, server := range servers {
		serverStatus[server] = Maintenance
	}
	fmt.Println("--------Final Status--------")
	info(serverStatus)
}
