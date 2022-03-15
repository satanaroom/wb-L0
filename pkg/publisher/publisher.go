package main

import (
	"io/ioutil"
	"os"

	stan "github.com/nats-io/stan.go"
	broker "github.com/satanaroom/L0"
	"github.com/sirupsen/logrus"
)

func main() {
	// Проверка количества переданных аргументов
	if len(os.Args) == 1 {
		logrus.Fatalln(broker.FError)
	}
	fileName := os.Args[1]
	// Получение модели слайсом байт, переданным первым аргументом для отправки в канал
	m, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Fatalf("error occured with file: %s", fileName)
	}

	// Подключение к серверу "prod"
	sc, err := stan.Connect("prod", "publisher")
	// Проверка на возможность подключения
	if err != nil {
		logrus.Fatalf("%s %s", broker.NSError, err.Error())
	} else {
		logrus.Println(broker.NSSuccess)
	}
	defer sc.Close()

	// Добавление данных в канал "orders"
	sc.Publish("orders", m)
}
