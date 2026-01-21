package main

import (
	"fmt"
	"os"

	"github.com/PwnySQL/bloggator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error while reading config: %v\n", err)
		os.Exit(1)
	}
	err = config.SetUser(&cfg, "PwnySQL")
	if err != nil {
		fmt.Printf("Error while updating config: %v\n", err)
		os.Exit(1)
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error while reading updated config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Updated config: %v\n", cfg)
}
