package operation

import (
	"github.com/sltet/garage/app/core"
	"gorm.io/gorm"
)

type Migration001 struct{}

func (m Migration001) Up(db *gorm.DB) {
	loc := core.NewLocalizedMessage()
	loc.Add(core.EN, "Vehicle Maintenance Services")
	loc.Add(core.FR, "Service de maintenance")
	so1 := NewServiceOperation(loc)

	db.Session(&gorm.Session{SkipHooks: false}).Create(&so1)
	loc = core.NewLocalizedMessage()
	loc.Add(core.FR, "Changement d'huile")
	loc.Add(core.EN, "Oil Change")

	des := core.NewLocalizedMessage()
	des.Add(core.EN, "Regular oil changes to ensure engine health")
	so1.AddOperation(NewOperation(loc, so1.GetID(), des))
	//so1.AddOperation(NewOperation("Fluid Checks and Refills", so1.GetID(), "Checking and refilling fluids like brake fluid, coolant, and transmission fluid."))
	//so1.AddOperation(NewOperation("Filter Replacement", so1.GetID(), "Replacing air filters, fuel filters, and cabin filters."))
	//so1.AddOperation(NewOperation("Tire Services", so1.GetID(), "Tire rotation, balancing, and alignment."))
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
