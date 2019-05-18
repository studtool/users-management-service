package messages

import (
	"github.com/studtool/common/queues"

	"github.com/studtool/users-management-service/beans"
)

func (c *QueueClient) handleUserCreated(body []byte) {
	data := &queues.CreatedUserData{}
	if err := c.parseMessageBody(body, data); err != nil {
		c.handleErr(err)
	} else {
		// TODO
	}
}

func (c *QueueClient) handleUserDeleted(body []byte) {
	data := &queues.DeletedUserData{}
	if err := c.parseMessageBody(body, data); err != nil {
		c.handleErr(err)
	} else {
		// TODO
	}
}

func (c *QueueClient) handleErr(err error) {
	if err != nil {
		beans.Logger.Error(err)
	}
}
