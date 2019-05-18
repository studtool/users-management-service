package messages

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"

	"github.com/studtool/common/consts"

	"github.com/studtool/users-management-service/beans"
)

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
	beans.Logger().Info(
		fmt.Sprintf("message published (%s): %v", queueName, data),
	)
}

func (c *MqClient) writeMessagePublicationErrorLog(
	queueName string,
	data easyjson.Marshaler,
) {
	beans.Logger().Error(
		fmt.Sprintf("message not published (%s): %v", queueName, data),
	)
}

func (c *MqClient) writeMessageConsumedLog(
	queueName string,
	data easyjson.Marshaler,
) {
	beans.Logger().Info(
		fmt.Sprintf("message consumed (%s): %v", queueName, data),
	)
}

func (c *MqClient) writeMessageConsumptionErrorLog(
	queueName string,
	data easyjson.Marshaler,
) {
	beans.Logger().Error(
		fmt.Sprintf("message not consumed (%s): %v", queueName, data),
	)
}

func (c *MqClient) writeErrorLog(err error) {
	if err != nil {
		beans.Logger().Error(err)
	}
}
