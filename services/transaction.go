package services

import (
	"context"
	"fmt"
	"mongodb/config"
	"mongodb/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func transactionContext()*mongo.Collection{

	client,_:=config.ConnectDataBase()

	return config.GetCollection(client,"sample_analytics","transactions")

}

func DisplayTransaction()([]*models.Transaction,error){
	ctx,_:=context.WithTimeout(context.Background(),100*time.Second)

	filter := bson.D{}

	result,err :=transactionContext().Find(ctx,filter)

	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}else {
		var transaction[]*models.Transaction
		for result.Next(ctx){
			var trans =&models.Transaction{}
			err := result.Decode(trans)
			if err!=nil{
				return nil,err
			}
		transaction = append(transaction, trans)

		// fmt.Println(trans)
		}
		if err:=result.Err();err !=nil{
			return nil,err
		}
		if len(transaction)==0{
			return []*models.Transaction{},nil
		}
		return transaction,nil

		}


	}
