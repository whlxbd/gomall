package mq

import (
    "context"
    "encoding/json"
    "time"
    rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/whlxbd/gomall/common/rmq"
)

type OrderMessage struct {
    OrderId string `json:"order_id"`
    UserId  uint32 `json:"user_id"`
}

func SendOrderTimeoutMessage(ctx context.Context, orderId string, userId uint32) error {
    msg := &OrderMessage{
        OrderId: orderId,
        UserId:  userId,
    }
    
    data, err := json.Marshal(msg)
    if err != nil {
        return err
    }

    message := &rmq_client.Message{
        Topic: rmq.Topic + "order",
        Body:  data,
    }
    // 15分钟后触发
    message.SetDelayTimestamp(time.Now().Add(time.Minute * 1))
    
    message.SetKeys(orderId)
    message.SetTag("timeout")
    
    _, err = orderMQ.producer.Send(ctx, message)
    return err
}