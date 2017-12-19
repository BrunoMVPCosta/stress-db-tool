package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	MongoConnectionString string
	InsertMany            bool
	BulkInsert            bool
	BulkInsertWithSleep   bool
}

func ReadConfiguration(fileName string) *Configuration {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("Could not open configuration file. ", err)
		return nil
	}

	decoder := json.NewDecoder(file)
	configuration := &Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error reading configuration file. ", err)
	}
	return configuration
}

// PrintConfiguration prints the current configuration
func (c *Configuration) PrintConfiguration() {
	fmt.Println("CONFIGURATION")
	fmt.Println("================")
	fmt.Println("Mongo Connection: ", c.MongoConnectionString)
	fmt.Println("InsertMany: ", c.InsertMany)
	fmt.Println("BulkInsert: ", c.BulkInsert)
	fmt.Println("BulkInsertWithSleep: ", c.BulkInsertWithSleep)
	fmt.Println("================")
}
