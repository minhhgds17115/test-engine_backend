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

func (u *UserServiceImpl) CreateUser(user *models.Users) error {
	fmt.Println("user collection created", u.usercollection.Name(), u.usercollection.Database().Name())

	_, err := u.usercollection.InsertOne(u.ctx, *user)
	return err
}

func (u *UserServiceImpl) GetUserEmail(email *string) (*models.Users, error) {
	var user *models.Users
	query := bson.D{bson.E{Key: "email", Value: email}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.Users, error) {
	var users []*models.Users
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.Users
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("user not found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.Users) error {
	filter := bson.D{primitive.E{Key: "firstname", Value: user.FirstName}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: user.FirstName}, primitive.E{Key: "last_name", Value: user.LastName}, primitive.E{Key: "Email", Value: user.Email}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(firstname *string) error {
	filter := bson.D{primitive.E{Key: "FirstName", Value: firstname}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched user found for delete")
	}
	return nil
}
