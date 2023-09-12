package repositories

import (
	"context"

	"gorm.io/gorm"
	"todoapp.com/domain/models"
	"todoapp.com/infrastructure/repositories/interfaces"
)

type todosRepositoryImpl struct {
	db *gorm.DB
}

func NewTodosRepository(db *gorm.DB) interfaces.TodosRepository {
	return &todosRepositoryImpl{db: db}
}

func (tr *todosRepositoryImpl) GetAll(context *context.Context) []models.Todo {
	todos := []models.Todo{}

	tr.db.WithContext(*context).Where("deleted_at is null").Find(todos)

	return todos
}

func (tr *todosRepositoryImpl) Create(context *context.Context, todo *models.Todo) error {
	return tr.db.WithContext(*context).Create(todo).Error
}

func (tr *todosRepositoryImpl) Update(context *context.Context, todo *models.Todo) error {
	return tr.db.WithContext(*context).Where("id = ? and deleted_at is null", todo.ID).Updates(todo).Error
}

func (tr *todosRepositoryImpl) Delete(context *context.Context, todo *models.Todo) error {
	return tr.db.WithContext(*context).Delete(todo).Error
}
