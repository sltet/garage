package operation

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
	FindAll(ctx *gin.Context) ([]ServiceOperation, core.DetailedError)
	FindById(ctx *gin.Context, id string) (ServiceOperation, core.DetailedError)
	Create(ctx *gin.Context, u ServiceOperation) (ServiceOperation, core.DetailedError)
	Save(ctx *gin.Context, u ServiceOperation) (ServiceOperation, core.DetailedError)
}

func NewRepository(db db.EntityManagerInterface, crud db.CrudRepositoryInterface) *Repository {
	return &Repository{crud, db}
}

func (r Repository) FindAll(ctx *gin.Context) ([]ServiceOperation, core.DetailedError) {
	var ops []ServiceOperation

	// Get all records
	err := r.db.Database().Preload("Operations").Find(&ops).WithContext(ctx).Error
	if err != nil {
		return nil, core.NewDatabaseError(err)
	}

	return ops, nil
}

func (r Repository) FindById(ctx *gin.Context, id string) (ServiceOperation, core.DetailedError) {
	var op ServiceOperation
	err := r.crudRepository.Read(ctx, id, &op)
	if err != nil {
		return op, core.NewDatabaseError(err)
	}
	return op, nil
}

func (r Repository) Create(ctx *gin.Context, op ServiceOperation) (ServiceOperation, core.DetailedError) {
	err := r.crudRepository.Create(ctx, &op)
	if err != nil {
		return op, core.NewDatabaseError(err)
	}
	return op, nil
}

func (r Repository) Save(ctx *gin.Context, u ServiceOperation) (ServiceOperation, core.DetailedError) {
	var op ServiceOperation
	err := r.db.Database().Save(u).WithContext(ctx).Error
	if err != nil {
		return op, core.NewDatabaseError(err)
	}
	return u, nil
}
