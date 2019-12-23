package user_service

import (
	"context"
	"github.com/hoaxoan/nc_student/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName  = "nc_student"
	ColName = "user"
)

type Repository interface {
	GetAll() ([]*User, error)
	Get(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) (interface{}, error)
}

type UserRepository struct {
	Client *mongo.Client
}

func NewUserRepository(client *mongo.Client) Repository {
	return &UserRepository{
		Client: client,
	}
}

func (repo *UserRepository) GetAll() ([]*User, error) {
	var users []*User
	cur, err := repo.Client.Database(DbName).Collection(ColName).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id int) (*User, error) {
	var user *User
	user.Id = id
	filter := bson.M{"id": id}
	err := repo.Client.Database(DbName).Collection(ColName).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	var user *User
	filter := bson.M{"email": email}
	err := repo.Client.Database(DbName).Collection(ColName).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *User) (interface{}, error) {
	cur, err := repo.Client.Database(DbName).Collection(ColName).InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	id := cur.InsertedID
	return id, nil
}
