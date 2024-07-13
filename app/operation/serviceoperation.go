package operation

import (
	"github.com/sltet/garage/app/core"
)

type ServiceOperation struct {
	core.Entity
	Name       core.LocalizedMessage `json:"name"`
	Operations []Operation           `json:"operations" gorm:"foreignKey:ServiceOperationId"`
}

func NewServiceOperation(name core.LocalizedMessage) ServiceOperation {
	return ServiceOperation{
		core.NewEntity(),
		name,
		[]Operation{},
	}
}

func (c *ServiceOperation) AddOperation(op Operation) {
	c.Operations = append(c.Operations, op)
}
