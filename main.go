package main

import (
	handlers2 "awesomeProject7/handlers"
	"awesomeProject7/repository"
	"awesomeProject7/server"
	services2 "awesomeProject7/services"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logger := logrus.Logger{}
	logger.Info("create router")
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	c := cron.New()
	c.AddFunc("CRON_TZ=Europe/Kyiv 38 18 * * *", services2.Sender)
	c.Start()

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loaing env variables:%s", err.Error())
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

	repo := repository.NewRepository(db)
	services := services2.NewService(repo)
	handlers := handlers2.NewHandler(services, logger)
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
