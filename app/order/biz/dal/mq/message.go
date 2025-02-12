package mq

import (
    "sync"
    rmq_client "github.com/apache/rocketmq-clients/golang/v5"
    "github.com/whlxbd/gomall/common/rmq"
)

var (
    orderMQ  *OrderMQ
    orderOnce sync.Once
)

type OrderMQ struct {
    producer rmq_client.Producer
    consumer rmq_client.SimpleConsumer
}

func Init(endpoint string) error {
    var err error
    orderOnce.Do(func() {
        producer, perr := rmq.InitProducer("order", endpoint)
        if perr != nil {
            err = perr
            return
        }

        consumer, cerr := rmq.InitConsumer("order", "order", endpoint)
        if cerr != nil {
            producer.GracefulStop()
            err = cerr
            return
        }

        orderMQ = &OrderMQ{
            producer: producer,
            consumer: consumer,
        }
    })
    return err
}