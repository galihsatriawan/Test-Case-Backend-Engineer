package user

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindByID(ID int) (User, error)
	FindByUsername(username string) (User, error)
	Create(user User) (User, error)
	Save(user User) (User, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindByUsername(username string) (User, error) {
	var user User
	err := repo.db.Where("username = ?", username).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
