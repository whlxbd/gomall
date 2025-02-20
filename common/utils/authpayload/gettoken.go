package authpayload

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/whlxbd/gomall/common/utils/authpayload/rpc"
)

func Token(ctx context.Context) (token string, err error) {
	rpc.InitClient()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", kerrors.NewBizStatusError(400, "metadata not found")
	}
	tokens := md.Get("Authorization")
	if len(tokens) == 0 || len(tokens[0]) < 7 {
		return "", kerrors.NewBizStatusError(400, "token not found")
	}
	token = tokens[0][7:]
	return token, nil
}
