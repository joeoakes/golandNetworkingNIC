package main

import (
	"fmt"
	"net"
)

func main() {
	// Get a list of all interfaces.
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate over all interfaces and print their details.
	for _, iface := range interfaces {
		fmt.Printf("Name: %v\n", iface.Name)
		fmt.Printf("MTU: %v\n", iface.MTU)
		fmt.Printf("Hardware Address: %v\n", iface.HardwareAddr)
		fmt.Printf("Flags: %v\n", iface.Flags)

		// Get addresses associated with the interface.
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, addr := range addrs {
			fmt.Printf("Address: %v\n", addr)
		}
		fmt.Println()
	}
}
