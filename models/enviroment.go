package models

// Enviroment is used as an enumerator VM Enviroments
type Enviroment int

const (
	// ENVDevelopment represents a Development enviroment
	ENVDevelopment Enviroment = iota
	// ENVTesting represents a Testing or QA enviroment
	ENVTesting
	// ENVProduction represents a Production enviroment
	ENVProduction
)

// String method returns the string representation of the Enviroment Value
func (d Enviroment) String() string {
	return [...]string{"Development", "Testing", "Production"}[d]
}
