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

func (u *AnswerServiceImpl) CreateAnswer(Answer *models.Answer) error {
	_, err := u.answerColllection.InsertOne(u.ctx, Answer)
	return err
}

func (u *AnswerServiceImpl) DeleteAnswer(Answer *models.Answer) error {
	filter := bson.D{primitive.E{Key: "Answer", Value: Answer.Answer}}
	result, _ := u.answerColllection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched question found for delete")
	}
	return nil
}

func (u *AnswerServiceImpl) UpdateAnswer(Answer *models.Answer) error {
	filter := bson.D{primitive.E{Key: "answer", Value: Answer.Answer}}
	result, _ := u.answerColllection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched answer found for delete")
	}
	return nil

}

func (u *AnswerServiceImpl) GetAnswer(Answerid int) (*models.Answer, error) {
	var answerid *models.Answer
	query := bson.D{bson.E{Key: "id", Value: Answerid}}
	err := u.answerColllection.FindOne(u.ctx, query).Decode(&Answerid)
	return answerid, err
}

func (u *AnswerServiceImpl) PostAnswer(Answer *models.Answer) error {
	fmt.Println("user answerd recoreded ", u.answerColllection.Name(), u.answerColllection.Database().Name())
	_, err := u.answerColllection.InsertOne(u.ctx, *Answer)
	return err
}
