package main

import (
	"os"
	"restapi/internal"
	"restapi/pkg/handler"
	"restapi/pkg/repository"
	"restapi/pkg/service"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"restapi/pkg/service/wb"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load("C:\\Users\\Нуржан\\OneDrive\\Рабочий стол\\Nurzhan\\school\\programming\\go\\go\\src\\assignment\\homeworks\\poi\\.env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	hub := wb.NewHub()
	wsHandler := wb.NewHandler(hub)
	go hub.Run()

	srv := new(internal.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes(wsHandler)); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())

	}
}

func initConfig() error {
	viper.AddConfigPath("C:\\Users\\Нуржан\\OneDrive\\Рабочий стол\\Nurzhan\\school\\programming\\go\\go\\src\\assignment\\homeworks\\poi\\configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
