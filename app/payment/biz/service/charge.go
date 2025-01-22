package service

import (
	"context"
	payment "github.com/whlxbd/gomall/rpc_gen/kitex_gen/payment"
	creditcard "github.com/durango/go-credit-card"
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
		Cvv: string(req.CreditCard.CreditCardCvv),
		Month: string(req.CreditCard.CreditCardExpirationMonth),
		Year: string(req.CreditCard.CreditCardExpirationYear),
	}

	err = cardInfo.Validate()
	if err != nil {
		return nil, err
	}

	return
}
