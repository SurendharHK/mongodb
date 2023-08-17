package services

import (
	"context"
	"fmt"
	"mongodb/config"
	"mongodb/models"
	"time"

	// "9fans.net/go/plan9/client"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func productContext() *mongo.Collection {
	client, _ := config.ConnectDataBase()
	return config.GetCollection(client, "inventory", "products")
}

func restaurantContext() *mongo.Collection {
	client, _ := config.ConnectDataBase()
	return config.GetCollection(client, "sample_restaurants", "restaurants")
}

// func InsertProduct(product models.Product) (*mongo.InsertOneResult, error) {
// 	var product models.Product
// 	product.ID = primitive.NewObjectID()
// 	product.Name = "Iphone"
// 	product.Price = 115000
// 	product.Description = "It is an awesome phone"
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, _ := config.ConnectDataBase()
// 	var productCollection *mongo.Collection = config.GetCollection(client, "inventrory", "products")
// 	result, err := productContext().InsertOne(ctx, product)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// 	return result, nil
// }

func InsertProductList(products []interface{}) (*mongo.InsertManyResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := productContext().InsertMany(ctx, products)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil

}
// func FindProducts() ([]*models.Product, error) {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	filter := bson.D{{"name", "OnePlus"}}
// 	var products []models.Product
// 	result, err := productContext().Find(ctx, filter)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	} else {
// 		//do something
// 		fmt.Println(result)
// 		//build the array of products for the cursor that we received
// 		var products []*models.Product
// 		for result.Next(ctx) {
// 			product := &models.Product{}
// 			err := result.Decode(product)
// 			if err != nil {
// 				return nil, err
// 			}
// 			//fmt.Println(product)
// 			products = append(products, product)
// 		}
// 		if err := result.Err(); err != nil {
// 			return nil, err
// 		}
// 		if len(products) == 0 {
// 			return []*models.Product{}, err
// 		}
// 		return products, nil
// 	}

// }

func DisplayContent() ([]*models.Restaurants, error) {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.D{}
	//var products []models.Product
	result, err := restaurantContext().Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		//do something
		fmt.Println(result)
		//build the array of products for the cursor that we received
		var products []*models.Restaurants
		for result.Next(ctx) {
			// fmt.Println("inside cursor:",len(products))
			product := &models.Restaurants{}
			err := result.Decode(product)
			if err != nil {
				return nil, err
			}
			//fmt.Println(product)
			products = append(products, product)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(products) == 0 {
			// fmt.Println("length of products",len(products))
			return []*models.Restaurants{}, err
		}
		// fmt.Println("length of products",len(products))
		return products, nil
	}

}
