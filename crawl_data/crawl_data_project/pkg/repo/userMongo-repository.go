package repo

import (
	"context"
	"crawl_data/pkg/model"
	"crawl_data/pkg/utils"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	InsertUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
	FindByEmail(email string) (*model.User, error)
	ProfileUser(userID string) (*model.User, error)
}

type userConnectionMon struct {
	postgresql UserRepositoryPostgres
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepositoryMon(postgresql UserRepositoryPostgres) UserRepository {
	return &userConnectionMon{
		postgresql: postgresql,
	}
}

var mongoDBNam = "golang_mongodb_api"

func (db *userConnectionMon) InsertUser(user *model.User) (*model.User, error) {

	user.Password = hashAndSalt([]byte(user.Password))
	//user.CreatedAt = time.Now().Format(time.RFC1123Z)
	user.CreatedAt = time.Now()
	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database(mongoDBNam).Collection("user")

	user.ID = uuid.Must(uuid.NewRandom())
	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Printf("Could not create User: %v", err)
		return nil, err
	}
	var userRes *model.User

	err = collection.FindOne(
		context.TODO(),
		bson.D{primitive.E{Key: "_id", Value: res.InsertedID}},
	).Decode(&userRes)
	if err != nil {
		log.Printf("Could not create User: %v", err)
		return nil, err
	}
	//userJson, err := json.Marshal(userRes)
	//if err != nil {
	//	log.Printf("Could not convert to json user: %v", err)
	//	return user
	//}

	_, err = db.postgresql.InsertUser(user)
	if err != nil {
		logrus.Error("Error insertUser Postgresql: ", err)
	}

	return userRes, nil
}

func (db *userConnectionMon) UpdateUser(user *model.User) (*model.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))

	user.UpdatedAt = time.Now()

	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	collection := client.Database(mongoDBNam).Collection("user")

	var userRes *model.User

	err := collection.FindOne(
		context.TODO(),
		bson.D{primitive.E{Key: "_id", Value: user.ID}},
	).Decode(&userRes)
	if err != nil {
		log.Printf("Could not create User: %v", err)
		return nil, err
	}
	user.CreatedAt = userRes.CreatedAt

	filter := bson.D{primitive.E{Key: "_id", Value: user.ID}}
	update := bson.M{
		"$set": user,
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	var updatedUser model.User
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedUser)
	if err != nil {
		log.Printf("Could not save User: %v", err)
		return nil, err
	}

	_, err = db.postgresql.UpdateUser(user)
	if err != nil {
		logrus.Error("Error updateUser Postgresql: ", err)
	}

	return user, nil
}

func (db *userConnectionMon) VerifyCredential(email string, password string) interface{} {
	var user model.User

	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database(mongoDBNam).Collection("user")

	error := collection.FindOne(
		context.TODO(),
		bson.D{primitive.E{Key: "email", Value: email}},
	).Decode(&user)
	if error == nil {
		return user
	}
	return nil
}

func (db *userConnectionMon) IsDuplicateEmail(email string) bool {
	var user model.User

	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database(mongoDBNam).Collection("user")

	error := collection.FindOne(
		context.TODO(),
		bson.D{primitive.E{Key: "email", Value: email}},
	).Decode(&user)
	if error == nil {
		return false
	}
	return true
}

func (db *userConnectionMon) FindByEmail(email string) (*model.User, error) {
	var user model.User

	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database(mongoDBNam).Collection("user")

	error := collection.FindOne(
		context.TODO(),
		bson.D{primitive.E{Key: "email", Value: email}},
	).Decode(&user)
	if error != nil {
		logrus.Error("Could not save User: %v", error)
		return nil, error
	}
	return &user, nil
}

func (db *userConnectionMon) ProfileUser(userID string) (*model.User, error) {
	id := uuid.MustParse(userID)

	var user model.User

	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database(mongoDBNam).Collection("user")

	error := collection.FindOne(
		context.TODO(),
		bson.D{primitive.E{Key: "_id", Value: id}},
	).Decode(&user)
	if error != nil {
		logrus.Error("Could not find User: %v", error)
	}
	return &user, nil
}

//func hashAndSalt1(pwd []byte) string {
//	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
//	if err != nil {
//		log.Println(err)
//		panic("Failed to hash a password")
//	}
//	return string(hash)
//}
