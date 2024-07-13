package vehicle

import (
	"github.com/sltet/garage/app/core"
)

type Vehicle struct {
	core.Entity
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func NewVehicle(make string, model string, year int) Vehicle {
	return Vehicle{
		core.NewEntity(),
		make,
		model,
		year,
	}
}
