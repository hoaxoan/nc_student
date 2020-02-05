package repository

import (
	"context"

	"github.com/hoaxoan/nc_student/model"
	"github.com/hoaxoan/nc_student/student"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName  = "nc_student"
	ColName = "student"
)

type studentRepository struct {
	Client *mongo.Client
}

func NewStudentRepository(Client *mongo.Client) student.Repository {
	return &studentRepository{Client}
}

func (repo *studentRepository) collection() *mongo.Collection {
	return repo.Client.Database(DbName).Collection(ColName)
}

func (repo *studentRepository) GetAll() ([]*model.Student, error) {
	var students []*model.Student
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

func (repo *studentRepository) Get(id int) (*model.Student, error) {
	var student *model.Student
	student.Id = id
	filter := bson.M{"id": id}
	err := repo.collection().FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (repo *studentRepository) Create(student *model.Student) error {
	_, err := repo.collection().InsertOne(context.TODO(), student)
	if err != nil {
		return err
	}
	return nil
}

func (repo *studentRepository) Update(student *model.Student) error {
	filter := bson.M{"id": student.Id}
	_, err := repo.collection().UpdateOne(context.TODO(), filter, bson.M{"$set": student})
	if err != nil {
		return err
	}
	return nil
}

func (repo *studentRepository) Delete(student *model.Student) error {
	filter := bson.M{"id": student.Id}
	_, err := repo.collection().UpdateOne(context.TODO(), filter, bson.M{"$set": student})
	if err != nil {
		return err
	}
	return nil
}
