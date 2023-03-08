package genericRepo

import (
	logService "backend-skeleton-golang/commons/app/services/log-service"
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Generic[T interface{}] struct {
	DB *gorm.DB
}

func (g *Generic[T]) FindById(id string) (*T, error) {
	model := new(T)
	res := g.DB.First(model, "id = ?", id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model, nil
	}

	if res.Error != nil {
		logService.Error(res.Error.Error())
	}

	return model, res.Error
}

func (g *Generic[T]) FindOne(query interface{}) (*T, error) {
	model := new(T)
	res := g.DB.Where(query).First(&model)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model, nil
	}

	if res.Error != nil {
		logService.Error(res.Error.Error())
	}

	return model, res.Error
}

func (g *Generic[T]) Find(query interface{}) ([]T, error) {
	models := []T{}
	res := g.DB.Where(query).Find(&models)

	if res.Error != nil {
		logService.Error(res.Error.Error())
	}

	return models, res.Error
}

func (g *Generic[T]) Create(data interface{}) (*T, error) {
	model := new(T)
	copier.Copy(&model, data)
	res := g.DB.Create(&model)

	if res.Error != nil {
		logService.Error(res.Error.Error())
	}

	return model, res.Error
}

func (g *Generic[T]) UpdateById(id string, data interface{}) (*T, error) {
	model := new(T)
	copier.Copy(&model, data)
	res := g.DB.Model(&model).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&model)

	if res.Error != nil {
		logService.Error(res.Error.Error())
	}

	return model, res.Error
}

func (g *Generic[T]) FindWithNot(queryNot interface{}, query interface{}) (*T, error) {
	model := new(T)
	res := g.DB.Not(queryNot).Where(query).First(&model)

	if res.Error != nil {
		logService.Error(res.Error.Error())
	}

	return model, res.Error
}

func (g *Generic[T]) DeleteById(id string) (*T, error) {
	model := new(T)
	res := g.DB.Delete(&model, "id = ?", id)

	if res.Error != nil {
		logService.Error(res.Error.Error())
	}

	return model, res.Error
}
