package main

import "fmt"

func displayPrompt() {

	fmt.Println("\nWelcome to GoWhoAteMyCPU!")
	fmt.Println("")
	fmt.Println("More info on features:")
	fmt.Println("https://github.com/buildingBuild/GoWhoAteMyCPU")
	fmt.Println("")
	displayDeviceInfo(deviceInfo)

	fmt.Println(Green + "\nOPTIONS: " + Reset)
	fmt.Println("1. CPU Monitoring Mode:")
	fmt.Println("2. Memory Monitoring Mode:")
	fmt.Println("3. Network Monitoring Mode:")
	fmt.Println("5. Hybrid Monitoring Mode:")
	fmt.Print("Select a Mode: ")

}
