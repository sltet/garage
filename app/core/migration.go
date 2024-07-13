package core

import "gorm.io/gorm"

type Migration interface {
	Up(db *gorm.DB)
}
