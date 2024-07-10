package company

import (
	"github.com/sltet/garage/app/core"
)

type Company struct {
	ID           string `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
}

func NewCompany(name, addressLine1, addressLine2 string) Company {
	return Company{
		ID:           core.GetTimeBasedUUID().String(),
		Name:         name,
		AddressLine1: addressLine1,
		AddressLine2: addressLine2,
	}
}
