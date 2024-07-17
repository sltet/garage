package user

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
	FindAll(ctx *gin.Context) ([]User, core.DetailedError)
	FindById(ctx *gin.Context, id string) (User, core.DetailedError)
	Create(ctx *gin.Context, u User) (User, core.DetailedError)
	Save(ctx *gin.Context, u User) (User, core.DetailedError)
}

func NewRepository(db db.EntityManagerInterface, crud db.CrudRepositoryInterface) *Repository {
	return &Repository{crud, db}
}

func (r Repository) FindAll(ctx *gin.Context) ([]User, core.DetailedError) {
	var users []User

	// Get all records
	err := r.db.Database().Find(&users).WithContext(ctx).Error
	if err != nil {
		return nil, core.NewDatabaseError(err)
	}

	return users, nil
}

func (r Repository) FindById(ctx *gin.Context, id string) (User, core.DetailedError) {
	var user User
	err := r.crudRepository.Read(ctx, id, &user)
	if err != nil {
		return user, core.NewDatabaseError(err)
	}
	return user, nil
}

func (r Repository) Create(ctx *gin.Context, u User) (User, core.DetailedError) {
	err := r.crudRepository.Create(ctx, &u)
	if err != nil {
		return u, core.NewDatabaseError(err)
	}
	return u, nil
}

func (r Repository) Save(ctx *gin.Context, u User) (User, core.DetailedError) {
	var user User
	err := r.db.Database().Save(u).WithContext(ctx).Error
	if err != nil {
		return user, core.NewDatabaseError(err)
	}
	return u, nil
}
