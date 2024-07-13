package servicerequest

import "github.com/sltet/garage/app/core"

type ServiceRequest struct {
	core.Entity
	Name string `json:"name"`
} //@name ServiceRequest

func NewServiceRequest(name string) ServiceRequest {
	return ServiceRequest{
		core.NewEntity(),
		name,
	}
}
