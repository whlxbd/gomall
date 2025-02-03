package rmq

import (
    "context"
    "testing"
    "time"
    
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
            endpoint:    "10.255.253.63:9878",
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
    client, err := Init("test", "test", "10.255.253.63:9878")
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
    client, err := Init("test", "test", "10.255.253.63:9878")
    assert.NoError(t, err)
    defer client.Close()

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
    defer cancel()

    done := make(chan struct{})
    go func() {
        client.ReceiveMsg(ctx)
        close(done)
    }()

    select {
    case <-done:
        // 正常退出
    case <-time.After(time.Second * 3):
        t.Error("接收消息超时")
    }
}