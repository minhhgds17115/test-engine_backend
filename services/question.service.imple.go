package services

import (
	"context"
	"errors"
	"strconv"

	"example.com/m/v2/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionsServiceImpl struct {
	QuestionsCollection *mongo.Collection
	ctx                 context.Context
}

func NewQuestionsServices(QuestionsCollection *mongo.Collection, ctx context.Context) *QuestionsServiceImpl {
	return &QuestionsServiceImpl{
		QuestionsCollection: QuestionsCollection,
		ctx:                 ctx,
	}
}

func (u *QuestionsServiceImpl) CreateQuestions(Question *models.Question) error {
	id := uuid.New()
	Question.ID = int(id.ID())
	Question.Answers.AnswerId = int(id.ID())

	_, err := u.QuestionsCollection.InsertOne(u.ctx, Question)
	return err
}

func (u *QuestionsServiceImpl) UpdateQuestions(Question *models.Question) error {
	filter := bson.D{primitive.E{Key: "id", Value: Question.ID}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "ID", Value: Question.ID}, primitive.E{Key: "Information", Value: Question.Information}, primitive.E{Key: "Topic", Value: Question.Topic}, primitive.E{Key: "Timeouts", Value: Question.Timeout}}}}
	result, _ := u.QuestionsCollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *QuestionsServiceImpl) DeleteQuestions(QuestionsId string) error {
	id, _ := strconv.Atoi(QuestionsId)
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	result, _ := u.QuestionsCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount == 0 {
		return errors.New("no matched Question found for delete")
	}
	return nil
}

func (u *QuestionsServiceImpl) GetQuestionsByID(QuestionsId string) (*models.Question, error) {
	var QuestionsIDs *models.Question
	query := bson.D{bson.E{Key: "id", Value: QuestionsIDs}}
	err := u.QuestionsCollection.FindOne(u.ctx, query).Decode(&QuestionsIDs)
	return QuestionsIDs, err
}

// func (u *QuestionsServiceImpl) GetAllAnswerQuestions(Question *models.Question) {

// }
