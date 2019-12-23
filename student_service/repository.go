package student_service

import (
	"context"
	"github.com/hoaxoan/nc_student/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName  = "nc_student"
	ColName = "student"
)

type Repository interface {
	GetAll() ([]*Student, error)
	Get(id int) (*Student, error)
	Create(user *Student) (interface{}, error)
}

type StudentRepository struct {
	Client *mongo.Client
}

func NewStudentRepository(client *mongo.Client) Repository {
	return &StudentRepository{
		Client: client,
	}
}

func (repo *StudentRepository) GetAll() ([]*Student, error) {
	var students []*Student
	cur, err := repo.Client.Database(DbName).Collection(ColName).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &students)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (repo *StudentRepository) Get(id int) (*Student, error) {
	var student *Student
	student.Id = id
	filter := bson.M{"id": id}
	err := repo.Client.Database(DbName).Collection(ColName).FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (repo *StudentRepository) Create(student *Student) (interface{}, error) {
	cur, err := repo.Client.Database(DbName).Collection(ColName).InsertOne(context.TODO(), student)
	if err != nil {
		return nil, err
	}
	id := cur.InsertedID
	return id, nil
}
