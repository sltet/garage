package company

import (
	"github.com/sltet/garage/app/core"
)

type Company struct {
	core.Entity
	Name         string `json:"name"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	State        string `json:"state"`
	PostalCode   string `json:"postal_code"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Website      string `json:"website"`
} //@name Company

func NewCompany(o CompanyCreate) Company {
	return Company{
		Entity:       core.NewEntity(),
		Name:         o.Name,
		AddressLine1: o.AddressLine1,
		AddressLine2: o.AddressLine2,
		State:        o.State,
		PostalCode:   o.PostalCode,
		City:         o.City,
		Country:      o.Country,
		Phone:        o.Phone,
		Email:        o.Email,
		Website:      o.Website,
	}
}
