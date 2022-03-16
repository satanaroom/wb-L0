package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	broker "github.com/satanaroom/L0"
	"github.com/satanaroom/L0/pkg/handler"
	"github.com/satanaroom/L0/pkg/repository"
	"github.com/satanaroom/L0/pkg/service"
	"github.com/satanaroom/L0/pkg/subscriber"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	// Объявление модели данных
	var m broker.Model

	// Установка формата логирования в виде JSON
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Ининциализация кофига
	err := initConfig()
	if err != nil {
		logrus.Fatalf("[cfg] failed to initialize configs: %s", err.Error())
	}

	// Загрузка переменных окружения
	err = godotenv.Load()
	if err != nil {
		logrus.Fatalf("[env] failed to load env variables: %s", err.Error())
	}
	// Инициализация базы данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize database: %s", err.Error())
	}
	logrus.Println("[db] initialized success")
	repos := repository.NewRepository(db)
	// Подписка на канал
	subscriber.Subscribe(&m, repos)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(broker.Server)
	go func() {
		err := srv.Run(viper.GetString("port"), handlers.InitRoutes())
		if err != nil {
			logrus.Fatalf("error occured while http server is running: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	err = srv.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	err = db.Close()
	if err != nil {
		logrus.Errorf("failed to db close: %s", err.Error())
	}

}
