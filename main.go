package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhinliang/gosimplifier"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "c", "", "Path to configuration file.")
	flag.Parse()
}

func main() {
	if configFile == "" {
		if _, err := os.Stat(".simplex.json"); err == nil {
			configFile = ".simplex.json"
		} else if _, err := os.Stat(filepath.Join(os.Getenv("HOME"), ".simplex.json")); err == nil {
			configFile = filepath.Join(os.Getenv("HOME"), ".simplex.json")
		} else {
			fmt.Println("No configuration file found. Please provide one using the -c option.")
			os.Exit(1)
		}
	}

	// Load configuration file
	config, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error reading configuration file: %s\n", err)
		os.Exit(1)
	}

	// Create new simplifier
	simplifier, err := gosimplifier.NewSimplifier(string(config))
	if err != nil {
		fmt.Printf("Error creating simplifier: %s\n", err)
		os.Exit(1)
	}

	// Create a new Scanner that scans os.Stdin.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		// Unmarshal input JSON
		var original map[string]interface{}
		err := json.Unmarshal([]byte(input), &original)
		if err != nil {
			fmt.Printf("Error decoding JSON: %s\n", err)
			os.Exit(1)
		}

		// Simplify JSON
		simplified, err := simplifier.Simplify(original)
		if err != nil {
			fmt.Printf("Error simplifying JSON: %s\n", err)
			os.Exit(1)
		}

		// Marshal JSON back to string
		simplifiedJSON, err := json.Marshal(simplified)
		if err != nil {
			fmt.Printf("Error encoding JSON: %s\n", err)
			os.Exit(1)
		}

		// Output simplified JSON
		fmt.Println(string(simplifiedJSON))
	}
	if scanner.Err() != nil {
		fmt.Printf("Scanner error: %s\n", scanner.Err())
		os.Exit(1)
	}
}
