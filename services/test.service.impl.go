package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"example.com/m/v2/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestServiceImpl struct {
	testCollection *mongo.Collection
	ctx            context.Context
}

func NewTestService(testCollection *mongo.Collection, ctx context.Context) *TestServiceImpl {
	return &TestServiceImpl{
		testCollection: testCollection,
		ctx:            ctx,
	}
}

func (t *TestServiceImpl) GetAllTest() ([]*models.Test, error) {
	var tests []*models.Test
	cursor, err := t.testCollection.Find(t.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(t.ctx) {
		var test bson.D

		err := cursor.Decode(&test)
		if err != nil {
			fmt.Println("this err ", err)
			return nil, err
		}

		fmt.Println("test ", test)

		data, err := bson.Marshal(test)
		if err != nil {
			fmt.Println("this err1 ", err)
			return nil, err
		}

		fmt.Println("data", data)

		_test := &models.Test{}
		err = bson.Unmarshal(data, _test)
		if err != nil {
			fmt.Println("this err2 ", err)
			return nil, err
		}

		fmt.Println(_test)
		tests = append(tests, _test)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(t.ctx)

	fmt.Println("tests ", tests)

	if len(tests) == 0 {
		return nil, errors.New("test not found")
	}
	return tests, nil
}

func (t *TestServiceImpl) GetTestID(TestID *int) (*models.Test, error) {
	var testId *models.Test
	fmt.Println(*TestID)
	query := bson.D{bson.E{Key: "global.test_id", Value: *TestID}}
	err := t.testCollection.FindOne(t.ctx, query).Decode(&testId)

	return testId, err
}

func (t *TestServiceImpl) UpdateTest(Test *models.Test) error {
	filter := bson.D{primitive.E{Key: "id", Value: Test.Global.TestID}}
	multichoice := make([]bool, 0)
	for _, question := range Test.Questions {
		multichoice = append(multichoice, question.Multichoice)
	}
	topic := make([]string, 0)
	for _, question := range Test.Questions {
		topic = append(topic, question.Topic)
	}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "timeout", Value: Test.Global.Timeout}, primitive.E{Key: "multichoice", Value: multichoice}, primitive.E{Key: "topic", Value: topic}}}}
	result, _ := t.testCollection.UpdateOne(t.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (t *TestServiceImpl) CreateTest(Test *mongo.Collection) error {

	_, err := t.testCollection.InsertOne(t.ctx, Test)
	return err
}

func (t *TestServiceImpl) Clicks(test *mongo.Collection) error {
	// var clicks models.UserAnswer

	return nil
}

func (t *TestServiceImpl) StoreAnswer(Results *models.UserAnswer, History *models.History, Result *models.Result, Stats *models.Stats) error {
	var UserAnswer models.UserAnswer
	var history models.History
	// var stats models.Stats
	var result models.Result

	id := uuid.New()
	UserAnswer.ID = int(id.ID())
	history.HistoryID = int(id.ID())

	StoreAnswer := []interface{}{
		bson.D{
			{Key: "id	", Value: UserAnswer.ID},
			{Key: "timeout	", Value: UserAnswer.Timeout},
			{Key: "question	", Value: UserAnswer.Question},
			{Key: "position	", Value: UserAnswer.Multichoice},
			{Key: "result	", Value: UserAnswer.Topic},
			{Key: "answer", Value: UserAnswer.Answers},
			{Key: " clicks", Value: UserAnswer.Clicks},
			{Key: "timer", Value: UserAnswer.Histories},
			{Key: "history	", Value: UserAnswer.Histories},
			{Key: "results", Value: UserAnswer.Results},
			{Key: "completed", Value: UserAnswer.Complete},
		},
		bson.D{
			{Key: "time_start", Value: time.Now().Unix()},
			{Key: "time_end", Value: time.Now().Unix()},
		},
		bson.D{
			{Key: "id	", Value: history.HistoryID},
			{Key: "pos", Value: history.Pos},
			{Key: "timestamp", Value: time.Now().Unix()},
		},
		bson.D{
			{Key: "answer", Value: result.Answer},
			{Key: "position	", Value: result.Position},
			{Key: "result	", Value: result.Result},
		}}
	////	timer
	// timer := time.NewTimer(2 * time.Second)

	// go func() {
	// 	<-timer.C

	// 	// Printed when timer is fired
	// 	fmt.Println("timer inactivated")
	// }()

	// stopTimer := timer.Stop()

	// if stopTimer {
	// 	fmt.Println("Out of times")
	// }
	// time.Sleep(10 * time.Second)
	// 	//clicks counter
	// 	Clicks := time.NewTicker(1 * time.Second)
	// 	done := make(chan bool)

	// 	go func() {
	// 		for {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case t := <-timer.C:
	// 				fmt.Println("Ticks at", t)
	// 			}
	// 		}
	// 	}()
	// 	time.Sleep(10 * time.Second)
	// 	Clicks.Stop()
	// 	done <- true
	// 	Timer :=
	_, err := t.testCollection.InsertMany(t.ctx, StoreAnswer)
	return err
}
