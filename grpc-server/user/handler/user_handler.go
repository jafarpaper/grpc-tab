package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	protoData "grpc-server/proto"
	"grpc-server/user"
)

type UserHandler struct {
	userUseCase user.UserUsecase
	protoData.UnimplementedUsersServer
}

func NewUserHandler(gr *grpc.Server, userUseCase user.UserUsecase) {
	userController := &UserHandler{userUseCase, protoData.UnimplementedUsersServer{}}

	protoData.RegisterUsersServer(gr, userController)
}

func (e *UserHandler) GetUserList(ctx context.Context, in *empty.Empty) (*protoData.UserList, error) {
	users, err := e.userUseCase.GetUsers()
	if err != nil {
		return nil, err
	}
	var userData = make([]*protoData.User, 0)
	for _, value := range users {
		var data = protoData.User{
			Id:      value.Id,
			Name:    value.Name,
			Email:   value.Email,
			Address: value.Address,
		}
		userData = append(userData, &data)
	}

	var userList = protoData.UserList{
		List: userData,
	}

	return &userList, nil
}

func (e *UserHandler) GetUserById(ctx context.Context, in *protoData.UserId) (*protoData.User, error) {
	user, err := e.userUseCase.FindUserById(*in)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *UserHandler) InsertUser(ctx context.Context, in *protoData.User) (*empty.Empty, error) {
	err := e.userUseCase.InsertUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}

func (e *UserHandler) UpdateUser(ctx context.Context, in *protoData.UserUpdate) (*empty.Empty, error) {
	err := e.userUseCase.UpdateUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}

func (e *UserHandler) DeleteUser(ctx context.Context, in *protoData.UserId) (*empty.Empty, error) {
	err := e.userUseCase.DeleteUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}
