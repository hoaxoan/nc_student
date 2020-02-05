package student

import (
	"context"
	"github.com/hoaxoan/nc_student/model"
)

type Usecase interface {
	Get(ctx context.Context, req *model.Student, res *model.StudentResponse) error
	GetAll(ctx context.Context, req *model.StudentRequest, res *model.StudentResponse) error
	Create(ctx context.Context, req *model.Student, res *model.StudentResponse) error
	Update(ctx context.Context, req *model.Student, res *model.StudentResponse) error
	Delete(ctx context.Context, req *model.Student, res *model.StudentResponse) error
}
