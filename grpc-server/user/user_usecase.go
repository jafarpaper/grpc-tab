package user

import protoData "grpc-server/proto"

type UserUsecase interface {
	GetUsers() ([]*protoData.User, error)
	FindUserById(id protoData.UserId) (*protoData.User, error)
	InsertUser(user *protoData.User) error
	UpdateUser(user *protoData.UserUpdate) error
	DeleteUser(id *protoData.UserId) error
}
