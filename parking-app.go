package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func create(capacity int) []string {
	parking := make([]string, capacity)
	fmt.Printf("Created parking lot with %d slots\n", capacity)
	return parking
}

func park(parking []string, car_number string) {
	found := false
	for i := range parking {
		if parking[i] == "" {
			parking[i] = car_number
			fmt.Printf("Allocated slot number %d\n", i+1)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Sorry, parking lot is full")
	}
}

func leave(parking []string, car_number string, hours int) {
	found := false
	for i := range parking {
		if parking[i] == car_number {
			charge := 10
			if hours > 2 {
				charge += (hours - 2) * charge
			}
			fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", car_number, i+1, charge)
			parking[i] = ""
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Registration number %s not found\n", car_number)
	}
}

func status(parking []string) {
	fmt.Println("Slot No. \t Registration No.")
	for i := range parking {
		if parking[i] != "" {
			fmt.Printf("%d\t\t %s\n", i+1, parking[i])
		}
	}
}

func main() {
	var parking []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter commands (type 'exit' to quit):")

	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			break
		}

		inputs := strings.Fields(text)
		if len(inputs) == 0 {
			continue
		}

		command := inputs[0]
		if command == "create_parking_lot" && len(inputs) == 2 {
			capacity, _ := strconv.Atoi(inputs[1])
			parking = create(capacity)
		} else if command == "park" && len(inputs) == 2 {
			park(parking, inputs[1])
		} else if command == "leave" && len(inputs) == 3 {
			hours, _ := strconv.Atoi(inputs[2])
			leave(parking, inputs[1], hours)
		} else if command == "status" || command == "Status" {
			status(parking)
		} else {
			fmt.Println("Command not recognized!")
		}
	}

	fmt.Println("Finished parking app. Goodbye!")
}
