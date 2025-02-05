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

type Producer struct {
	producer rmq_client.Producer
}

type SimpleConsumer struct {
	simpleConsumer rmq_client.SimpleConsumer
}

var (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 60
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 30
	// receive messages in a loop
)

func InitProducer(topicSuffix string, Endpoint string) (*Producer, error) {
	topicName := Topic + topicSuffix

	producer, err := rmq_client.NewProducer(&rmq_client.Config{
		Endpoint: Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithTopics(topicName),
		rmq_client.WithMaxAttempts(3),
	)
	if err != nil {
		fmt.Println("Init Producer failed: ", err)
		klog.Error("Init Producer failed： ", err)
		return nil, kerrors.NewBizStatusError(400, "Init Producer failed")
	}

	if err := producer.Start(); err != nil {
		producer.GracefulStop()
        return nil, fmt.Errorf("start producer failed: %w", err)
    }
	return &Producer{producer: producer}, err
}

func InitConsumer(topicSuffix string, groupSuffix string, Endpoint string) (*SimpleConsumer, error) {
	topicName := Topic + topicSuffix

	simpleConsumer, err := rmq_client.NewSimpleConsumer(&rmq_client.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: ConsumerGroup + groupSuffix,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithAwaitDuration(awaitDuration),
		rmq_client.WithSubscriptionExpressions(map[string]*rmq_client.FilterExpression{
			topicName: rmq_client.SUB_ALL,
		}),// 订阅所有消息
	)
	if err != nil {
		fmt.Println("Init Consumer failed")
		klog.Error("Init Consumer failed")
		return nil, kerrors.NewBizStatusError(400, "Init Consumer failed")
	}

    // 启动consumer
    if err := simpleConsumer.Start(); err != nil {
        simpleConsumer.GracefulStop()
        return nil, fmt.Errorf("start consumer failed: %w", err)
    }

	return &SimpleConsumer{simpleConsumer: simpleConsumer}, err
}

func (c *Producer) Close() {
	if c != nil && c.producer != nil {
		c.producer.GracefulStop()
	}
}

func (c *SimpleConsumer) Close() {
	if c != nil && c.simpleConsumer != nil {
		c.simpleConsumer.GracefulStop()
	}
}

func (c *Producer) SendMsgAsync(ctx context.Context, message string, topicSuffix string, key string, tag string) {
	msg := &rmq_client.Message{
		Topic: Topic + topicSuffix,
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

func (c *Producer) SendMsgSync(ctx context.Context, message string, topicSuffix string, key string, tag string) error {
	msg := &rmq_client.Message{
		Topic: Topic + topicSuffix,
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

func (c *SimpleConsumer) ReceiveMsg(ctx context.Context) {
	for {
		msgs, err := c.simpleConsumer.Receive(ctx, maxMessageNum, invisibleDuration)
		fmt.Printf("Receive message %+v\n", msgs)
		if err != nil {
			klog.Error("Receive msg failed: ", err)
			fmt.Println("Receive msg failed： ", err)
		}
		
		for _, mv := range msgs {
			if ackErr := c.simpleConsumer.Ack(ctx, mv); ackErr != nil {
				klog.Error("Ack msg failed: ", ackErr)
			}
			fmt.Println(mv)
		}
	}
}
