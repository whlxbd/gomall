package authenticator

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
	"github.com/whlxbd/gomall/common/middleware/authenticator/rpc"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type Payload struct {
	UserId int32
	Type   string
}

func GetToken(ctx context.Context) (token string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", kerrors.NewBizStatusError(400, "metadata not found")
	}
	tokens := md.Get("Authorization")
	if len(tokens) == 0 || len(tokens[0]) < 7 {
		klog.Errorf("tokens: %v", tokens)
		klog.Errorf("ctx: %v", ctx)
		return "", kerrors.NewBizStatusError(400, "token not found")
	}
	token = tokens[0][7:]
	return token, nil
}

func GetPayload(ctx context.Context, token string) (payload *Payload, err error) {
	authVerifyResp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{Token: token})
	if err != nil {
		klog.Errorf("verify token failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}
	if !authVerifyResp.Res {
		klog.Errorf("token invalid")
		return nil, kerrors.NewBizStatusError(401, "token invalid")
	}

	authPayloadResp, err := rpc.AuthClient.GetPayload(ctx, &auth.GetPayloadReq{Token: token})
	if err != nil {
		klog.Errorf("get payload failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}

	payload = &Payload{
		UserId: authPayloadResp.UserId,
		Type:   authPayloadResp.Type,
	}

	serviceName, _ := kitexutil.GetIDLServiceName(ctx)
	methodName, _ := kitexutil.GetMethod(ctx)
	klog.Infof("service: %s, method: %s, payload: %v", serviceName, methodName, payload)
	checkPermissionResp, err := rpc.AuthClient.CheckPermission(ctx, &auth.CheckPermissionReq{
		Role:   payload.Type,
		Router: serviceName + "." + methodName,
	})
	if err != nil {
		klog.Errorf("check permission failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}
	if !checkPermissionResp.Ok {
		klog.Errorf("permission denied")
		return nil, kerrors.NewBizStatusError(403, "permission denied")
	}

	return payload, nil
}

func AuthenticatorMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	rpc.InitClient()
	return func(ctx context.Context, req, resp interface{}) (err error) {
		// Finish your business logic.
		// Get Router
		serviceName, _ := kitexutil.GetIDLServiceName(ctx)
		methodName, _ := kitexutil.GetMethod(ctx)
		router := serviceName + "." + methodName
		// Check White Router
		checkWhiteResp, err := rpc.AuthClient.CheckWhite(ctx, &auth.CheckWhiteReq{
			Router: router,
		})
		if err != nil {
			klog.Errorf("check white failed: %v", err)
		}
		if checkWhiteResp != nil && checkWhiteResp.Ok {
			return next(ctx, req, resp)
		}

		// Get Token and Payload
		token, err := GetToken(ctx)
		if err != nil {
			klog.Errorf("get token failed: %v", err)
			return err
		}
		payload, err := GetPayload(ctx, token)
		if err != nil {
			klog.Errorf("get payload failed: %v", err)
			return err
		}

		// Check Permission
		checkPermissionResp, err := rpc.AuthClient.CheckPermission(ctx, &auth.CheckPermissionReq{
			Role:   payload.Type,
			Router: serviceName + "." + methodName,
		})
		if err != nil {
			klog.Errorf("check permission failed: %v", err)
			return err
		}

		if !checkPermissionResp.Ok {
			klog.Errorf("permission denied")
			return kerrors.NewBizStatusError(403, "permission denied")
		}

		return next(ctx, req, resp)
	}
}
