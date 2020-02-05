package student

import (
	"github.com/hoaxoan/nc_student/model"
)

type Repository interface {
	GetAll() ([]*model.Student, error)
	Get(id int) (*model.Student, error)
	Create(student *model.Student) error
	Update(student *model.Student) error
	Delete(student *model.Student) error
}
