package services

import (
	"context"
	"errors"
	"fmt"

	"example.com/m/v2/models"
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
		return nil, errors.New("Test not found")
	}
	return tests, nil
}

func (t *TestServiceImpl) GetTestID(TestID *int) (*models.Test, error) {
	var testId *models.Test
	query := bson.D{bson.E{Key: "testId", Value: testId}}
	err := t.testCollection.FindOne(t.ctx, query).Decode(&testId)
	return testId, err

}

func (t *TestServiceImpl) UpdateTest(Test *models.Test) error {
	filter := bson.D{primitive.E{Key: "id", Value: Test.Global.TestID}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "timeout", Value: Test.Global.Timeout}, primitive.E{Key: "multichoice", Value: Test.Questions[0].Multichoice}, primitive.E{Key: "topic", Value: Test.Questions[0].Topic}}}}
	result, _ := t.testCollection.UpdateOne(t.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (t *TestServiceImpl) CreateTest(test *mongo.Collection) error {
	_, err := t.testCollection.InsertOne(t.ctx, test)
	return err
}
