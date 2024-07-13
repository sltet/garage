package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/db"
)

type Repository struct {
	crudRepository db.CrudRepositoryInterface
	db             db.EntityManagerInterface
}

type RepositoryInterface interface {
	FindAll(ctx *gin.Context) ([]User, error)
	FindById(ctx *gin.Context, id string) (User, error)
	Create(ctx *gin.Context, u User) (User, error)
	Save(ctx *gin.Context, u User) (User, error)
}

func NewRepository(db db.EntityManagerInterface, crud db.CrudRepositoryInterface) *Repository {
	return &Repository{crud, db}
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

func (r Repository) FindById(ctx *gin.Context, id string) (User, error) {
	var user User
	err := r.crudRepository.Read(ctx, id, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r Repository) Create(ctx *gin.Context, u User) (User, error) {
	err := r.crudRepository.Create(ctx, &u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r Repository) Save(ctx *gin.Context, u User) (User, error) {
	var user User
	err := r.db.Database().Save(u).WithContext(ctx).Error
	if err != nil {
		return user, err
	}
	return u, nil
}
