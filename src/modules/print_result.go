package modules

import "fmt"

func PrintResult(input Input, result []Connection, networkChoice string) {

	var network string

	if networkChoice == "-u" {
		network = "UDP"
	} else {
		network = "TCP"
	}

	fmt.Printf("\n\t\t  %v %v\n\n", input.host, network)
	fmt.Printf("\t\tPORT\t\t OPEN\n")
	fmt.Println("-----------------------------------------------------")

	for _, el := range result {
		var open string
		if el.isOpen {
			open = "Yes"
		} else {
			open = "No"
		}

		fmt.Printf("\t\t%v\t\t %v\n", el.port, open)
	}
	fmt.Println()
}
