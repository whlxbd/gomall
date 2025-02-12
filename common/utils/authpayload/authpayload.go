// Get retrieves the authentication payload from the context.
// It extracts the token from the metadata in the incoming context,
// verifies the token using the AuthClient, and then fetches the payload.
//
// Parameters:
//   - ctx: The context containing the metadata with the token.
//
// Returns:
//   - payload: The authentication payload response.
//   - err: An error if the metadata is not found, the token is not found or invalid,
//          or if there is an error during the RPC calls to verify the token or get the payload.

package authpayload

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/whlxbd/gomall/common/utils/authpayload/rpc"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

// 传入微服务的上下文，该函数将从metadata中提取token并校验token的有效性，最后获取token的payload并返回
func Get(ctx context.Context) (payload *auth.GetPayloadResp, err error) {
	rpc.InitClient()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, kerrors.NewBizStatusError(400, "metadata not found")
	}
	tokens := md.Get("token")
	if len(tokens) == 0 || tokens[0] == "" {
		return nil, kerrors.NewBizStatusError(400, "token not found")
	}
	token := tokens[0]

	authVerifyResp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{Token: token})
	if err != nil {
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}
	if !authVerifyResp.Res {
		return nil, kerrors.NewBizStatusError(401, "token invalid")
	}

	payload, err = rpc.AuthClient.GetPayload(ctx, &auth.GetPayloadReq{Token: token})
	if err != nil {
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}

	return
}
