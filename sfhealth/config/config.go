package config

import (
	"bufio"
	"context"
	"errors"
	"log"
	"os"
	"strings"
)

//TODO: This is a demo code.
// Put more needed details here and in corresponding data file as per the need

var configFile = "./config/config.dat"
var config *Config

//Config provides the information from the application configuration file
type Config struct {
	Driver  string
	ConnStr string
	//TODO: should have all components separate in it from the conn string
}

// Initialize Iniatializes the configuration
func Initialize(ctx context.Context) error {

	cfg, err := getConfig(ctx, configFile)
	if err != nil {
		log.Fatalf("Could not initialize config %v", err)
		return err
	}

	config = cfg
	return nil
}

// getConfig loads the Config struct from config file
func getConfig(ctx context.Context, cfgFilePath string) (*Config, error) {

	log.Printf("Reading values from %v", cfgFilePath)
	values, err := readValues(cfgFilePath)
	if err != nil {
		log.Fatalf("Unable to read values %v", err)
		return &Config{}, err
	}

	driver, ok := values["dbDriver"]
	if !ok {
		log.Fatalf("driver key does not exist in %s", cfgFilePath)
		return &Config{}, errors.New("unable to get driver value")
	}

	connStr, ok := values["dbConnStr"]
	if !ok {
		log.Fatalf("dbConnStr key does not exist in %s", cfgFilePath)
		return &Config{}, errors.New("unable to get dbConnStr value")
	}

	return &Config{driver, connStr}, nil
}

//readValues reads the info from config file and returns as map
func readValues(propertyFileName string) (map[string]string, error) {
	file, err := os.Open(propertyFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	toReturn := make(map[string]string, len(lines))

	for _, line := range lines {
		currLine := strings.TrimSpace(line)

		//Improvement work
		//Can be made more generic with different options for delim
		if delimIdx := strings.Index(currLine, "="); delimIdx >= 0 {
			key := strings.TrimSpace(currLine[:delimIdx])
			value := strings.TrimSpace(currLine[delimIdx+1:])

			toReturn[key] = value
		}
	}

	return toReturn, nil
}

// Get Provides the initialized config
func Get() *Config {
	return config
}

// Set Just writting it but may not use
func Set(newConfig *Config) {
	config = newConfig
}
