// Package core implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// ReadFromFile returns its argument string reversed rune-wise left to right.
func ReadFromFile(filaName string, obj interface{}) {
	// read file
	data, err := ioutil.ReadFile(filaName)
	if err != nil {
		fmt.Print(err)
	}

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Print(err)
	}
}

// WriteData returns its argument string reversed rune-wise left to right.
func WriteData(fileName string, data interface{}) {
	var jsonData []byte
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(fileName, jsonData, 0644)
}

//  // Open our jsonFile
//  jsonFile, err := os.Open("users.json")
//  // if we os.Open returns an error then handle it
//  if err != nil {
// 	 fmt.Println(err)
//  }
//  fmt.Println("Successfully Opened users.json")
//  // defer the closing of our jsonFile so that we can parse it later on
//  defer jsonFile.Close()

//  byteValue, _ := ioutil.ReadAll(jsonFile)

//  var result map[string]interface{}
//  json.Unmarshal([]byte(byteValue), &result)

//  fmt.Println(result["users"])
