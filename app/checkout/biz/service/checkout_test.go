package service

import (
	"context"
	"testing"

	checkout "github.com/whlxbd/gomall/rpc_gen/kitex_gen/checkout"
	"google.golang.org/grpc/metadata"
)

func TestCheckout_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckoutService(ctx)
	// init req and assert value

	req := &checkout.CheckoutReq{
		UserId: 1,
	}

	// todo: edit your unit test

	md := metadata.New(map[string]string{
		"authorization": "Bearer your_token_here",
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := s.Run(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatalf("expected response, got nil")
	}
	t.Logf("response: %v", resp)
}
