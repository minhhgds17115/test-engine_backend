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
	for _, question := range Test.Question {
		multichoice = append(multichoice, question.Multichoice)
	}
	topic := make([]string, 0)
	for _, question := range Test.Question {
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
func (t *TestServiceImpl) StoreUserInfo(userInformation *models.UserInformation) error {
	// var ReturnedUserInformations models.ReturnedUserInformation
	// var global models.Global

	fmt.Println("this is global ", userInformation)

	id := uuid.New()
	userInformation.Global.TestID = int(id.ID())
	userInformation.Candidate.ID = int(id.ID())

	userInformation.Candidate.TimeStart = time.Now().Unix()
	// ReturnInfo := []interface{}{
	// 	bson.D{
	// 		{Key: "global", Value: userInformation.TestID},
	// 		{Key: "name", Value: global.Name},
	// 		{Key: "company", Value: global.Company},
	// 		{Key: "timeout", Value: global.Timeout},
	// 		{Key: "randomize", Value: global.Randomize},
	// 	},
	// 	bson.D{
	// 		{Key: "time_start", Value: ReturnedUserInformations.TimeStart.Day()},
	// 		{Key: "firstname", Value: ReturnedUserInformations.FirstName},
	// 		{Key: "lastname", Value: ReturnedUserInformations.LastName},
	// 		{Key: "contact", Value: ReturnedUserInformations.Contact},
	// 		{Key: "send_feedback", Value: ReturnedUserInformations.SendFeedback},
	// 		{Key: "feedback", Value: ReturnedUserInformations.Feedback},
	// 	},
	// }

	_, err := t.testCollection.InsertOne(t.ctx, userInformation)
	if err != nil {
		return err
	}
	// if ReturnInfo == nil && err == nil {
	// 	return errors.New("no return info")
	// }
	// fmt.Println("Information return ")
	return nil
}

func (t *TestServiceImpl) ReturnAnswer(returnAnswer *models.ReturnedAnswer) error {

	// uuid validation
	id := uuid.New()
	returnAnswer.Global.TestID = int(id.ID())
	returnAnswer.Questions.ID = int(id.ID())
	returnAnswer.Questions.Question.ID = int(id.ID())
	returnAnswer.Questions.Answer.AnswerId = int(id.ID())
	returnAnswer.Questions.Question.Answers.AnswerId = int(id.ID())
	returnAnswer.Questions.Histories.HistoryID = int(id.ID())

	// time
	returnAnswer.Stats.TimeStart = time.Now().Unix()
	returnAnswer.Stats.TimeEnd = time.Now().Unix()
	returnAnswer.Questions.Histories.Timestamp = time.Now().Unix()
	returnAnswer.ReturnedUserInformation.TimeStart = time.Now().Unix()

	_, err := t.testCollection.InsertOne(t.ctx, returnAnswer)
	if err != nil {
		return err
	}
	return err
}
