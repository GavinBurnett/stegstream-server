package main

import (
	"fmt"
	"testing"
)

func TestReadConfigFile(t *testing.T) {

	// Set up test data
	var configData Config = Config{DEFAULT_PORT}

	// The tests to run
	var tests = []struct {
		name           string
		input          string
		expectedResult Config
	}{
		{"NoParameterData", "", configData},
		{"FileDoesNotExist", "fakeFile", configData},
		{"FileExists", CONFIG_FILE, configData},
	}

	// Write name of function being tested to test results file
	LogResult("ReadConfigFile")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := ReadConfigFile(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %d Expected: %d", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %d Expected: %d", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}
}

func TestCheckConfigFile(t *testing.T) {

	// Set up test data
	var defaultData Config = Config{DEFAULT_PORT}
	var portChanged Config = Config{8181}
	var invalidPort Config = Config{-1}

	// The tests to run
	var tests = []struct {
		name            string
		input           Config
		expectedConfig  Config
		expectedBoolean bool
	}{
		{"DefaultData", defaultData, defaultData, true},
		{"PortChanged", portChanged, portChanged, true},
		{"InvalidPort", invalidPort, defaultData, true},
	}

	// Write name of function being tested to test results file
	LogResult("CheckConfigFile")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			resultConfig, resultBoolean := CheckConfigFile(currentTest.input)

			if resultConfig != currentTest.expectedConfig && resultBoolean != currentTest.expectedBoolean {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %v Got: %d %v Expected: %d %v", currentTest.input, resultConfig, resultBoolean, currentTest.expectedConfig, currentTest.expectedBoolean) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %v Got: %d %v Expected: %d %v", currentTest.input, resultConfig, resultBoolean, currentTest.expectedConfig, currentTest.expectedBoolean) + " - PASS")
			}
		})
	}
}

func TestParseStringToInt(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		input          string
		expectedResult int
	}{
		{"NoParameterData", "", -1},
		{"InvalidData1", "invalidData", -1},
		{"InvalidData2", "34e", -1},
		{"InvalidData3", "e56", -1},
		{"ZeroNumber", "0", -1},
		{"MinusNumber", "-4", -1},
		{"10", "10", 10},
		{"8080", "8080", 8080},
	}

	// Write name of function being tested to test results file
	LogResult("ParseStringToInt")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := ParseStringToInt(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %d Expected: %d", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %d Expected: %d", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}
}