package company

type CompanyCreate struct {
	Name         string `json:"name"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
}
