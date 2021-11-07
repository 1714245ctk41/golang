package repo

import (
	"crawl_data/pkg/model"

	"github.com/olivere/elastic/v7"
)

type VascaraElasticRepository interface {
	InsertProduct(product *model.ProductDetail, client *elastic.Client) error
}

type vascaraElasticRepository struct {
	elasticClient *elastic.Client
}

func NewVascaraElasticRepository(elasticClient *elastic.Client) VascaraElasticRepository {
	return &vascaraElasticRepository{
		elasticClient: elasticClient,
	}
}

func (elas *vascaraElasticRepository) InsertProduct(product *model.ProductDetail, client *elastic.Client) error {
	return nil
}
