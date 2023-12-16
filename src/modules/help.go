package modules

import "fmt"

func PrintHelp() {
	fmt.Println("\nUSAGE: go run active.go [OPTIONS] [HOST] [PORT | PORT-PORT]")
	fmt.Println("Options:")
	fmt.Println("  -u\t\t UDP Scan")
	fmt.Println("  -t\t\t TCP Scan")
	fmt.Println("  --help\t Show this message and exit.")
	fmt.Println()
}
