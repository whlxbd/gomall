package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/whlxbd/gomall/app/payment/biz/dal/model"
	"github.com/whlxbd/gomall/app/payment/biz/dal/mysql"
	"github.com/whlxbd/gomall/common/utils/authpayload"
	payment "github.com/whlxbd/gomall/rpc_gen/kitex_gen/payment"
	"gorm.io/gorm"

	"github.com/whlxbd/gomall/app/payment/infra/rpc"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/order"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	payload, err := authpayload.Get(s.ctx)
	if err != nil {
		klog.Errorf("get payload failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "get payload failed")
	}
	if payload.UserId != int32(req.UserId) {
		klog.Errorf("permission denied, user id: %v, req user id: %v", payload.UserId, req.UserId)
		return nil, kerrors.NewBizStatusError(400, "permission denied")
	}
	cardInfo := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	err = cardInfo.Validate(true)
	if err != nil {
		klog.Errorf("validate credit card failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, err.Error())
	}

	transactionId, err := uuid.NewRandom()
	if err != nil {
		klog.Errorf("generate transaction id failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		err = model.Create(mysql.DB, s.ctx, &model.PaymentRecord{
			TransactionId: transactionId.String(),
			Amount:        req.Amount,
			OrderId:       req.OrderId,
			UserId:        req.UserId,
			PayAt:         time.Now(),
		})
		if err != nil {
			klog.Errorf("create payment record failed: %v", err)
			return kerrors.NewBizStatusError(500, err.Error())
		}

		_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{
			OrderId: req.OrderId,
			UserId:  req.UserId,
		})
		if err != nil {
			klog.Errorf("mark order paid failed: %v", err)
			return kerrors.NewBizStatusError(500, err.Error())
		}

		return nil
	})
	if err != nil {
		klog.Errorf("transaction failed: %v", err)
		return nil, err
	}

	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}
