package mocks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/stretchr/testify/mock"
)

type KinesisClient struct {
	mock.Mock
}

func (_m *KinesisClient) ListStreamConsumers(ctx context.Context, params *kinesis.ListStreamConsumersInput, optFns ...func(*kinesis.Options)) (*kinesis.ListStreamConsumersOutput, error) {
	ret := _m.Called(ctx, params, optFns)
	if ret.Get(1) == nil {
		return ret.Get(0).(*kinesis.ListStreamConsumersOutput), nil
	}
	return nil, ret.Get(1).(error)
}
