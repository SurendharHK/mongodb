package main

import (
	"fmt"
	// "mongodb/models"
	"mongodb/services"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
)

func main() {
	//client,_:=config.ConnectDataBase()
	//collection :=config.GetCollection(client,"sample_training","zips")
	fmt.Println("MongoDB successfully connected...")
	// product:= models.Product{ID: primitive.NewObjectID(),Name: "Samsung",Price: 1000000,Description: "Budget Phone"}
	// services.InsertProduct(product)
	// products := []interface{}{models.Product{ID: primitive.NewObjectID(), Name: "OnePlus", Price: 1000000, Description: "Budget Phone"},

	// 	models.Product{ID: primitive.NewObjectID(), Name: "Vivo", Price: 100000, Description: "China Phone"}}
	// services.InsertProductList(products)

	// product, _ := services.DisplayContent()

	// for index, result := range product {
	// 	fmt.Println(index, result)
	// }

	// transaction, _ := services.DisplayTransaction()

	// for index, result := range transaction {
	// 	fmt.Println(index, result)
	// }

	// services.FetchAggregatedTransaction()

	result,err := services.UpdateTransaction()

	

}
