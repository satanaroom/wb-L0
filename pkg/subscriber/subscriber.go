package subscriber

import (
	"encoding/json"

	stan "github.com/nats-io/stan.go"
	broker "github.com/satanaroom/L0"
	"github.com/sirupsen/logrus"
)

func Subscribe(model *broker.Model) {
	// Подключение к серверу "prod"
	sc, err := stan.Connect("prod", "subscriber1", stan.NatsURL("nats://localhost:4222"))
	// Проверка на возможность подключения
	if err != nil {
		logrus.Fatalf("%s: %s", broker.NSError, err.Error())
	} else {
		logrus.Println(broker.NSSuccess)
	}
	defer sc.Close()

	// Подписываемся на канал для чтения данных (последней публикации)
	sub, err := sc.Subscribe("orders", func(m *stan.Msg) {
		err := json.Unmarshal(m.Data, model)
		if err != nil {
			logrus.Fatalf("transferred data isn't a json-object: %s", err.Error())
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		logrus.Fatalf("couldn't subscribe on orders channel: %s", err.Error())
	}
	defer sub.Unsubscribe()

	// w := sync.WaitGroup{}
	// w.Add(1)
	// w.Wait()
}
