package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	broker "github.com/satanaroom/L0"
	"github.com/satanaroom/L0/pkg/repository"
	"github.com/satanaroom/L0/pkg/subscriber"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// func htmlHandler() {

// }

// func InitRoutes() *gin.Engine {
// 	router := gin.New()
// 	orders := router.Group("/orders")
// 	{
// 		orders.GET("/", htmlHandler())
// 		orders.POST("/", )
// 	}

// 	return router
// }

func main() {
	// Объявление модели данных
	var m broker.Model

	// Установка формата логирования в виде JSON
	logrus.SetFormatter(new(logrus.JSONFormatter))
	// Подписка на канал
	subscriber.Subscribe(&m)
	fmt.Println(m)
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
	// srv := new(broker.Server)

	// go func() {
	// 	if err := srv.Run(viper.GetString("port"), InitRoutes()); err != nil {
	// 		logrus.Fatalf("error occured while http server is running: %s", err.Error())
	// 	}
	// }()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// <-quit

	// if err := srv.Shutdown(context.Background()); err != nil {
	// 	logrus.Errorf("error occured on server shutting down: %s", err.Error())
	// }

	if err := db.CloseDB(); err != nil {
		logrus.Errorf("error occured on database connection close: %s", err.Error())
	}

}
