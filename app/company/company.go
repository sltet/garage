package company

import (
	"github.com/sltet/garage/app/core"
)

type Company struct {
	core.Entity
	Name         string `json:"name"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
}

func NewCompany(name, addressLine1, addressLine2 string) Company {
	return Company{
		Entity:       core.NewEntity(),
		Name:         name,
		AddressLine1: addressLine1,
		AddressLine2: addressLine2,
	}
}
