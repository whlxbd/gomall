package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/whlxbd/gomall/app/payment/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/payment/biz/model"
	payment "github.com/whlxbd/gomall/rpc_gen/kitex_gen/payment"
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

	err = model.Create(mysql.DB, s.ctx, &model.PaymentRecord{
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		OrderId:       req.OrderId,
		UserId:        req.UserId,
		PayAt:         time.Now(),
	})
	if err != nil {
		klog.Errorf("create payment record failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}

	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}
