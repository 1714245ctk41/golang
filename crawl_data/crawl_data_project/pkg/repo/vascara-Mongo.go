package repo

import (
	"context"
	"crawl_data/pkg/model"
	"crawl_data/pkg/utils"
	"log"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VascaraMongoRepository interface {
	InsertProduct(product *model.ProductDetail, client *mongo.Client) error
	InsertCategory(category []model.CategoryFa) error
	FindProductById(id int) (*model.ProductDetail, error)
	GetAllCategory() ([]model.CategoryFa, error)
}

type vascaraMongoRepository struct {
	postgres VascaraPostgresRepository
}

func NewVascaraMongoRepository(postgres VascaraPostgresRepository) VascaraMongoRepository {
	return &vascaraMongoRepository{
		postgres: postgres,
	}
}

var collectionVascaraProduct = "vascara"
var collectionVascaraCategory = "vascara_category"

func (vas *vascaraMongoRepository) InsertProduct(product *model.ProductDetail, client *mongo.Client) error {
	//user.CreatedAt = time.Now().Format(time.RFC1123Z)

	collection := client.Database(mongoDBNam).Collection(collectionVascaraProduct)
	collectionCate := client.Database(mongoDBNam).Collection(collectionVascaraCategory)

	var category []model.CategoryFa
	x, err := collectionCate.Find(context.TODO(), bson.M{})
	err = x.All(context.TODO(), &category)
	if err != nil {
		logrus.Error("Error get category: %v", err)
	}
	for _, v := range category {
		for _, vc := range v.CategoryChilds {
			if strings.Contains(vc.Name, product.CategoryName) {
				product.CategoryId = vc.ID
				break
			}
		}
	}

	product.BaseModel.CreatedAt = time.Now()
	product.BaseModel.UpdatedAt = time.Now()

	filter := bson.D{primitive.E{Key: "_id", Value: product.ID}}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	update := bson.D{primitive.E{Key: "$set", Value: product}}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Err()
	if err != nil {
		logrus.Error("Could not find to update product just inert: %v", err)
	}
	if err = vas.postgres.InsertProduct(product); err != nil {
		logrus.Error("Could not insert product: %v", err)

	}
	if err != nil {
		return err
	}
	return nil
}

func (vas *vascaraMongoRepository) InsertCategory(category []model.CategoryFa) error {
	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	collection := client.Database(mongoDBNam).Collection(collectionVascaraCategory)

	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	_, err := collection.DeleteMany(context.TODO(), bson.D{{}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range category {
		if err := vas.postgres.InsertCategory(&v); err != nil {
			return err
		}
		v.CreatedAt = time.Now()
		_, err := collection.InsertOne(context.TODO(), v)
		if err != nil {
			log.Printf("Could not insert product at: %v | %v", i, err)
			return err
		}
	}

	return nil
}

func (vas *vascaraMongoRepository) GetAllCategory() ([]model.CategoryFa, error) {
	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	collection := client.Database(mongoDBNam).Collection(collectionVascaraCategory)
	var category []model.CategoryFa
	x, err := collection.Find(context.TODO(), bson.M{})
	err = x.All(context.Background(), &category)

	return category, err

}
func (vas *vascaraMongoRepository) FindCategoryById(id int) (*model.CategoryChild, error) {
	var categoryChild model.CategoryChild

	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database(mongoDBNam).Collection(collectionVascaraCategory)

	error := collection.FindOne(
		context.TODO(),
		bson.D{primitive.E{Key: "id", Value: id}},
	).Decode(&categoryChild)
	if error != nil {
		logrus.Error("Could not find product: %v", error)
		return nil, error
	}
	return &categoryChild, nil
}

func (vas *vascaraMongoRepository) FindProductById(id int) (*model.ProductDetail, error) {
	var productDetail model.ProductDetail

	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database(mongoDBNam).Collection(collectionVascaraProduct)

	error := collection.FindOne(
		context.TODO(),
		bson.D{primitive.E{Key: "_id", Value: id}},
	).Decode(&productDetail)
	if error != nil {
		logrus.Error("Could not find product: %v", error)
		return nil, error
	}
	return &productDetail, nil
}
