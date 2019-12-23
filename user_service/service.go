package user_service

import (
	"context"
	"github.com/hoaxoan/nc_student/models"
)

type UserService struct {
	Repo Repository
	//tokenService Authable
}

func (srv *UserService) Get(ctx context.Context, req *User, res *Response) error {
	user, err := srv.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *Request, res *Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

// func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
// 	log.Println("Logging in with:", req.Email, req.Password)
// 	user, err := srv.repo.GetByEmail(req.Email)
// 	log.Println(user)
// 	if err != nil {
// 		return err
// 	}

// 	// Compares our given password against the hashed password
// 	// stored in the database
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
// 		return err
// 	}

// 	token, err := srv.tokenService.Encode(user)
// 	if err != nil {
// 		return err
// 	}
// 	res.Token = token
// 	return nil
// }

// func (srv *service) Create(ctx context.Context, req *User, res *Response) error {

// 	// Generates a hashed version of our password
// 	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	req.Password = string(hashedPass)
// 	if err := srv.repo.Create(req); err != nil {
// 		return err
// 	}
// 	res.User = req
// 	return nil
// }

// func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

// 	// Decode token
// 	claims, err := srv.tokenService.Decode(req.Token)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println(claims)

// 	if claims.User.Id == "" {
// 		return errors.New("invalid user")
// 	}

// 	res.Valid = true

// 	return nil
// }