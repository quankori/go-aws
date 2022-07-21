package services

import (
	"fmt"
	"os"
)

// Hostname get the host name
func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return hostname
}
