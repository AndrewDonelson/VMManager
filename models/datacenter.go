package models

import (
	//"github.com/ghodss/yaml"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// TestDataFile is the file to be used for testing
const TestDataFile = "./data/dc.yml"

// Datacenter represents a datacenter with all its clusters, vm's etc
type Datacenter struct {
	// TODO: Convert YAML to JSON, Create Struct from JSON, Load YAML
}

// NewDataCenter given a valid YML Datacenter configuration file
// will create and return a new Datacenter object
func NewDataCenter(dataFile string) (*Datacenter, error) {
	dc := Datacenter{}

	// file exists & readable?
	_, err := os.Stat(dataFile)
	if err != nil {
		return nil, err
	}

	// load the file
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	// decode the file as yaml
	yerr := yaml.Unmarshal(data, dc)
	if yerr != nil {
		return nil, yerr
	}

	// return valid pointer datacenter object
	return &dc, nil
}
