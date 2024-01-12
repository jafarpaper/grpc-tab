package repository

import (
	protoData "grpc-server/proto"
	"grpc-server/user"

	"github.com/jinzhu/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func CreateUserRepoImpl(db *gorm.DB) user.UserRepo {
	return &UserRepoImpl{db}
}

func (e *UserRepoImpl) GetUsers() ([]*protoData.User, error) {
	var users []*protoData.User
	err := e.db.Table("user").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (e *UserRepoImpl) FindUserById(id protoData.UserId) (*protoData.User, error) {
	var user protoData.User
	err := e.db.Table("user").Where("id = ?", id.Id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (e *UserRepoImpl) InsertUser(user *protoData.User) error {
	err := e.db.Table("user").Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *UserRepoImpl) UpdateUser(user *protoData.UserUpdate) error {
	var us protoData.User
	err := e.db.Table("user").Where("id = ?", user.Id).First(&us).Update(&user.User).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *UserRepoImpl) DeleteUser(id *protoData.UserId) error {
	var user protoData.User
	err := e.db.Table("user").Where("id = ?", id.Id).First(&user).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
