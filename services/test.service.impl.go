package services

import (
	"context"
	"errors"
	"fmt"

	"example.com/m/v2/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// test serviceImplementation
type TestServiceImpl struct {
	testCollection *mongo.Collection
	ctx            context.Context
}

// new Test Service
func NewTestService(testCollection *mongo.Collection, ctx context.Context) *TestServiceImpl {
	return &TestServiceImpl{
		testCollection: testCollection,
		ctx:            ctx,
	}
}

// Get all test services
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

// Get test by id
func (t *TestServiceImpl) GetTestID(TestID *int) (*models.ReturnedAnswer, error) {
	var test *models.ReturnedAnswer
	fmt.Println(*TestID)
	query := bson.D{bson.E{Key: "test_id"}}
	err := t.testCollection.FindOne(t.ctx, query).Decode(&test)

	return test, err
}

// Update test
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

// Create a new Test
func (t *TestServiceImpl) CreateTest(Test *mongo.Collection) error {

	_, err := t.testCollection.InsertOne(t.ctx, Test)
	return err
}

// Store Candidate
func (t *TestServiceImpl) StoreCandidateInfo(candidateInformation *models.CandidateInformation) error {

	// var global models.Global
	// newCandidateInformation := NewCandidateInterface{}
	// vErrs, err := validation.Run(NewCandidateInterface)
	fmt.Println("this is global ", candidateInformation)

	// 	return err
	// }
	// if len(vErrs) > 0 {
	// }

	id := uuid.New()
	candidateInformation.Global.TestID = int(id.ID())

	_, err := t.testCollection.InsertOne(t.ctx, candidateInformation)
	if err != nil {
		return err
	}

	return nil
}

// Returned Answer
func (t *TestServiceImpl) ReturnAnswer(returnAnswer *models.ReturnedAnswer) error {

	// uuid validation
	id := uuid.New()
	returnAnswer.Global.TestID = int(id.ID())

	// returnAnswer.Questions.Question = []models.Questions.Question.

	_, err := t.testCollection.InsertOne(t.ctx, returnAnswer)
	if err != nil {
		return err
	}
	return err
}
