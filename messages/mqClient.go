package messages

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"

	"github.com/studtool/common/consts"
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

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	c.createdUsersQueue, err = ch.QueueDeclare(
		queues.CreatedUsersQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.deletedUsersQueue, err = ch.QueueDeclare(
		queues.DeletedUsersQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.profilesToCreateQueue, err = ch.QueueDeclare(
		queues.ProfilesToCreateQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.profilesToDeleteQueue, err = ch.QueueDeclare(
		queues.ProfilesToDeleteQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.documentUsersToCreateQueue, err = ch.QueueDeclare(
		queues.DocumentUsersToCreateQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.documentUsersToDeleteQueue, err = ch.QueueDeclare(
		queues.DocumentUsersToDeleteQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.channel = ch
	c.connection = conn

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
	if err := c.recvCreatedUsersData(); err != nil {
		return err
	}
	if err := c.recvDeletedUsersData(); err != nil {
		return err
	}
	return nil
}

func (c *MqClient) recvCreatedUsersData() error {
	messages, err := c.channel.Consume(
		c.createdUsersQueue.Name,
		consts.EmptyString,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range messages {
			c.handleUserCreated(d.Body)
		}
	}()

	return nil
}

func (c *MqClient) recvDeletedUsersData() error {
	messages, err := c.channel.Consume(
		c.deletedUsersQueue.Name,
		consts.EmptyString,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range messages {
			c.handleUserDeleted(d.Body)
		}
	}()

	return nil
}

func (c *MqClient) marshalMessageBody(v easyjson.Marshaler) ([]byte, error) {
	return easyjson.Marshal(v)
}

func (c *MqClient) unmarshalMessageBody(data []byte, v easyjson.Unmarshaler) error {
	return easyjson.Unmarshal(data, v)
}
