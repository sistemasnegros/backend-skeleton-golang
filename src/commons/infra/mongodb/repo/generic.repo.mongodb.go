package genericRepoMongodb

import (
	logService "backend-skeleton-golang/commons/app/services/log-service"
	serviceDomain "backend-skeleton-golang/commons/domain/service"
	"context"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IGeneric[T interface{}, D interface{}] interface {
	Find(query interface{}) ([]D, error)
	FindById(id string) (*D, error)
	FindOne(query interface{}) (*D, error)
	FindPagination(query interface{}, limit int64, page int) (*serviceDomain.PaginationData[D], error)
	UpdateById(id string, data interface{}) (*D, error)
	Create(D) (D, error)
	FindWithNot(queryNot map[string]interface{}, query map[string]interface{}) (*D, error)
	DeleteById(id string) error
}

type Generic[T interface{}, D interface{}] struct {
	DB *mongo.Collection
}

func ObjectIdToString(dataDomain interface{}, dataModel interface{}) {
	var fieldId interface{}

	// check if interface
	if reflect.ValueOf(dataModel).Kind() != reflect.Ptr {
		fieldId = reflect.ValueOf(dataModel).FieldByName("Id").Interface()
	}

	// check if struct
	if reflect.ValueOf(dataModel).Kind() == reflect.Ptr {
		fieldId = reflect.ValueOf(dataModel).Elem().FieldByName("Id").Interface()
	}

	// check if interface
	if reflect.ValueOf(dataDomain).Kind() != reflect.Ptr {
		reflect.ValueOf(dataDomain).FieldByName("Id").SetString(fieldId.(primitive.ObjectID).Hex())
		return
	}

	reflect.ValueOf(dataDomain).Elem().FieldByName("Id").SetString(fieldId.(primitive.ObjectID).Hex())
}

func (g *Generic[T, D]) FindById(id string) (*D, error) {
	model := new(T)
	idConverted, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: idConverted}}

	err := g.DB.FindOne(context.TODO(), filter).Decode(&model)

	dataDomain := new(D)
	copier.Copy(&dataDomain, model)

	if err == mongo.ErrNoDocuments {
		return dataDomain, nil
	}

	if err != nil {
		logService.Error(err.Error())
	}

	ObjectIdToString(dataDomain, model)
	return dataDomain, err
}

func (g *Generic[T, D]) FindOne(query interface{}) (*D, error) {
	model := new(T)

	err := g.DB.FindOne(context.TODO(), query).Decode(&model)

	dataDomain := new(D)
	copier.Copy(&dataDomain, model)

	if err == mongo.ErrNoDocuments {
		return dataDomain, nil
	}

	if err != nil {
		logService.Error(err.Error())
	}

	ObjectIdToString(dataDomain, model)
	return dataDomain, err
}

func (g *Generic[T, D]) Find(query interface{}) ([]D, error) {
	models := []T{}

	cursor, err := g.DB.Find(context.TODO(), query)

	if err != nil {
		logService.Error(err.Error())
		return nil, err
	}

	err = cursor.All(context.TODO(), &models)

	if err != nil {
		logService.Error(err.Error())
		return nil, err
	}
	var resultsDomain []D

	for _, result := range models {
		dataDomain := new(D)
		copier.Copy(&dataDomain, result)
		ObjectIdToString(dataDomain, result)

		resultsDomain = append(resultsDomain, *dataDomain)

	}

	return resultsDomain, err
}

func (g *Generic[T, D]) FindPagination(query interface{}, limit int64, page int) (*serviceDomain.PaginationData[D], error) {
	models := []T{}
	opts := options.Count().SetHint("_id_")
	count, err := g.DB.CountDocuments(context.TODO(), query, opts)

	pages := int(math.Ceil(float64(count) / float64(limit)))

	skip := limit * int64((page - 1))

	optsFind := options.Find().SetSkip(skip).SetLimit(limit)
	cursor, err := g.DB.Find(context.TODO(), query, optsFind)

	if err != nil {
		logService.Error(err.Error())
		return nil, err
	}

	err = cursor.All(context.TODO(), &models)

	if err != nil {
		logService.Error(err.Error())
		return nil, err
	}
	var resultsDomain []D

	for _, result := range models {
		dataDomain := new(D)
		copier.Copy(&dataDomain, result)
		ObjectIdToString(dataDomain, result)

		resultsDomain = append(resultsDomain, *dataDomain)

	}

	paginationData := &serviceDomain.PaginationData[D]{
		Pages:   pages,
		Page:    page,
		Data:    resultsDomain,
		Records: int(count),
	}

	return paginationData, err
}

func (g *Generic[T, D]) Create(data D) (D, error) {
	model := new(T)

	Id := reflect.ValueOf(data).FieldByName("Id")
	idConverted, _ := primitive.ObjectIDFromHex(Id.String())

	copier.Copy(&model, data)

	reflect.ValueOf(model).Elem().FieldByName("Id").Set(reflect.ValueOf(idConverted))
	reflect.ValueOf(model).Elem().FieldByName("CreatedAt").Set(reflect.ValueOf(time.Now()))
	reflect.ValueOf(model).Elem().FieldByName("UpdatedAt").Set(reflect.ValueOf(time.Now()))

	_, err := g.DB.InsertOne(context.TODO(), &model)

	if err != nil {
		logService.Error(err.Error())
	}

	copier.Copy(&data, model)

	return data, err
}

func (g *Generic[T, D]) UpdateById(id string, data interface{}) (*D, error) {
	model := new(T)
	copier.Copy(&model, data)

	idConverted, _ := primitive.ObjectIDFromHex(id)

	dataReflect := reflect.ValueOf(data)
	var dataBson bson.D

	for i := 0; i < dataReflect.NumField(); i++ {

		field := dataReflect.Type().Field(i).Tag.Get("json")
		value := dataReflect.Field(i).Interface()

		if field == "id" {
			continue
		}

		if value == "" {
			continue
		}

		dataBson = append(dataBson, bson.E{Key: field, Value: value})

	}

	dataBson = append(dataBson, bson.E{Key: "updateAt", Value: time.Now()})

	dataBsonSet := bson.D{{Key: "$set", Value: dataBson}}

	opts := options.Update().SetUpsert(true)
	_, err := g.DB.UpdateByID(context.TODO(), idConverted, dataBsonSet, opts)

	if err != nil {
		logService.Error(err.Error())
    return nil, err
	}

	dataDomain, _ := g.FindById(idConverted.Hex())

	return dataDomain, err
}

func (g *Generic[T, D]) FindWithNot(queryNot map[string]interface{}, query map[string]interface{}) (*D, error) {
	model := new(T)
	var queryBson bson.D

	for k, v := range query {
		queryBson = append(queryBson, bson.E{Key: k, Value: v})
	}

	for k, v := range queryNot {
		if k == "id" {
			k = "_id"
			objectId, _ := primitive.ObjectIDFromHex(fmt.Sprint(v))
			v = objectId
		}

		queryBson = append(queryBson, bson.E{Key: k, Value: bson.D{{Key: "$ne", Value: v}}})
	}

	err := g.DB.FindOne(context.TODO(), queryBson).Decode(&model)

	dataDomain := new(D)
	copier.Copy(&dataDomain, model)

	if err == mongo.ErrNoDocuments {
		return dataDomain, nil
	}

	if err != nil {
		logService.Error(err.Error())
	}

	ObjectIdToString(dataDomain, model)
	return dataDomain, err
}

func (g *Generic[T, D]) DeleteById(id string) error {
	idConverted, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: idConverted}}

	_, err := g.DB.DeleteOne(
		context.TODO(),
		filter,
	)

	if err != nil {
		logService.Error(err.Error())
	}

	return err
}
