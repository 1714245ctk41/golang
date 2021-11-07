package service

import (
	"crawl_data/pkg/model"
	"crawl_data/pkg/repo"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

//UserService is a contract.....
type UserService interface {
	Update(user *model.UserUpdateDTO) (*model.User, error)
	Profile(userID string) (*model.User, error)
}

type userService struct {
	userRepository repo.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo repo.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user *model.UserUpdateDTO) (*model.User, error) {
	// customIDuser := strings.ReplaceAll(user.ID, "-", "")
	// byteID := []byte(customIDuser)
	idUser := uuid.MustParse(user.ID)
	// if err != nil {
	// 	logrus.Error(err)
	// 	return nil, err
	// }
	userToUpdate := &model.User{
		ID:        idUser,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Token:     "",
		Role:      user.Role,
		BaseModel: model.BaseModel{},
	}

	updatedUser, err := service.userRepository.UpdateUser(userToUpdate)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return updatedUser, nil
}

func (service *userService) Profile(userID string) (*model.User, error) {
	userProfile, err := service.userRepository.ProfileUser(userID)
	if err != nil {
		logrus.Error("error to get profile user: ", err)
		return nil, err
	}
	return userProfile, nil
}
