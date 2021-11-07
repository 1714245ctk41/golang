package repo

import (
	"crawl_data/pkg/model"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryPostgres interface {
	InsertUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
	FindByEmail(email string) (*model.User, error)
	ProfileUser(userID string) (*model.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepositoryPostgres(db *gorm.DB) UserRepositoryPostgres {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user *model.User) (*model.User, error) {
	if err := db.connection.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (db *userConnection) UpdateUser(user *model.User) (*model.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser *model.User
		if err := db.connection.Find(&tempUser, user.ID).Error; err != nil {
			return nil, err
		}
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user, nil
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user model.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) bool {
	var user model.User
	db.connection.Where("email = ?", email).Take(&user)
	if user.ID != uuid.Nil {
		return false
	}
	return true
}

func (db *userConnection) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := db.connection.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *userConnection) ProfileUser(userID string) (*model.User, error) {
	var user model.User
	if err := db.connection.Preload("Books").Preload("Books.User").Find(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
