package service

import (
	"crawl_data/pkg/model"
	"crawl_data/pkg/repo"
	"github.com/sirupsen/logrus"
	"log"

	"github.com/mashingan/smapping"

	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user *model.RegisterDTO) (*model.User,error)
	FindByEmail(email string) (*model.User,error)
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repo.UserRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repo.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(model.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user *model.RegisterDTO) (*model.User,error) {
	userToCreate := &model.User{}
	err := smapping.FillStruct(userToCreate, smapping.MapFields(user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res, err := service.userRepository.InsertUser(userToCreate)
	if err != nil{
		logrus.Error("Error create User: ", err)
		return nil, err
	}
	return res, nil
}

func (service *authService) FindByEmail(email string) (*model.User,error) {
	user, err := service.userRepository.FindByEmail(email)
	if err != nil{
		logrus.Error("Error service findbyEmail: ", err)
		return nil, err
	}
	return user, nil
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return res

}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
