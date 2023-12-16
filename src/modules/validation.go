package modules

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Input struct {
	host  string
	ports []string
}

func ValidateArgs(args []string) (Input, error) {

	var input Input
	//option flag

	if len(args) < 2 {
		return input, errors.New("Missing option, please provide a valid option flag")
	}

	if args[1] != "-u" && args[1] != "-t" && args[1] != "--help" {
		return input, errors.New("Please provide a valid option flag")
	}

	//exeption --help

	if args[1] == "--help" {
		return input, nil
	}

	//host flag

	if len(args) < 3 {
		return input, errors.New("Missing host, please provide a valid host")
	}

	if !strings.Contains(args[2], ".") || !is_numeric(args[2]) || is_alphabetic(args[2]) {
		return input, errors.New("Please provide a valid format host IP address")
	}

	input.host = args[2]

	//port flag

	if len(args) < 4 {
		return input, errors.New("Missing port | port-port range, please provide a port or a port range")
	}

	if args[3] != "-p" {
		return input, errors.New("Missing port flag, please provide a valid port flag")
	}

	//port

	if len(args) < 5 {
		return input, errors.New("Missing valid port or port range.")
	}

	if is_alphabetic(args[4]) || !is_numeric(args[4]) {
		return input, errors.New("Please provide a valid port or a port range")
	}

	if strings.Contains(args[4], "-") {

		port1, _ := strconv.Atoi(strings.Split(args[4], "-")[0])
		port2, _ := strconv.Atoi(strings.Split(args[4], "-")[1])

		if port1 > port2 {
			return input, errors.New("Port range must be incremental")
		}

		if port1 == port2 {
			return input, errors.New("Port ranges much differ")
		}

		if port2 > 65535 {
			return input, errors.New("Max port reached, try lowering your port number")
		}

		for i := port1; i <= port2; i++ {
			str := strconv.Itoa(i)
			input.ports = append(input.ports, str)
		}
	} else {

		port, _ := strconv.Atoi(args[4])

		if port > 65535 {
			return input, errors.New("Max port reached, try lowering your port number")
		}

		input.ports = append(input.ports, args[4])
	}
	return input, nil

}

func is_numeric(word string) bool {
	return regexp.MustCompile(`\d`).MatchString(word)
	// calling regexp.MustCompile() function to create the regular expression.
	// calling MatchString() function that returns a bool that
	// indicates whether a pattern is matched by the string.
}

func is_alphabetic(word string) bool {
	return regexp.MustCompile(`[A-Za-z]`).MatchString(word)
	// calling regexp.MustCompile() function to create the regular expression.
	// calling MatchString() function that returns a bool that
	// indicates whether a pattern is matched by the string.
}
