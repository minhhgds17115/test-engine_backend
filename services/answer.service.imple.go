package services

import (
	"context"
	"errors"

	"example.com/m/v2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnswerServiceImpl struct {
	answerColllection *mongo.Collection
	ctx               context.Context
}

func NewAnswerServices(answerColllection *mongo.Collection, ctx context.Context) *AnswerServiceImpl {
	return &AnswerServiceImpl{
		answerColllection: answerColllection,
		ctx:               ctx,
	}
}

func (u *AnswerServiceImpl) createAnswer(Answer *models.Answer) error {
	_, err := u.answerColllection.InsertOne(u.ctx, Answer)
	return err
}

func (u *AnswerServiceImpl) deleteAnswer(Answer *models.Answer) error {
	filter := bson.D{primitive.E{Key: "answerid", Value: Answer.AnswerId}}
	result, _ := u.answerColllection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched question found for delete")
	}
	return nil
}

func (u *AnswerServiceImpl) updateAnswer(Answer *models.Answer) error {
	filter := bson.D{primitive.E{Key: "answerid", Value: Answer.AnswerId}}
	result, _ := u.answerColllection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched answer found for delete")
	}
	return nil

}

func (u *AnswerServiceImpl) getAnswer(Answerid int) (*models.Answer, error) {
	var answerid *models.Answer
	query := bson.D{bson.E{Key: "id", Value: Answerid}}
	err := u.answerColllection.FindOne(u.ctx, query).Decode(&Answerid)
	return answerid, err
}

func (u *AnswerServiceImpl) postAnswer(Answer *models.Answer) error {
	return nil
}
