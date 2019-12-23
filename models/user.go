package models

type User struct {
	Id        int    `json:"id,omitempty" bson:"id"`
	FirstName string `json:"first_name,omitempty" bson:"first_name`
	LastName  string `json:"last_name,omitempty" bson:"las_name"`
	Email     string `json:"email,omitempty" bson:"email"`
	Password  string `json:"password,omitempty" bson:"password`
}

type Request struct {
}

type Response struct {
	User     *User      `json:"user,omitempty"`
	Users    []*User    `json:"users,omitempty"`
	Student  *Student   `json:"student,omitempty"`
	Students []*Student `json:"students,omitempty"`
	Errors   []*Error   `json:"errors,omitempty"`
}

type Token struct {
	Token  string   `json:"token,omitempty"`
	Valid  bool     `json:"valid,omitempty"`
	Errors []*Error `json:"errors,omitempty"`
}

type Error struct {
	Code        int32  `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}
