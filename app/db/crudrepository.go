package db

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
	"gorm.io/gorm/clause"
)

type CrudRepository struct {
	EntityManager EntityManagerInterface
}

type CrudRepositoryInterface interface {
	Create(ctx *gin.Context, entity core.ORMAwareEntity) (err error)
	Read(ctx *gin.Context, entityId string, model interface{}) error
}

func NewCrudRepository(em EntityManagerInterface) *CrudRepository {
	return &CrudRepository{EntityManager: em}
}

func (r CrudRepository) Create(ctx *gin.Context, entity core.ORMAwareEntity) (err error) {
	err = r.EntityManager.Database().Create(entity).WithContext(ctx).Error
	if err != nil {
		return err
	}
	return nil
}

func (r CrudRepository) Read(ctx *gin.Context, entityId string, model interface{}) error {
	err := r.EntityManager.Database().Preload(clause.Associations).First(model, "id = ?", entityId).WithContext(ctx).Error
	if err != nil {
		return err
	}
	return nil
}
