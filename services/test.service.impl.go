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
func (t *TestServiceImpl) StoreUserInfo(Test *mongo.Collection) error {
	return nil
}

func (t *TestServiceImpl) StoreAnswer(Global *models.Global, ReturnedUserInformation *models.ReturnedUserInformation, Results *models.UserAnswer, History *models.History, Result *models.Result, Stats *models.Stats) error {
	var UserAnswers models.UserAnswer
	var ReturnedUserInformations models.ReturnedUserInformation
	var global models.Global
	var history models.History
	// var stats models.Stats
	var result models.Result

	id := uuid.New()
	UserAnswers.ID = int(id.ID())
	history.HistoryID = int(id.ID())

	StoreAnswer := []interface{}{
		bson.D{
			{Key: "global", Value: global.TestID},
			{Key: "name", Value: global.Name},
			{Key: "company", Value: global.Company},
			{Key: "timeout", Value: global.Timeout},
			{Key: "randomize", Value: global.Randomize},
		},
		bson.D{
			{Key: "time_start", Value: ReturnedUserInformations.TimeStart.Day()},
			{Key: "firstname", Value: ReturnedUserInformations.FirstName},
			{Key: "lastname", Value: ReturnedUserInformations.LastName},
			{Key: "Contact", Value: ReturnedUserInformations.Contact},
			{Key: "send_feedback", Value: ReturnedUserInformations.SendFeedback},
			{Key: "Feedback", Value: ReturnedUserInformations.Feedback},
		},
		bson.D{
			{Key: "time_start", Value: time.Now().Unix()},
			{Key: "time_end", Value: time.Now().Add(30 * time.Second).Unix()},
		},
		bson.D{
			{Key: "id	", Value: UserAnswers.ID},
			{Key: "timeout	", Value: UserAnswers.Timeout},
			{Key: "question	", Value: UserAnswers.Question},
			{Key: "position	", Value: UserAnswers.Multichoice},
			{Key: "result	", Value: UserAnswers.Topic},
			{Key: "answer", Value: UserAnswers.Answers},
			{Key: " clicks", Value: UserAnswers.Clicks},
			{Key: "timer", Value: UserAnswers.Histories},
			{Key: "history	", Value: UserAnswers.Histories},
			{Key: "results", Value: UserAnswers.Results},
			{Key: "completed", Value: UserAnswers.Complete},
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
		},
	}
	////((	timer
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
	// 	Timer :=))

	_, err := t.testCollection.InsertMany(t.ctx, StoreAnswer)
	return err
}
