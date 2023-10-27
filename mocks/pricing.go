package mocks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/stretchr/testify/mock"
)

type PricingClient struct {
	mock.Mock
}

func (_m *PricingClient) GetProducts(ctx context.Context, input *pricing.GetProductsInput, opt ...func(*pricing.Options)) (*pricing.GetProductsOutput, error) {
	ret := _m.Called(ctx, input, opt)
	if ret.Get(1) == nil {
		return ret.Get(0).(*pricing.GetProductsOutput), nil
	}
	return nil, ret.Get(1).(error)
}
