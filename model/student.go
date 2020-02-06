package model

type Student struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	ClassName string `json:"class_name" bson:"class_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type StudentRequest struct {
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	ClassName string `json:"class_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty"`
}

type StudentResponse struct {
	Student  *Student   `json"student,omitempty"`
	Students []*Student `json:"students,omitempty"`
	Errors   []*Error   `json:"errors,omitempty"`
}
