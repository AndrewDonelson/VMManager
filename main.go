package main

import "github.com/AndrewDonelson/VMManager/models"

func main() {
	dc, dcErr := models.NewDataCenter(models.TestDataFile)
	if dcErr != nil {
		panic(dcErr)
	}
	println(dc)
}
