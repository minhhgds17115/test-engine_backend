package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"strconv"

	"example.com/m/v2/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(Candidate *models.Candidate) error {
	fmt.Println("Candidate collection created", u.usercollection.Name(), u.usercollection.Database().Name())
	id := uuid.New()
	Candidate.ID = int(id.ID())

	Candidate.TimeStart = time.Now().Unix()
	fmt.Println(*Candidate)
	_, err := u.usercollection.InsertOne(u.ctx, *Candidate)
	return err
}

func (u *UserServiceImpl) GetUserEmail(Contact *string) (*models.Candidate, error) {
	var Candidate *models.Candidate
	query := bson.D{bson.E{Key: "contact", Value: Contact}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&Candidate)
	return Candidate, err
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.Candidate, error) {
	var Candidate []*models.Candidate
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var candidate models.Candidate
		err := cursor.Decode(&Candidate)
		if err != nil {
			return nil, err
		}
		Candidate = append(Candidate, &candidate)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(Candidate) == 0 {
		return nil, errors.New("candidate not found")
	}
	return Candidate, nil
}

func (u *UserServiceImpl) UpdateUser(Candidate *models.Candidate) error {
	filter := bson.D{primitive.E{Key: "first_name", Value: Candidate.FirstName}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "first_name", Value: Candidate.FirstName}, primitive.E{Key: "last_name", Value: Candidate.LastName}, primitive.E{Key: "Contact", Value: Candidate.Contact}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error {
	idNumber, _ := strconv.Atoi(*id)
	filter := bson.D{primitive.E{Key: "id", Value: idNumber}}

	fmt.Println(*id)
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched Candidate found for delete")
	}
	return nil
}
