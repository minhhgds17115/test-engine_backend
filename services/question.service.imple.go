package services

import (
	"context"
	"errors"
	"strconv"

	"example.com/m/v2/models"
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

func (u *QuestionsServiceImpl) CreateQuestions(Questions *models.Questions) error {
	_, err := u.QuestionsCollection.InsertOne(u.ctx, Questions)
	return err
}

func (u *QuestionsServiceImpl) UpdateQuestions(Questions *models.Questions) error {
	filter := bson.D{primitive.E{Key: "id", Value: Questions.ID}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "ID", Value: Questions.ID}, primitive.E{Key: "Information", Value: Questions.Information}, primitive.E{Key: "Topic", Value: Questions.Topic}, primitive.E{Key: "Timeouts", Value: Questions.Timeout}}}}
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
		return errors.New("no matched Questions found for delete")
	}
	return nil
}

func (u *QuestionsServiceImpl) GetQuestions(QuestionsId string) (*models.Questions, error) {
	var QuestionsIDs *models.Questions
	query := bson.D{bson.E{Key: "id", Value: QuestionsIDs}}
	err := u.QuestionsCollection.FindOne(u.ctx, query).Decode(&QuestionsIDs)
	return QuestionsIDs, err
}
