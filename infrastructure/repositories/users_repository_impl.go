package repositories

import (
	"context"

	"gorm.io/gorm"
	"todoapp.com/domain/interfaces"
	"todoapp.com/domain/models"
)

type usersRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) interfaces.UsersRepository {
	return &usersRepositoryImpl{db: db}
}

func (ur *usersRepositoryImpl) GetAll(context context.Context) []models.User {
	users := &[]models.User{}

	ur.db.WithContext(context).Where("deleted_at is null").Find(&users)

	return *users
}

func (ur *usersRepositoryImpl) Get(context context.Context, id *uint) models.User {
	user := &models.User{}

	ur.db.WithContext(context).Where("id = ? and deleted_at is null", (*id)).Find(&user)

	return *user
}

func (ur *usersRepositoryImpl) Create(context context.Context, user *models.User) error {
	return ur.db.WithContext(context).Create(user).Error
}

func (ur *usersRepositoryImpl) Update(context context.Context, user *models.User) error {
	return ur.db.WithContext(context).Where("id = ? and deleted_at is null", (*user).ID).Updates(user).Error
}

func (ur *usersRepositoryImpl) Delete(context context.Context, user *models.User) error {
	return ur.db.WithContext(context).Delete(user).Error
}

// Integration tests only
func (ur *usersRepositoryImpl) CleanUp(context context.Context) int64 {
	return ur.db.Unscoped().WithContext(context).Where("name like ? or email like ?", "%[test]%", "%[test]%").Delete(&[]models.User{}).RowsAffected
}
