package usecase

import (
	"context"

	"github.com/hoaxoan/nc_student/model"
	"github.com/hoaxoan/nc_student/student"
)

type studentUsecase struct {
	Repo student.Repository
}

func NewStudentUsecase(repo student.Repository) student.Usecase {
	return &studentUsecase{
		Repo: repo,
	}
}

func (scase *studentUsecase) Get(ctx context.Context, req *model.Student, res *model.StudentResponse) error {
	student, err := scase.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.Student = student
	return nil
}

func (scase *studentUsecase) GetAll(ctx context.Context, req *model.StudentRequest, res *model.StudentResponse) error {
	students, err := scase.Repo.GetAll(req)
	if err != nil {
		return err
	}
	res.Students = students
	return nil
}

func (scase *studentUsecase) Create(ctx context.Context, req *model.Student, res *model.StudentResponse) error {
	if err := scase.Repo.Create(req); err != nil {
		return err
	}
	res.Student = req
	return nil
}

func (scase *studentUsecase) Update(ctx context.Context, req *model.Student, res *model.StudentResponse) error {
	if err := scase.Repo.Update(req); err != nil {
		return err
	}
	res.Student = req
	return nil
}

func (scase *studentUsecase) Delete(ctx context.Context, req *model.Student, res *model.StudentResponse) error {
	if err := scase.Repo.Update(req); err != nil {
		return err
	}
	res.Student = req
	return nil
}
