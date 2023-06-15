package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/RogerWaldron/go-reserveration-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "users"

type UserStore interface {
	Dropper
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, Map, types.UpdateUserParams) error
} 

type MongoUserStore struct {
	client *mongo.Client
	collect *mongo.Collection
}

func NewMongoUserStore(c *mongo.Client, dbName string) *MongoUserStore {
	return &MongoUserStore{
		client: c,
		collect: c.Database(dbName).Collection(collectionName),
	}
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("--- (´･_･`) --- dropping collection: " + collectionName)
	return s.collect.Drop(ctx)
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User

	cur, err := s.collect.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &users)
	if err != nil { 
		return []*types.User{}, err
	}

	return users, nil
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = s.collect.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
 
func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	result, err := s.collect.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error {
	objID, err := primitive.ObjectIDFromHex(filter["_id"].(string))
	if err != nil {
		return err
	}

	filter["_id"] = objID
	update := bson.M{"$set": params.ToBSON()}
	_, err = s.collect.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	
	filter := bson.M{"_id": objID}
	result, deleteErr := s.collect.DeleteOne(ctx, filter)
	if deleteErr != nil {
		return deleteErr
	}

	if result.DeletedCount == 0 {
		return errors.New(id + " was not found")
	}

	return nil
}