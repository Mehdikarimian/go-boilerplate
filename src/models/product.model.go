package models

import (
	"context"
	"fmt"
	"go-boilerplate/src/common"
	"go-boilerplate/src/core/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProductsModel() *BaseModel {
	mod := &BaseModel{
		ModelConstructor: &common.ModelConstructor{
			Collection: db.GetMongoDb().Collection("ProductsCollection"),
		},
	}

	return mod
}

// models definitions
type Product struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `db:"title" json:"title"`
	Content   string             `db:"content" json:"content"`
	Price     int                `db:"price" json:"price"`
	UpdatedAt time.Time          `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time          `db:"created_at" json:"created_at"`
}

type CreateProductForm struct {
	Title   string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Content string `form:"content" json:"content" binding:"required,min=3,max=1000"`
	Price   int    `form:"price" json:"price" binding:"required"`
}

type UpdateProductForm struct {
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
	Price   int    `form:"price" json:"price"`
}

type FindProductForm struct {
	ID int64 `form:"id" json:"id" binding:"required"`
}

func (mod *BaseModel) GetAllProducts(limit int64, skip int64, search string) (products []Product, err error) {
	var results []Product
	opts := options.Find().SetLimit(limit).SetSkip(skip)

	filter := bson.D{{Key: "$or", Value: bson.A{bson.D{{Key: "title", Value: bson.M{"$regex": search}}}, bson.D{{Key: "content", Value: bson.M{"$regex": search}}}}}}

	cur, err := mod.Collection.Find(context.TODO(), filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}

	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	count, err := mod.Collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	return results, err
}

func (mod *BaseModel) GetOneProduct(id string) (product Product, err error) {
	var result Product
	objID, err := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objID}}
	err = mod.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}

	return result, err
}

func (mod *BaseModel) CreateProduct(body CreateProductForm) (productID interface{}, err error) {
	result, err := mod.Collection.InsertOne(context.TODO(), bson.M{"title": body.Title, "content": body.Content, "price": body.Price})

	if err != nil {
		return
	}

	return result.InsertedID, err
}

func (mod *BaseModel) UpdateProduct(id string, body UpdateProductForm) (result *mongo.UpdateResult, err error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{{Key: "$set", Value: bson.M{"title": body.Title, "content": body.Content, "price": body.Price}}}

	res, err := mod.Collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	return res, err
}

func (mod *BaseModel) DeleteProduct(id string) (result *mongo.DeleteResult, err error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}

	res, err := mod.Collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	return res, err
}
