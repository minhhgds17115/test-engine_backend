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
	as          *services.AnswerServiceImpl
	qs          *services.QuestionsServiceImpl
	ts          *services.TestServiceImpl
	uc          *controller.UserController
	qc          *controller.QuestionsController
	ac          *controller.AnswerController
	tc          *controller.TestController
	ctx         context.Context
	userc       *mongo.Collection
	questionc   *mongo.Collection
	answerc     *mongo.Collection
	testc       *mongo.Collection
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

	userc = mongoclient.Database("test-engine").Collection("User")
	questionc = mongoclient.Database("test-engine").Collection("Question")
	answerc = mongoclient.Database("test-engine").Collection("Answer")
	testc = mongoclient.Database("test-engine").Collection("Test")
	us = services.NewUserService(userc, ctx)
	as = services.NewAnswerServices(answerc, ctx)
	qs = services.NewQuestionsServices(questionc, ctx)
	ts = services.NewTestService(testc, ctx)
	uc = controller.NewController(us)
	qc = controller.NewQuestionsController(qs)
	ac = controller.NewAnswerController(as)
	tc = controller.NewTestController(ts)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1/")
	uc.RegisterRouterGroup(basepath)
	qc.RegisterQuestionsRouterGroup(basepath)
	ac.RegisterAnswerRouterGroup(basepath)
	tc.RegisterTestRouterGroup(basepath)

	log.Fatal(server.Run(":9000"))

}
