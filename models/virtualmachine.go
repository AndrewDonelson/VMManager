package models

// VirtualMachine Model
type VirtualMachine struct {
	Name        string     `schema:"name" json:"name"`
	OSType      string     `schema:"os" json:"os"`
	IP          string     `schema:"ip" json:"ip"`
	Environment Enviroment `schema:"env" json:"env"`
}
