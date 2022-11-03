package main

import (
	"fmt"
	config "nid-go/cmd"
	"os"
)

func main() {
	var configPath string

	if len(os.Args) == 2 {
		configPath = os.Args[1]
	} else {
		configPath = "nid-go.yaml"
	}

	res, err := config.ParseConfigFile(configPath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v", res)
}
