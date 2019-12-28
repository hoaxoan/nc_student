package model

type Student struct {
	Id        int    `json:"id,omitempty" bson:"id"`
	FirstName string `json:"first_name,omitempty" bson:"first_name`
	LastName  string `json:"last_name,omitempty" bson:"las_name"`
	Email     string `json:"email,omitempty" bson:"email"`
	Password  string `json:"password,omitempty" bson:"password`
}

type StudentResponse struct {
	Student  *Student   `json"student,omitempty"`
	Students []*Student `json:"students,omitempty"`
	Errors   []*Error   `json:"errors,omitempty"`
}
