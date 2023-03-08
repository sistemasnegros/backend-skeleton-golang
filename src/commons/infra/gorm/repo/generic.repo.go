package genericRepo

// import (
// 	logService "backend-skeleton-golang/commons/app/services/log-service"
// 	gormInfra "backend-skeleton-golang/commons/infra/gorm"
// 	"errors"

// 	"github.com/jinzhu/copier"
// 	"gorm.io/gorm"
// )

// func FindById[ModelType comparable](model ModelType, id string) (ModelType, error) {
// 	res := gormInfra.Db.First(&model, "id = ?", id)

// 	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
// 		return model, nil
// 	}

// 	if res.Error != nil {
// 		logService.Error(res.Error.Error())
// 	}

// 	return model, res.Error
// }

// func FindOne[ModelType comparable](model ModelType, query interface{}) (ModelType, error) {
// 	res := gormInfra.Db.Where(query).First(&model)

// 	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
// 		return model, nil
// 	}

// 	if res.Error != nil {
// 		logService.Error(res.Error.Error())
// 	}

// 	return model, res.Error
// }

// func Find[ModelType comparable](model []ModelType, query interface{}) ([]ModelType, error) {
// 	res := gormInfra.Db.Where(query).Find(&model)

// 	if res.Error != nil {
// 		logService.Error(res.Error.Error())
// 	}

// 	return model, res.Error
// }

// func Create[ModelType comparable](model ModelType, data interface{}) (ModelType, error) {
// 	copier.Copy(&model, data)
// 	res := gormInfra.Db.Create(&model)

// 	if res.Error != nil {
// 		logService.Error(res.Error.Error())
// 	}

// 	return model, res.Error
// }

// func UpdateById[ModelType comparable](model ModelType, id string, data interface{}) (ModelType, error) {
// 	copier.Copy(&model, data)
// 	res := gormInfra.Db.Model(&model).Updates(&model)

// 	if res.Error != nil {
// 		logService.Error(res.Error.Error())
// 	}

// 	return model, res.Error
// }

// func FindWithNot[ModelType comparable](model ModelType, queryNot interface{}, query interface{}) (ModelType, error) {
// 	res := gormInfra.Db.Not(queryNot).Where(query).First(&model)

// 	if res.Error != nil {
// 		logService.Error(res.Error.Error())
// 	}

// 	return model, res.Error
// }

// func DeleteById[ModelType comparable](model ModelType, id string) (ModelType, error) {
// 	res := gormInfra.Db.Delete(&model, "id = ?", id)

// 	if res.Error != nil {
// 		logService.Error(res.Error.Error())
// 	}

// 	return model, res.Error
// }
