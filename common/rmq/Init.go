package rmq

import (
	"context"
	"fmt"
	"time"

	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

const (
	AccessKey		= "test"
	SecretKey   	= "test"
	ConsumerGroup	= "gomall_group_"
	Topic			= "gomall_topic_"
)

type RMQClient struct {
    producer       rmq_client.Producer
    simpleConsumer rmq_client.SimpleConsumer
}

var (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// receive messages in a loop
)

func Init(topicPrefix string, groupPrefix string, Endpoint string) (*RMQClient, error) {
	producer, err := rmq_client.NewProducer(&rmq_client.Config{
		Endpoint: Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithTopics(Topic + topicPrefix),
		rmq_client.WithMaxAttempts(3),
	)
	if err != nil {
		fmt.Println("Init Producer failed")
		klog.Error("Init Producer failed")
		return &RMQClient{}, kerrors.NewBizStatusError(400, "Init Producer failed")
	}

	simpleConsumer, err := rmq_client.NewSimpleConsumer(&rmq_client.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: ConsumerGroup + groupPrefix,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithAwaitDuration(awaitDuration),
		rmq_client.WithSubscriptionExpressions(map[string]*rmq_client.FilterExpression{
			Topic: rmq_client.SUB_ALL,
		}),
	)

	if err != nil {
		fmt.Println("Init Consumer failed")
		klog.Error("Init Consumer failed")
		return &RMQClient{}, kerrors.NewBizStatusError(400, "Init Consumer failed")
	}

	return &RMQClient{
		producer: producer,
		simpleConsumer: simpleConsumer, 
	}, err
}

func (c *RMQClient) Close() {
	c.producer.GracefulStop()
	c.simpleConsumer.GracefulStop()
}

func (c *RMQClient) SendMsgAsync(ctx context.Context, message string, topicPrefix string, key string, tag string) {
	msg := &rmq_client.Message{
		Topic: Topic + topicPrefix,
		Body: []byte(message),
	}
	msg.SetKeys(key)
	msg.SetTag(tag)
	c.producer.SendAsync(context.TODO(), msg, func(ctx context.Context, resp []*rmq_client.SendReceipt, err error) {
		if err != nil {
			fmt.Printf("SendAsnyc failed\n")
			klog.Error("SendAsnyc failed")
		}
		for i := 0; i < len(resp); i++ {
			fmt.Printf("%#v\n", resp[i])
		}
	})
}

func (c *RMQClient) SendMsgSync(ctx context.Context, message string, topicPrefix string, key string, tag string) error {
	msg := &rmq_client.Message{
		Topic: Topic + topicPrefix,
		Body: []byte(message),
	}
	msg.SetKeys(key)
	msg.SetTag(tag)
	_, err := c.producer.Send(ctx, msg)
	if err != nil {
		fmt.Printf("SendAsnyc failed\n")
		klog.Error("SendAsnyc failed")
	}
	return err
}

func (c *RMQClient) ReceiveMsg(ctx context.Context) {
	for {
		msgs, err := c.simpleConsumer.Receive(ctx, maxMessageNum, invisibleDuration)
		fmt.Printf("Receive message %+v\n", msgs)
		if err != nil {
			klog.Error("Receive msg failed")
			fmt.Println("Receive msg failed")
		}
		
		for _, mv := range msgs {
			c.simpleConsumer.Ack(ctx, mv)
		}
	}
}

