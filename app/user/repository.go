package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/db"
)

type Repository struct {
	db db.DatabaseInterface
}

type RepositoryInterface interface {
	FindAll(ctx *gin.Context) ([]User, error)
}

func NewRepository(db db.DatabaseInterface) *Repository {
	return &Repository{db}
}

func (r Repository) FindAll(ctx *gin.Context) ([]User, error) {
	var users []User

	// Get all records
	err := r.db.Database().Find(&users).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
