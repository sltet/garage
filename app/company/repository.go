package company

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
	"github.com/sltet/garage/app/db"
)

type Repository struct {
	crudRepository db.CrudRepositoryInterface
	db             db.EntityManagerInterface
}

type RepositoryInterface interface {
	FindAll(ctx *gin.Context) ([]Company, core.DetailedError)
	Create(ctx *gin.Context, company Company) (comp Company, err core.DetailedError)
	FindById(ctx *gin.Context, id string) (company Company, err core.DetailedError)
}

func NewRepository(crud db.CrudRepositoryInterface, db db.EntityManagerInterface) *Repository {
	return &Repository{crud, db}
}

func (r Repository) FindAll(ctx *gin.Context) ([]Company, core.DetailedError) {
	var companies []Company

	// Get all records
	err := r.db.Database().Find(&companies).WithContext(ctx).Error
	if err != nil {
		return nil, core.NewDatabaseError(err)
	}

	return companies, nil
}

func (r Repository) Create(ctx *gin.Context, company Company) (Company, core.DetailedError) {
	err := r.crudRepository.Create(ctx, &company)
	if err != nil {
		return company, core.NewDatabaseError(err)
	}
	return company, nil
}

func (r Repository) FindById(ctx *gin.Context, id string) (Company, core.DetailedError) {
	var company Company
	err := r.crudRepository.Read(ctx, id, &company)
	if err != nil {
		return company, core.NewDatabaseError(err)
	}

	return company, nil
}
