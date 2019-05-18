package messages

import (
	"github.com/studtool/common/queues"
)

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

func (c *MqClient) publishAvatarToCreateMessage(
	createdUserData *queues.CreatedUserData,
) {
	c.publishMessage(c.avatarsToCreateQueue.Name,
		&queues.AvatarToCreateData{
			UserID: createdUserData.UserID,
		},
	)
}

func (c *MqClient) publishAvatarToDeleteMessage(
	deletedUserData *queues.DeletedUserData,
) {
	c.publishMessage(c.avatarsToDeleteQueue.Name,
		&queues.AvatarToDeleteData{
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
