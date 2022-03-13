package main

import (
	"github.com/satanaroom/L0/subscriber"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	subscriber.Subscribe()
}
