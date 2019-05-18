package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/dig"

	"github.com/studtool/common/utils"

	"github.com/studtool/users-management-service/beans"
	"github.com/studtool/users-management-service/messages"
)

func main() {
	c := dig.New()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	signal.Notify(ch, syscall.SIGTERM)

	utils.AssertOk(c.Provide(messages.NewMqClient))
	utils.AssertOk(c.Invoke(func(q *messages.MqClient) {
		if err := q.OpenConnection(); err != nil {
			beans.Logger.Fatal(err)
		} else {
			beans.Logger.Info("queue: connection open")
		}
		if err := q.Run(); err != nil {
			beans.Logger.Fatal(err)
		} else {
			beans.Logger.Info("queue: ready to receive messages")
		}
	}))
	defer func() {
		utils.AssertOk(c.Invoke(func(q *messages.MqClient) {
			if err := q.CloseConnection(); err != nil {
				beans.Logger.Fatal(err)
			} else {
				beans.Logger.Info("queue: connection closed")
			}
		}))
	}()

	<-ch
}
