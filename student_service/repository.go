package student_service

import (
	"context"
	md "github.com/hoaxoan/nc_student/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName  = "nc_student"
	ColName = "student"
)

type Repository interface {
	GetAll() ([]*md.Student, error)
	Get(id int) (*md.Student, error)
	Create(user *md.Student) (interface{}, error)
}

type StudentRepository struct {
	Client *mongo.Client
}

func (repo *StudentRepository) collection() *mongo.Collection {
	return repo.Client.Database(DbName).Collection(ColName)
}

func (repo *StudentRepository) GetAll() ([]*md.Student, error) {
	var students []*md.Student
	cur, err := repo.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &students)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (repo *StudentRepository) Get(id int) (*md.Student, error) {
	var student *md.Student
	student.Id = id
	filter := bson.M{"id": id}
	err := repo.collection().FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (repo *StudentRepository) Create(student *md.Student) (interface{}, error) {
	cur, err := repo.collection().InsertOne(context.TODO(), student)
	if err != nil {
		return nil, err
	}
	id := cur.InsertedID
	return id, nil
}
