package operation

import (
	"github.com/sltet/garage/app/core"
	"gorm.io/gorm"
)

type Migration001 struct{}

func (m Migration001) Up(db *gorm.DB) {
	serviceOperationName := core.NewLocalizedMessage()
	serviceOperationName.Add(core.EN, "Vehicle Maintenance Services")
	serviceOperationName.Add(core.FR, "Service de maintenance")
	so1 := NewServiceOperation(serviceOperationName)

	db.Session(&gorm.Session{SkipHooks: false}).Create(&so1)

	operationName := core.NewLocalizedMessage()
	operationName.Add(core.FR, "Changement d'huile")
	operationName.Add(core.EN, "Oil Change")
	operationDescription := core.NewLocalizedMessage()
	operationDescription.Add(core.EN, "Regular oil changes to ensure engine health")
	operation := NewOperation(operationName, so1.GetID(), operationDescription)
	so1.AddOperation(operation)

	operationName = core.NewLocalizedMessage()
	operationName.Add(core.FR, "Verification des fluides")
	operationName.Add(core.EN, "Fluid Checks and Refills")
	operationDescription = core.NewLocalizedMessage()
	operationDescription.Add(core.EN, "Checking and refilling fluids like brake fluid, coolant, and transmission fluid.")
	operation = NewOperation(operationName, so1.GetID(), operationDescription)
	so1.AddOperation(operation)

	operationName = core.NewLocalizedMessage()
	operationName.Add(core.FR, "Remplacement de filtre")
	operationName.Add(core.EN, "Filter Replacement")
	operationDescription = core.NewLocalizedMessage()
	operationDescription.Add(core.EN, "Replacing air filters, fuel filters, and cabin filters.")
	operation = NewOperation(operationName, so1.GetID(), operationDescription)
	so1.AddOperation(operation)

	operationName = core.NewLocalizedMessage()
	operationName.Add(core.FR, "Pneumatique")
	operationName.Add(core.EN, "Tire Services")
	operationDescription = core.NewLocalizedMessage()
	operationDescription.Add(core.EN, "Tire replacement, rotation, balancing, and alignment.")
	operationDescription.Add(core.FR, "Remplacement de pneu, alignement, balancement.")
	operation = NewOperation(operationName, so1.GetID(), operationDescription)
	so1.AddOperation(operation)

	// OPERATION 2

	db.Session(&gorm.Session{SkipHooks: false}).Save(&so1)
	serviceOperationName = core.NewLocalizedMessage()
	serviceOperationName.Add(core.EN, "Vehicle Repair Services")
	serviceOperationName.Add(core.FR, "Service de réparation")
	so1 = NewServiceOperation(serviceOperationName)

	db.Session(&gorm.Session{SkipHooks: false}).Create(&so1)

	operationName = core.NewLocalizedMessage()
	operationName.Add(core.FR, "Réparation moteur")
	operationName.Add(core.EN, "Engine Repair")
	operationDescription = core.NewLocalizedMessage()
	operationDescription.Add(core.EN, "Diagnosing and fixing engine-related issues.")
	operation = NewOperation(operationName, so1.GetID(), operationDescription)
	so1.AddOperation(operation)

	operationName = core.NewLocalizedMessage()
	operationName.Add(core.FR, "Réparation de la transmission")
	operationName.Add(core.EN, "Transmission Repair")
	operationDescription = core.NewLocalizedMessage()
	operationDescription.Add(core.EN, "Repairing or replacing transmission systems.")
	operation = NewOperation(operationName, so1.GetID(), operationDescription)
	so1.AddOperation(operation)

	db.Session(&gorm.Session{SkipHooks: false}).Save(&so1)

	//
	//so2 := NewServiceOperation("Vehicle Repair Services")
	//db.Session(&gorm.Session{SkipHooks: false}).Create(&so2)
	//so2.AddOperation(NewOperation("Engine Repair", so2.GetID(), "Diagnosing and fixing engine-related issues."))
	//so2.AddOperation(NewOperation("Transmission Repair", so2.GetID(), "Repairing or replacing transmission systems."))
	//so2.AddOperation(NewOperation("Brake Repair", so2.GetID(), "Repairing or replacing brake pads, rotors, and calipers."))
	//so2.AddOperation(NewOperation("Suspension Repair", so2.GetID(), "Fixing issues with shocks, struts, and other suspension components."))
	//so2.AddOperation(NewOperation("Exhaust System Repair", so2.GetID(), "Repairing or replacing exhaust systems and mufflers"))
	//so2.AddOperation(NewOperation("Electrical System Repair", so2.GetID(), "Fixing issues with the vehicle’s electrical system, including battery replacement and alternator repair."))
	//db.Session(&gorm.Session{SkipHooks: false}).Save(so2)
}
