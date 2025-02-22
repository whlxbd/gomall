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
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/common/utils/authpayload/rpc"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"

	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
)

// 传入微服务的上下文，该函数将从metadata中提取token并校验token的有效性，最后获取token的payload并返回
func Get(ctx context.Context) (payload *auth.GetPayloadResp, err error) {
	token, err := Token(ctx)
	if err != nil {
		klog.Errorf("get token failed: %v", err)
		return nil, err
	}

	authVerifyResp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{Token: token})
	if err != nil {
		klog.Errorf("verify token failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}
	if !authVerifyResp.Res {
		klog.Errorf("token invalid")
		return nil, kerrors.NewBizStatusError(401, "token invalid")
	}

	payload, err = rpc.AuthClient.GetPayload(ctx, &auth.GetPayloadReq{Token: token})
	if err != nil {
		klog.Errorf("get payload failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}

	serviceName, _ := kitexutil.GetIDLServiceName(ctx)
	methodName, _ := kitexutil.GetMethod(ctx)
	klog.Infof("service: %s, method: %s, payload: %v", serviceName, methodName, payload)
	checkPermissionResp, err := rpc.RuleClient.CheckPermission(ctx, &rule.CheckPermissionReq{
		Role:   payload.Type,
		Router: serviceName + "/" + methodName,
	})
	if err != nil {
		klog.Errorf("check permission failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}
	if !checkPermissionResp.Ok {
		klog.Errorf("permission denied")
		return nil, kerrors.NewBizStatusError(403, "permission denied")
	}

	return
}
