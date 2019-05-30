package models

import (
	//"github.com/ghodss/yaml"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const TestDataFile = "./data/dc.yml"

type Datacenter struct {
}

// NewDataCenter given a valid YML Datacenter configuration file
// will create and return a new Datacenter object
func NewDataCenter(dataFile string) (*Datacenter, error) {
	dc := Datacenter{}

	_, err := os.Stat(dataFile)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	yerr := yaml.Unmarshal(data, dc)
	if yerr != nil {
		return nil, yerr
	}

	return &dc, nil
}
