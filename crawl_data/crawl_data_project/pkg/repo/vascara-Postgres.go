package repo

import (
	"crawl_data/pkg/model"

	"gorm.io/gorm"
)

type VascaraPostgresRepository interface {
	InsertProduct(product *model.ProductDetail) error
	InsertCategory(category *model.CategoryFa) error
	FindProductById(id uint) error
	DeleteAllCategory(id uint) error
}

type vascaraPostgresRepository struct {
	connection *gorm.DB
}

func NewVascaraPostgresRepository(db *gorm.DB) VascaraPostgresRepository {
	return &vascaraPostgresRepository{
		connection: db,
	}
}

func (db *vascaraPostgresRepository) InsertProduct(product *model.ProductDetail) error {
	err := db.connection.Model(&model.ProductDetail{}).Where("id=?", product.ID).Save(product).Error
	// err := db.connection.wher
	if err != nil {
		return err
	}
	return nil
}
func (db *vascaraPostgresRepository) InsertCategory(category *model.CategoryFa) error {
	err := db.connection.Model(&model.CategoryFa{}).Where("id=?", category.ID).Save(category).Error
	// Create(&model.CategoryFa{
	// 	ID:             category.ID,
	// 	Name:           category.Name,
	// 	Link:           category.Link,
	// 	LinkCategory:   category.LinkCategory,
	// 	CategoryChilds: category.CategoryChilds,
	// 	BaseModel:      model.BaseModel{},
	// }).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *vascaraPostgresRepository) FindProductById(id uint) error {
	result := map[string]interface{}{}
	err := db.connection.Model(&model.ProductDetail{}).First(&result, "ID = ?", 1).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *vascaraPostgresRepository) DeleteAllCategory(id uint) error {
	if err := db.connection.Delete(&model.CategoryFa{}, id).Error; err != nil {
		return err
	}
	return nil
}
