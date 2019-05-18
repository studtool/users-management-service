package main

import (
	"os"
	"os/signal"

	"go.uber.org/dig"

	"github.com/studtool/common/utils"

	"github.com/studtool/users-management-service/beans"
	"github.com/studtool/users-management-service/messages"
)

func main() {
	c := dig.New()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill)

	utils.AssertOk(c.Provide(messages.NewQueueClient))
	utils.AssertOk(c.Invoke(func(q *messages.QueueClient) {
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
		utils.AssertOk(c.Invoke(func(q *messages.QueueClient) {
			if err := q.CloseConnection(); err != nil {
				beans.Logger.Fatal(err)
			} else {
				beans.Logger.Info("queue: connection closed")
			}
		}))
	}()

	<-ch
}
