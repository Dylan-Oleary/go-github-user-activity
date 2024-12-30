package main

import (
	"errors"
	"os"
)

func getUsernameFromCli() (string, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return "", errors.New("No username provided")
	}

	return args[0], nil
}
