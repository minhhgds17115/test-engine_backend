package services

import (
	"context"
	"errors"

	"example.com/m/v2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionServiceImpl struct {
	questionCollection *mongo.Collection
	ctx                context.Context
}

func NewQuestionServices(questionCollection *mongo.Collection, ctx context.Context) *QuestionServiceImpl {
	return &QuestionServiceImpl{
		questionCollection: questionCollection,
		ctx:                ctx,
	}
}

func (u *QuestionServiceImpl) CreateQuestion(Question *models.Question) error {
	_, err := u.questionCollection.InsertOne(u.ctx, Question)
	return err
}

func (u *QuestionServiceImpl) UpdateQuestion(Question *models.Question) error {
	filter := bson.D{primitive.E{Key: "questionid", Value: Question.QuestionID}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "ID", Value: Question.QuestionID}, primitive.E{Key: "Information", Value: Question.Information}, primitive.E{Key: "Topic", Value: Question.Topic}, primitive.E{Key: "Timeouts", Value: Question.Timeouts}}}}
	result, _ := u.questionCollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *QuestionServiceImpl) DeleteQuestion(QuestionId string) error {
	filter := bson.D{primitive.E{Key: "QuestionId", Value: QuestionId}}
	result, _ := u.questionCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched question found for delete")
	}
	return nil
}

func (u *QuestionServiceImpl) GetQuestion(QuestionId string) (*models.Question, error) {
	var questionIDs *models.Question
	query := bson.D{bson.E{Key: "id", Value: questionIDs}}
	err := u.questionCollection.FindOne(u.ctx, query).Decode(&questionIDs)
	return questionIDs, err
}
