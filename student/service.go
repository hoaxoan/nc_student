package student

import (
	"context"
	md "github.com/hoaxoan/nc_course/nc_student/model"
)

type StudentService struct {
	Repo *StudentRepository
}

func (srv *StudentService) Get(ctx context.Context, req *md.Student, res *md.StudentResponse) error {
	student, err := srv.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.Student = student
	return nil
}

// func (srv *StudentService) GetAll(ctx context.Context, req *Request, res *Response) error {
// 	students, err := srv.Repo.GetAll()
// 	if err != nil {
// 		return err
// 	}
// 	res.Students = students
// 	return nil
// }