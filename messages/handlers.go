package messages

import (
	"github.com/studtool/common/queues"
)

type messageHandler func(data []byte)

func (c *MqClient) handleUserCreated(body []byte) {
	createdUserData := &queues.CreatedUserData{}
	if err := c.unmarshalMessageBody(body, createdUserData); err != nil {
		c.writeErrorLog(err)
	} else {
		c.writeMessageConsumedLog(c.createdUsersQueue.Name, createdUserData)

		c.publishProfileToCreateMessage(createdUserData)
		c.publishAvatarToCreateMessage(createdUserData)
		c.publishDocumentUserToCreateMessage(createdUserData)
	}
}

func (c *MqClient) handleUserDeleted(body []byte) {
	deletedUserData := &queues.DeletedUserData{}
	if err := c.unmarshalMessageBody(body, deletedUserData); err != nil {
		c.writeErrorLog(err)
	} else {
		c.writeMessageConsumedLog(c.createdUsersQueue.Name, deletedUserData)

		c.publishProfileToDeleteMessage(deletedUserData)
		c.publishAvatarToDeleteMessage(deletedUserData)
		c.publishDocumentUserToDeletedMessage(deletedUserData)
	}
}
