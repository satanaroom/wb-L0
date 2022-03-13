package subscriber

import (
	"fmt"
	"sync"

	stan "github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func Subscribe() {
	sc, err := stan.Connect("prod", "user1", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		logrus.Fatalf("couldn't connect to nats-streaming: %s", err.Error())
	} else {
		logrus.Println("connection to nats-streaming success")
	}
	defer sc.Close()

	sub, err := sc.Subscribe("orders", func(m *stan.Msg) {
		fmt.Printf("Вот что опубликовали: %s\n", string(m.Data))
	}, stan.StartWithLastReceived())
	if err != nil {
		logrus.Fatalf("couldn't subscribe on orders channel: %s", err.Error())
	}
	defer sub.Unsubscribe()

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
