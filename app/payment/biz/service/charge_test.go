package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	payment "github.com/whlxbd/gomall/rpc_gen/kitex_gen/payment"
)

func TestCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeService(ctx)

	req := &payment.ChargeReq{
		OrderId: "1",
		UserId:  1,
		Amount:  1,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "424242424242424242",
			CreditCardCvv:             123,
			CreditCardExpirationMonth: 12,
			CreditCardExpirationYear:  2030,
		},
	}
	resp, err := s.Run(req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.TransactionId)
}
