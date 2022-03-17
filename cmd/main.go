package main

// Защита от потери данных в случае ошибок реализована путем создания
// директории с файлами, в которые пишутся все публикации
// nats-streaming-server -cid prod -store file -dir store

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	broker "github.com/satanaroom/L0"
	"github.com/satanaroom/L0/cache"
	"github.com/satanaroom/L0/pkg/handler"
	"github.com/satanaroom/L0/pkg/repository"
	"github.com/satanaroom/L0/pkg/service"
	"github.com/satanaroom/L0/pkg/subscriber"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	// Добавляем путь, по которому нужно искать файлы конфигов
	viper.AddConfigPath("configs")
	// Добавляем имя файла
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
		logrus.Fatalf("[db] failed to initialize database: %s", err.Error())
	}
	logrus.Println("[db] initialized success")

	// Создание кэша
	cache := cache.New(5*time.Minute, 10*time.Minute)
	// Передача БД и кэша в слой репозитория для его создания
	repos := repository.NewRepository(db, cache)
	// Подписка на канал с передачей БД и кэша для записи
	subscriber.Subscribe(&m, repos, cache)
	// Передача репозитория в слой сервисов
	services := service.NewService(repos)
	// Передача сервисов в слой обработчиков
	handlers := handler.NewHandler(services)
	// Инициализация http-сервера
	srv := new(broker.Server)
	// Запуск http-сервера
	go func() {
		err := srv.Run(viper.GetString("port"), handlers.InitRoutes())
		if err != nil {
			logrus.Fatalf("[srv] failed to server running: %s", err.Error())
		}
	}()

	// Инициализация канала, ожидающего сигнал завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Остановка сервера
	err = srv.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("[srv] failed to server shutting down: %s", err.Error())
	}

	// Закрытие БД
	err = db.Close()
	if err != nil {
		logrus.Errorf("[db] failed to db close: %s", err.Error())
	}
}
