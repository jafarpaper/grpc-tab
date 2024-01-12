package usecase

import (
	protoData "grpc-server/proto"
	"grpc-server/user"
)

type UserUseCaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUseCase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUseCaseImpl{userRepo}
}

func (e *UserUseCaseImpl) GetUsers() ([]*protoData.User, error) {
	return e.userRepo.GetUsers()
}

func (e *UserUseCaseImpl) FindUserById(id protoData.UserId) (*protoData.User, error) {
	return e.userRepo.FindUserById(id)
}

func (e *UserUseCaseImpl) InsertUser(user *protoData.User) error {
	return e.userRepo.InsertUser(user)
}

func (e *UserUseCaseImpl) UpdateUser(user *protoData.UserUpdate) error {
	return e.userRepo.UpdateUser(user)
}

func (e *UserUseCaseImpl) DeleteUser(id *protoData.UserId) error {
	return e.userRepo.DeleteUser(id)
}
