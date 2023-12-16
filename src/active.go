package main

import (
	mod "active/modules"
	"fmt"
	"os"
)

func main() {
	//validate args

	input, err := mod.ValidateArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var result []mod.Connection

	switch os.Args[1] {
	case "-u":
		result = mod.UdpConn(input)
		mod.PrintResult(input, result, os.Args[1])

	case "-t":
		result = mod.TcpConn(input)
		mod.PrintResult(input, result, os.Args[1])
		
	case "--help":
		mod.PrintHelp()
	}

}