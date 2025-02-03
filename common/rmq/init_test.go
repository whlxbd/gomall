package rmq

import (
    "context"
    "testing"
    "time"
    "fmt"
    
    "github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
    tests := []struct {
        name        string
        topicSuffix string
        groupSuffix string
        endpoint    string
        wantErr     bool
    }{
        {
            name:        "正常初始化",
            topicSuffix: "test",
            groupSuffix: "group",
            endpoint:    "127.0.0.1:9878",
            wantErr:     false,
        },
        {
            name:        "错误endpoint",
            topicSuffix: "topic",
            groupSuffix: "group",
            endpoint:    "invalid-endpoint",
            wantErr:     true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            client, err := Init(tt.topicSuffix, tt.groupSuffix, tt.endpoint)
			if tt.wantErr {
                assert.Error(t, err)
                return
            }
            assert.NoError(t, err)
            assert.NotNil(t, client)
            defer client.Close()
        })
    }
}

func TestSendMessage(t *testing.T) {
    client, err := Init("test", "test", "127.0.0.1:9878")
    assert.NoError(t, err)
    defer client.Close()

    ctx := context.Background()
    
    // 测试同步发送
    err = client.SendMsgSync(ctx, "test message", "test", "test_key", "test_tag")
    assert.NoError(t, err)

    // 测试异步发送
    client.SendMsgAsync(ctx, "test message async", "test", "test_key", "test_tag")
    time.Sleep(time.Second) // 等待异步完成
}

func TestReceiveMessage(t *testing.T) {
    client, err := Init("test", "test", "127.0.0.1:9878")
    assert.NoError(t, err)
    defer client.Close()

    ctx := context.Background()

    receivedMsgs := make(chan string, 5)
    
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                msgs, err := client.simpleConsumer.Receive(ctx, maxMessageNum, invisibleDuration)
                if err != nil {
                    continue
                }
                
                for _, msg := range msgs {
                    // 确认消息
                    _ = client.simpleConsumer.Ack(ctx, msg)
                    receivedMsgs <- string(msg.GetBody())
                }
            }
        }
    }()

    // 发送测试消息
    for i := 0; i < 5; i++ {
        msgContent := fmt.Sprintf("test message %d", i)
        err := client.SendMsgSync(ctx, msgContent, "test", fmt.Sprintf("key_%d", i), "test_tag")
        assert.NoError(t, err)
        time.Sleep(time.Second)
    }

    // 验证接收到的消息
    receivedCount := 0
    timeout := time.After(time.Second * 30)
    for receivedCount < 5 {
        select {
        case msg := <-receivedMsgs:
            t.Logf("收到消息: %s", msg)
            receivedCount++
        case <-timeout:
            t.Errorf("接收超时, 已接收 %d 条消息", receivedCount)
            return
        }
    }
    t.Logf("成功接收 %d 条消息", receivedCount)
}