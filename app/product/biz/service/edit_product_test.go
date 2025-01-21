package service

import (
	"context"
	"testing"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
)

func TestEditProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEditProductService(ctx)
	// init req and assert value

	req := &product.EditProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
