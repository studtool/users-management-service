package messages

import (
	"github.com/studtool/common/queues"
)

func (c *MqClient) handleUserCreated(body []byte) {
	createdUserData := &queues.CreatedUserData{}
	if err := c.unmarshalMessageBody(body, createdUserData); err != nil {
		c.writeErrorLog(err)
	} else {
		c.publishProfileToCreateMessage(createdUserData)
		c.publishDocumentUserToCreateMessage(createdUserData)

		c.writeMessageConsumedLog(c.createdUsersQueue.Name, createdUserData)
	}
}

func (c *MqClient) handleUserDeleted(body []byte) {
	deletedUserData := &queues.DeletedUserData{}
	if err := c.unmarshalMessageBody(body, deletedUserData); err != nil {
		c.writeErrorLog(err)
	} else {
		c.publishProfileToDeleteMessage(deletedUserData)
		c.publishDocumentUserToDeletedMessage(deletedUserData)

		c.writeMessageConsumedLog(c.createdUsersQueue.Name, deletedUserData)
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
