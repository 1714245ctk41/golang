package conf

import (
	"github.com/caarlos0/env/v6"
)

// AppConfig presents app conf
type AppConfig struct {
	Port string `env:"PORT" envDefault:"8081"`

	DBHost   string `env:"DB_HOST" envDefault:"localhost"`
	DBPort   string `env:"DB_PORT" envDefault:"5432"`
	DBUser   string `env:"DB_USER" envDefault:"root"`
	DBPass   string `env:"DB_PASS" envDefault:"2520"`
	DBName   string `env:"DB_NAME" envDefault:"crawl_golang"`
	DBSchema string `env:"DB_SCHEMA" envDefault:""`
	EnableDB string `env:"ENABLE_DB" envDefault:"true"`
	AMQP_URL string `env:"AMQP_URL" envDefault:"amqp://guest:guest@localhost:5672"`

	MongoDBName     string `env:"MONGODB_NAME" envDefault:"golang_mongodb_api"`
	MongoConnectStr string `env:"MONGO_CONNECT_STR" envDefault:"mongodb+srv://admin:admin@cluster0.uigy4.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"`

	MsConsumerURL string `env:"MS_CONSUMER_URL" envDefault:"http://localhost:8081/api/v2/consumer/"`
}

var config AppConfig

func SetEnv() {
	_ = env.Parse(&config)
}

func LoadEnv() AppConfig {
	return config
}
