package cfg

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Cfg struct {
	Server struct {
		BindIP     string `env:"BIND_IP" env-default:"127.0.0.1"`
		ListenPort string `env:"LISTEN_PORT" env-default:"8080"`
	}
	Postgresql struct {
		Host     string `env:"PSQL_HOST"`
		Port     string `env:"PSQL_PORT"`
		Database string `env:"PSQL_DBNAME"`
		Username string `env:"PSQL_USER"`
		Password string `env:"PSQL_PASSWORD"`
	}
	JWT struct {
		SecretKey string `env:"JWT_SECRET_KEY"`
		ExpTime   int    `env:"JWT_TOKEN_EXP_TIME"`
	}
}

var instance *Cfg
var once sync.Once

func GetConfig() *Cfg {
	once.Do(func() {
		log.Println("load env from file")

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("error load file .env:", err)
		}

		log.Println("read app configuration")
		instance = &Cfg{}
		err = cleanenv.ReadEnv(instance)
		if err != nil {
			help, err := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatalln(err)
		}
	})
	return instance
}
