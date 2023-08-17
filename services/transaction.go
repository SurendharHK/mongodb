package services

import (
	"context"
	"encoding/json"
	"fmt"
	"mongodb/config"

	// "mongodb/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func transactionContext() *mongo.Collection {

	client, _ := config.ConnectDataBase()

	return config.GetCollection(client, "sample_analytics", "transactions")

}

// func DisplayTransaction()([]*models.Transaction,error){
// 	ctx,_:=context.WithTimeout(context.Background(),100*time.Second)

// 	filter := bson.M{"transaction_count":bson.D{{"$eq",100},{"account_id":bson.D{"$gt",7000}}}}
// 	options := options.Find().SetSort(bson.D{{"transaction_count",1}}).SetSkip(30).SetLimit(2)

// 	result,err :=transactionContext().Find(ctx,filter,options)

// 	if err!=nil{
// 		fmt.Println(err.Error())
// 		return nil,err
// 	}else {
// 		var transaction[]*models.Transaction
// 		for result.Next(ctx){
// 			var trans =&models.Transaction{}
// 			err := result.Decode(trans)
// 			if err!=nil{
// 				return nil,err
// 			}
// 		transaction = append(transaction, trans)

// 		// fmt.Println(trans)
// 		}
// 		if err:=result.Err();err !=nil{
// 			return nil,err
// 		}
// 		if len(transaction)==0{
// 			return []*models.Transaction{},nil
// 		}
// 		return transaction,nil

// 		}

// 	}

func FetchAggregatedTransaction() {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	matchStage := bson.D{{"$match", bson.D{{"account_id", 278866}}}}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$account_id"},
			{"total_count", bson.D{{"$sum", "$transaction_count"}}},
		}}}
	result, err := transactionContext().Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var showsWithInfo []bson.M
		if err = result.All(ctx, &showsWithInfo); err != nil {
			panic(err)
		}
		//fmt.Println(showsWithInfo)
		formatted_data, err := json.MarshalIndent(showsWithInfo, "", " ")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(formatted_data))
		}
	}
}

func UpdateTransaction(initalValue int,newValue int)(*mongo.UpdateResult,error){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	filter :=bson.D{{"account_id",initalValue}}
	update := bson.D{{"$set",bson.D{{"account_id",newValue}}}}
	result,err:=transactionContext().UpdateOne(ctx,filter,update)
	if err!=nil{
		fmt.Println(err.Error())
	}
	return result,nil

}
