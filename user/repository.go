package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
	FindById(id int) (User, error)
	Store(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) FindById(id int) (User, error) {
	var user User
	err := r.db.Find(&user).Error

	return user, err
}

func (r *repository) Store(user User) (User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
