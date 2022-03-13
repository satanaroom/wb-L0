package main

import (
	"io/ioutil"
	"os"

	stan "github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func main() {
	// Проверка количества переданных аргументов
	if len(os.Args) == 1 {
		logrus.Fatalln("no files to publish on server. use: go run publisher.go fileName")
	}
	fileName := os.Args[1]
	// Получение модели слайсом байт, переданным первым аргументом для отправки в канал
	m, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Fatalf("error occured with file: %s", fileName)
	}

	// Подключение к серверу "prod"
	sc, err := stan.Connect("prod", "client-123", stan.NatsURL("nats://localhost:4222"))
	// Проверка на возможность подключения
	if err != nil {
		logrus.Fatalf("couldn't connect to nats-streaming: %s", err.Error())
	} else {
		logrus.Println("connection to nats-streaming success")
	}
	defer sc.Close()

	// Добавление данных в канал "orders"
	sc.Publish("orders", m)
}
