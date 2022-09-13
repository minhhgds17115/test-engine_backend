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

func (t *TestServiceImpl) StoreAnswer(Results *models.UserAnswer, History *models.UserAnswer) error {
	var result models.UserAnswer

	id := uuid.New()
	result.Results.ID = int(id.ID())
	StoreAnswer := []interface{}{
		bson.D{
			{Key: "id	", Value: result.Results.ID},
			{Key: "answer	", Value: result.Results.Answer},
			{Key: "position	", Value: result.Results.Position},
			{Key: "result	", Value: result.Results.Result},
			{Key: "question	", Value: result.Results.Questions},
			{Key: " clicks", Value: result.Results.Clicks},
		},
		bson.D{
			{Key: "history_id	", Value: result.Results.ID},
			{Key: "pos", Value: result.Results.Position},
			{Key: "timestamp", Value: time.Now().Unix()},
		},
	}

	_, err := t.testCollection.InsertMany(t.ctx, StoreAnswer)
	return err
}
