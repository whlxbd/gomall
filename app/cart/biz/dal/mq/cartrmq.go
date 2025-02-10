package mq

import (
	"context"
	"encoding/json"
	"sync"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/whlxbd/gomall/app/cart/biz/dal/model"
	"github.com/whlxbd/gomall/common/rmq"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"gorm.io/gorm"
)

type CartMessage struct {
	Operation string  `json:"operation"`
	UserID    uint32  `json:"user_id"`
	CartList  []*model.Cart `json:"cart"`
}

const (
    OperationAdd    = "add"
    OperationUpdate = "update"
    OperationDelete = "delete"
)

// CartMQ 单例管理器
type CartMQ struct {
    producer rmq_client.Producer
    consumer rmq_client.SimpleConsumer
    db       *gorm.DB
    redis    *redis.Client
}

var (
    cartMQ  *CartMQ
    cartMQOnce sync.Once
)

// InitCartMQ 单例初始化
func Init(db *gorm.DB, redis *redis.Client) error {
    os.Setenv("mq.consoleAppender.enabled", "true")
	rmq_client.ResetLogger()
    var err error
    cartMQOnce.Do(func() {
        producer, perr := rmq.InitProducer("cart", os.Getenv("RMQENDPOINT"))
        if perr != nil {
            err = perr
            return
        }

        consumer, cerr := rmq.InitConsumer("cart", "cart_consumer", os.Getenv("RMQENDPOINT"))
        if cerr != nil {
            err = cerr
            return
        }

        cartMQ = &CartMQ{
            producer: producer,
            consumer: consumer,
            db:       db,
            redis:    redis,
        }
    })
    return err
}

// SendCartMessage 发送购物车消息
func (mq *CartMQ) SendCartMessage(ctx context.Context, msg *CartMessage) error {
    data, err := json.Marshal(msg)
    if err != nil {
        return err
    }
    return rmq.SendMsgSync(mq.producer, ctx, string(data), "cart", msg.Operation, "cart")
}

// StartConsumer 启动消费者
func (mq *CartMQ) StartConsumer(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                msgs, err := mq.consumer.Receive(ctx, rmq.MaxMessageNum, rmq.InvisibleDuration)
                if err != nil {
                    continue
                }

                for _, msg := range msgs {
                    var cartMsg CartMessage
                    if err := json.Unmarshal(msg.GetBody(), &cartMsg); err != nil {
                        continue
                    }

                    // 处理消息
                    if err := mq.handleMessage(ctx, &cartMsg); err != nil {
                        continue
                    }

                    mq.consumer.Ack(ctx, msg)
                }
            }
        }
    }()
}

// handleMessage 处理消息
func (mq *CartMQ) handleMessage(ctx context.Context, msg *CartMessage) error {
    // 1. 更新数据库
    
	fmt.Printf("handleMessage %+v\n", msg)


    return nil
}

// GetCartMQ 获取单例
func GetCartMQ() *CartMQ {
    return cartMQ
}

// Close 关闭连接
func (mq *CartMQ) Close() {
    if mq == nil {
        return
    }

    if mq.producer != nil {
        mq.producer.GracefulStop()
    }
    if mq.consumer != nil {
        mq.consumer.GracefulStop()
    }
}