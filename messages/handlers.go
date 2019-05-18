package messages

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/queues"

	"github.com/studtool/users-management-service/beans"
)

func (c *MqClient) handleUserCreated(body []byte) {
	createdUserData := &queues.CreatedUserData{}
	if err := c.unmarshalMessageBody(body, createdUserData); err != nil {
		c.writeErrorLog(err)
	} else {
		c.publishProfileToCreateMessage(createdUserData)
		c.publishDocumentUserToCreateMessage(createdUserData)
	}
}

func (c *MqClient) handleUserDeleted(body []byte) {
	deletedUserData := &queues.DeletedUserData{}
	if err := c.unmarshalMessageBody(body, deletedUserData); err != nil {
		c.writeErrorLog(err)
	} else {
		c.publishProfileToDeleteMessage(deletedUserData)
		c.publishDocumentUserToDeletedMessage(deletedUserData)
	}
}

func (c *MqClient) publishProfileToCreateMessage(
	createdUserData *queues.CreatedUserData,
) {
	c.publishMessage(c.profilesToCreateQueue.Name,
		&queues.ProfileToCreateData{
			UserID: createdUserData.UserID,
		},
	)
}

func (c *MqClient) publishProfileToDeleteMessage(
	deletedUserData *queues.DeletedUserData,
) {
	c.publishMessage(c.profilesToDeleteQueue.Name,
		&queues.ProfileToDeleteData{
			UserID: deletedUserData.UserID,
		},
	)
}

func (c *MqClient) publishDocumentUserToCreateMessage(
	createdUserData *queues.CreatedUserData,
) {
	c.publishMessage(c.documentUsersToCreateQueue.Name,
		&queues.ProfileToCreateData{
			UserID: createdUserData.UserID,
		},
	)
}

func (c *MqClient) publishDocumentUserToDeletedMessage(
	deletedUserData *queues.DeletedUserData,
) {
	c.publishMessage(c.documentUsersToDeleteQueue.Name,
		&queues.ProfileToDeleteData{
			UserID: deletedUserData.UserID,
		},
	)
}

func (c *MqClient) publishMessage(
	queueName string,
	data easyjson.Marshaler,
) {
	body, err := c.marshalMessageBody(data)

	err = c.channel.Publish(
		consts.EmptyString,
		queueName,
		false,
		false,
		amqp.Publishing{
			Body:        body,
			ContentType: "application/json",
		},
	)
	if err == nil {
		c.writeMessagePublishedLog(queueName, data)
	} else {
		c.writeMessagePublicationErrorLog(queueName, data)
	}
}

func (c *MqClient) writeMessagePublishedLog(
	queueName string,
	data easyjson.Marshaler,
) {
	beans.Logger.Info(
		fmt.Sprintf("message published (%s): %v", queueName, data),
	)
}

func (c *MqClient) writeMessagePublicationErrorLog(
	queueName string,
	data easyjson.Marshaler,
) {
	beans.Logger.Error(
		fmt.Sprintf("message not published (%s): %v", queueName, data),
	)
}

func (c *MqClient) writeErrorLog(err error) {
	if err != nil {
		beans.Logger.Error(err)
	}
}
