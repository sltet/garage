package company

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/db"
)

type Repository struct {
	crudRepository db.CrudRepositoryInterface
	db             db.EntityManagerInterface
}

type RepositoryInterface interface {
	FindAll(ctx *gin.Context) ([]Company, error)
	Create(ctx *gin.Context, company Company) (comp Company, err error)
	FindById(ctx *gin.Context, id string) (company Company, err error)
}

func NewRepository(crud db.CrudRepositoryInterface, db db.EntityManagerInterface) *Repository {
	return &Repository{crud, db}
}

func (r Repository) FindAll(ctx *gin.Context) ([]Company, error) {
	var companies []Company

	// Get all records
	err := r.db.Database().Find(&companies).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (r Repository) Create(ctx *gin.Context, company Company) (comp Company, err error) {
	// Get all records
	err = r.db.Database().Create(company).WithContext(ctx).Error
	if err != nil {
		return comp, err
	}

	return company, nil
}

func (r Repository) FindById(ctx *gin.Context, id string) (company Company, err error) {
	err = r.crudRepository.Read(ctx, id, &company)
	if err != nil {
		return company, err
	}

	return company, nil
}
