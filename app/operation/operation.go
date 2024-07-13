package operation

import (
	"github.com/sltet/garage/app/core"
)

type Operation struct {
	core.Entity
	ServiceOperationId string                `json:"service_operation_id"`
	Name               core.LocalizedMessage `json:"name"`
	Description        core.LocalizedMessage `json:"description"`
}

func NewOperation(name core.LocalizedMessage, serviceOperationId string, description core.LocalizedMessage) Operation {
	return Operation{
		core.NewEntity(),
		serviceOperationId,
		name,
		description,
	}
}
