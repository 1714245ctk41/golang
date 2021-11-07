package service

import (
	"context"
	"crawl_data/pkg/model"
	"crawl_data/pkg/repo"
	"crawl_data/pkg/utils"
	rabbitmq_go "crawl_data/pkg/utils/rabbitmq_go"
	"log"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

type VascaraService interface {
	InsertCategory() ([]model.CategoryChild, error)
	ReadMessage()
	InsertProduct(urlStr string, client *mongo.Client, productBestSale []string, count *int) error
	GetProductReview() ([]model.ProductView, error)
}

var count int = 0
var ProductView []model.ProductView

type vascaraService struct {
	vascaraRepository repo.VascaraMongoRepository
}

func NewVascaraService(vascaraRepository repo.VascaraMongoRepository) VascaraService {
	return &vascaraService{
		vascaraRepository: vascaraRepository,
	}
}

func (vas *vascaraService) InsertProduct(urlStr string, client *mongo.Client, productBestSale []string, count *int) error {
	productDetail := rabbitmq_go.ProductDetailVascava(string(urlStr))
	if productDetail.ProductCode == "" && productDetail.Title == "" {
		logrus.Error("Sản phẩm rỗng")
		return nil
	}
	if *count <= len(productBestSale) {
		for _, x := range productBestSale {
			if strings.Contains(productDetail.Title, x) {
				*count++
				productDetail.BestSale = true
				break
			}
		}
	}
	err := vas.vascaraRepository.InsertProduct(&productDetail, client)
	if err != nil {
		logrus.Error("Failed insert product to mongo")
		return err
	}
	return nil
}

func (vas *vascaraService) GetAllCategory() {

}

func (vas *vascaraService) InsertCategory() ([]model.CategoryChild, error) {
	cate := rabbitmq_go.GetCategoryContainer("https://www.vascara.com/")
	cateCollecFa, cateCollecChild := rabbitmq_go.GetCategory(cate)
	err := vas.vascaraRepository.InsertCategory(cateCollecFa)
	if err != nil {
		logrus.Error("Error insert category: %v", err)
		return nil, err
	}
	return cateCollecChild, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (vas *vascaraService) ReadMessage() {
	productBestSale := rabbitmq_go.BestSaleHandle()
	var count *int
	count = new(int)
	*count = 0

	vas.CreateReadMessage("vascara_crawl", productBestSale, count)

}

func (vas *vascaraService) CreateReadMessage(nameQueue string, productBestSale []string, count *int) {
	conn := rabbitmq_go.Connect()
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		nameQueue, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)
	client := utils.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	queueNumber := 15
	for i := 0; i <= queueNumber; i++ {
		go vas.Consumer(msgs, client, productBestSale, count)

	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
func (vas *vascaraService) Consumer(msgs <-chan amqp.Delivery, client *mongo.Client, productBestSale []string, count *int) {
	for d := range msgs {
		urlStr := d.Body
		vas.InsertProduct(string(urlStr), client, productBestSale, count)
		log.Printf("Received a message: %s", d.Body)
	}
}
func (vas *vascaraService) GetProductReview() ([]model.ProductView, error) {
	category, err := vas.InsertCategory()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	ProductView := rabbitmq_go.GetProductReview(category)
	return ProductView, nil
}
