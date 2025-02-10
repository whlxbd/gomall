package rmq

import (
	"context"
	"fmt"
	"sync"
	"time"

	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

const (
	AccessKey     = "test"
	SecretKey     = "test"
	ConsumerGroup = "gomall_group_"
	Topic         = "gomall_topic_"
	// maximum waiting time for receive func
	AwaitDuration = time.Second * 60
	// maximum number of messages received at one time
	MaxMessageNum int32 = 16
	// invisibleDuration should > 20s
	InvisibleDuration = time.Second * 30
	// receive messages in a loop
)

var (
	closeMutex sync.Mutex
)

func InitProducer(topicSuffix string, Endpoint string) (rmq_client.Producer, error) {
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
		defer producer.GracefulStop()
		return nil, fmt.Errorf("start producer failed: %w", err)
	}
	return producer, err
}

func InitConsumer(topicSuffix string, groupSuffix string, Endpoint string) (rmq_client.SimpleConsumer, error) {
	topicName := Topic + topicSuffix

	simpleConsumer, err := rmq_client.NewSimpleConsumer(&rmq_client.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: ConsumerGroup + groupSuffix,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithAwaitDuration(AwaitDuration),
		rmq_client.WithSubscriptionExpressions(map[string]*rmq_client.FilterExpression{
			topicName: rmq_client.SUB_ALL,
		}), // 订阅所有消息
	)
	if err != nil {
		fmt.Println("Init Consumer failed")
		klog.Error("Init Consumer failed")
		return nil, kerrors.NewBizStatusError(400, "Init Consumer failed")
	}

	// 启动consumer
	if err := simpleConsumer.Start(); err != nil {
		defer simpleConsumer.GracefulStop()
		return nil, fmt.Errorf("start consumer failed: %w", err)
	}

	return simpleConsumer, err
}

func Close(p rmq_client.Producer, c rmq_client.SimpleConsumer) (err error) {
	closeMutex.Lock()
	defer closeMutex.Unlock()

	if p != nil {
		err = p.GracefulStop()
	}
	if err != nil {
		return
	}

	if c != nil {
		err = c.GracefulStop()
	}

	return 
}

func SendMsgAsync(producer rmq_client.Producer , ctx context.Context, message string, topicSuffix string, key string, tag string) {
	msg := &rmq_client.Message{
		Topic: Topic + topicSuffix,
		Body:  []byte(message),
	}
	msg.SetKeys(key)
	msg.SetTag(tag)
	producer.SendAsync(context.TODO(), msg, func(ctx context.Context, resp []*rmq_client.SendReceipt, err error) {
		if err != nil {
			fmt.Printf("SendAsnyc failed\n")
			klog.Error("SendAsnyc failed")
		}
		for i := 0; i < len(resp); i++ {
			fmt.Printf("%#v\n", resp[i])
		}
	})
}

func SendMsgSync(producer rmq_client.Producer, ctx context.Context, message string, topicSuffix string, key string, tag string) error {
	msg := &rmq_client.Message{
		Topic: Topic + topicSuffix,
		Body:  []byte(message),
	}
	msg.SetKeys(key)
	msg.SetTag(tag)
	_, err := producer.Send(ctx, msg)
	if err != nil {
		fmt.Printf("SendAsnyc failed\n")
		klog.Error("SendAsnyc failed")
	}

	return err
}

func ReceiveMsg(simpleConsumer rmq_client.SimpleConsumer, ctx context.Context) {
	for {
		msgs, err := simpleConsumer.Receive(ctx, MaxMessageNum, InvisibleDuration)
		fmt.Printf("Receive message %+v\n", msgs)
		if err != nil {
			klog.Error("Receive msg failed: ", err)
			fmt.Println("Receive msg failed： ", err)
		}

		for _, mv := range msgs {
			if ackErr := simpleConsumer.Ack(ctx, mv); ackErr != nil {
				klog.Error("Ack msg failed: ", ackErr)
			}
			fmt.Println(mv)
		}
	}
}
