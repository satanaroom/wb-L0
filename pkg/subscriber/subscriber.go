package subscriber

import (
	"encoding/json"

	stan "github.com/nats-io/stan.go"
	broker "github.com/satanaroom/L0"
	"github.com/satanaroom/L0/pkg/repository"
	"github.com/sirupsen/logrus"
)

func Subscribe(model *broker.Model, repos *repository.Repository) {
	// Подключение к серверу "prod"
	sc, err := stan.Connect("prod", "subscriber1", stan.NatsURL("nats://localhost:4222"))
	// Проверка на возможность подключения
	if err != nil {
		logrus.Fatalf("%s: %s", broker.NSError, err.Error())
	} else {
		logrus.Println(broker.NSSuccess)
	}

	// Подписываемся на канал для чтения данных (последней публикации)
	_, err = sc.Subscribe("orders", func(m *stan.Msg) {
		err := json.Unmarshal(m.Data, model)
		if err != nil {
			logrus.Fatalf("transferred data isn't a json-object: %s", err.Error())
		}
		repos.CreateModelMain(*model)
		repos.CreateModelDeliveries(*model)
	}, stan.StartWithLastReceived())
	if err != nil {
		logrus.Fatalf("couldn't subscribe on orders channel: %s", err.Error())
	}
}
