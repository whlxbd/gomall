package mq

import (
    "context"
    "encoding/json"
    "github.com/whlxbd/gomall/app/order/biz/dal/model"
    "github.com/whlxbd/gomall/app/order/biz/dal/mysql"
	"github.com/whlxbd/gomall/common/rmq"
)

func StartOrderConsumer(ctx context.Context) {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                msgs, err := orderMQ.consumer.Receive(ctx, rmq.MaxMessageNum, rmq.InvisibleDuration)
                if err != nil {
                    continue
                }

                for _, mv := range msgs {
                    var msg OrderMessage
                    if err := json.Unmarshal(mv.GetBody(), &msg); err != nil {
                        continue
                    }

                    // 处理超时订单
                    if err := handleTimeoutOrder(ctx, &msg); err == nil {
                        orderMQ.consumer.Ack(ctx, mv)
                    }
                }
            }
        }
}

func handleTimeoutOrder(ctx context.Context, msg *OrderMessage) error {
    // 查询当前订单状态
    order, err := model.GetOrder(mysql.DB, ctx, msg.UserId, msg.OrderId)
    if err != nil {
        return err
    }

    // 只取消未支付订单
    if order.OrderState != model.OrderStatePlaced {
        return nil
    }

    // 更新为已取消状态
    return model.UpdateOrderState(mysql.DB, ctx, msg.UserId, msg.OrderId, model.OrderStateCanceled)
}