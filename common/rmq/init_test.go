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

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
    defer cancel()

    done := make(chan struct{})
    msgSent := make(chan struct{})
    go func() {
        client.ReceiveMsg(ctx)
        close(done)
    }()
    
    go func() {
        // 限制发送5条消息
        for i := 0; i < 5; i++ {
            err := client.SendMsgSync(ctx, fmt.Sprintf("test message %d", i), "test", "test_key", "test_tag")
            if err != nil {
                t.Errorf("发送消息失败: %v", err)
                return
            }
            time.Sleep(time.Second) // 间隔发送
        }
        close(msgSent)
    }()

    select {
    case <-done:
        // 正常退出
    case <-time.After(time.Second * 3):
        t.Error("接收消息超时")
    }
}