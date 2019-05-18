package messages

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/studtool/common/queues"
	"github.com/studtool/common/utils"

	"github.com/studtool/users-management-service/beans"
	"github.com/studtool/users-management-service/config"
)

type MqClient struct {
	connStr    string
	connection *amqp.Connection

	channel *amqp.Channel

	createdUsersQueue amqp.Queue
	deletedUsersQueue amqp.Queue

	profilesToCreateQueue amqp.Queue
	profilesToDeleteQueue amqp.Queue

	avatarsToCreateQueue amqp.Queue
	avatarsToDeleteQueue amqp.Queue

	documentUsersToCreateQueue amqp.Queue
	documentUsersToDeleteQueue amqp.Queue
}

func NewMqClient() *MqClient {
	return &MqClient{
		connStr: fmt.Sprintf("amqp://%s:%s@%s:%d/",
			config.MqUser.Value(), config.MqPassword.Value(),
			config.MqHost.Value(), config.MqPort.Value(),
		),
	}
}

func (c *MqClient) OpenConnection() error {
	var conn *amqp.Connection
	err := utils.WithRetry(func(n int) (err error) {
		if n > 0 {
			beans.Logger().Info(fmt.Sprintf("opening message queue connection. retry #%d", n))
		}
		conn, err = amqp.Dial(c.connStr)
		return err
	}, config.MqConnNumRet.Value(), config.MqConnRetItv.Value())
	if err != nil {
		return err
	}

	c.connection = conn

	c.channel, err = conn.Channel()
	if err != nil {
		return err
	}

	c.createdUsersQueue, err =
		c.declareQueue(queues.CreatedUsersQueueName)
	if err != nil {
		return err
	}

	c.deletedUsersQueue, err =
		c.declareQueue(queues.DeletedUsersQueueName)
	if err != nil {
		return err
	}

	c.profilesToCreateQueue, err =
		c.declareQueue(queues.ProfilesToCreateQueueName)
	if err != nil {
		return err
	}

	c.profilesToDeleteQueue, err =
		c.declareQueue(queues.ProfilesToDeleteQueueName)
	if err != nil {
		return err
	}

	c.avatarsToCreateQueue, err =
		c.declareQueue(queues.AvatarsToCreateQueueName)
	if err != nil {
		return err
	}

	c.avatarsToDeleteQueue, err =
		c.declareQueue(queues.AvatarsToDeleteQueueName)
	if err != nil {
		return err
	}

	c.documentUsersToCreateQueue, err =
		c.declareQueue(queues.DocumentUsersToCreateQueueName)
	if err != nil {
		return err
	}

	c.documentUsersToDeleteQueue, err =
		c.declareQueue(queues.DocumentUsersToDeleteQueueName)
	if err != nil {
		return err
	}

	return nil
}

func (c *MqClient) CloseConnection() error {
	if err := c.channel.Close(); err != nil {
		return err
	}
	return c.connection.Close()
}

type MessageHandler func(data []byte)

func (c *MqClient) Run() error {
	err := c.runConsumer(
		queues.CreatedUsersQueueName,
		c.handleUserCreated,
	)
	if err != nil {
		return err
	}

	err = c.runConsumer(
		queues.DeletedUsersQueueName,
		c.handleUserDeleted,
	)
	if err != nil {
		return err
	}

	return nil
}

type messageHandler func(data []byte)
