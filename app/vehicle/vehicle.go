package vehicle

import (
	"github.com/sltet/garage/app/core"
)

type Vehicle struct {
	ID    string `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func NewVehicle(make string, model string, year int) Vehicle {
	return Vehicle{
		core.GetTimeBasedUUID().String(),
		make,
		model,
		year,
	}
}
