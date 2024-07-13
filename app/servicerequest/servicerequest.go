package servicerequest

import "github.com/sltet/garage/app/core"

type ServiceRequest struct {
	core.Entity
	Name string `json:"name"`
}

func NewServiceRequest(name string) ServiceRequest {
	return ServiceRequest{
		core.NewEntity(),
		name,
	}
}
