package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhinliang/gosimplifier"
	"muzzammil.xyz/jsonc"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "c", "", "Path to configuration file.")
	flag.Parse()
}

func parseConfigFile() string {
	if configFile != "" {
		return configFile
	}

	if _, err := os.Stat(".simplex.json"); err == nil {
		return ".simplex.json"
	} else if _, err := os.Stat(".simplex.jsonc"); err == nil {
		return ".simplex.jsonc"
	} else if _, err := os.Stat(filepath.Join(os.Getenv("HOME"), ".simplex.json")); err == nil {
		return filepath.Join(os.Getenv("HOME"), ".simplex.json")
	} else if _, err := os.Stat(filepath.Join(os.Getenv("HOME"), ".simplex.jsonc")); err == nil {
		return filepath.Join(os.Getenv("HOME"), ".simplex.jsonc")
	}
	fmt.Println("No configuration file found. Please provide one using the -c option.")
	os.Exit(1)
	return ""
}

func main() {
	// Config file handling
	configFilePath := parseConfigFile()

	// Load configuration file
	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Error reading configuration file: %s\n", err)
		os.Exit(1)
		return
	}

	// If it is a JSONC file, convert it to JSON
	if filepath.Ext(configFilePath) == ".jsonc" {
		configData = jsonc.ToJSON(configData)
	}

	// Create new simplifier
	simplifier, err := gosimplifier.NewSimplifier(string(configData))
	if err != nil {
		fmt.Printf("Error creating simplifier: %s\n", err)
		os.Exit(1)
		return
	}

	// Create a new Scanner that scans os.Stdin.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		// Attempt to unmarshal input JSON as a map
		var original map[string]interface{}
		if err := json.Unmarshal([]byte(input), &original); err != nil {
			// If that fails, attempt to unmarshal as an array
			var originalArr []interface{}
			err = json.Unmarshal([]byte(input), &originalArr)
			if err != nil {
				// If that also fails, exit
				fmt.Printf("Error decoding JSON: %s\n", err)
				os.Exit(1)
				return
			}

			// Simplify each object in the array
			var simplifiedArr []interface{}
			for _, obj := range originalArr {
				simplified, err := simplifier.Simplify(obj)
				if err != nil {
					fmt.Printf("Error simplifying JSON: %s\n", err)
					os.Exit(1)
				}
				simplifiedArr = append(simplifiedArr, simplified)
			}
			// Output the simplified array
			simplifiedJSON, err := json.Marshal(simplifiedArr)
			if err != nil {
				fmt.Printf("Error encoding JSON: %s\n", err)
				os.Exit(1)
			}
			fmt.Println(string(simplifiedJSON))
			continue
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
