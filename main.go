package main

import (
	"context"
	"fmt"
	"log"

	"example.com/m/v2/controller"
	"example.com/m/v2/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	us          *services.UserServiceImpl
	uc          *controller.UserController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("Users").Collection("Users")
	us = services.NewUserService(userc, ctx)
	uc = controller.NewController(us)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1/")
	uc.RegisterRouterGroup(basepath)

	log.Fatal(server.Run(":9000"))

}