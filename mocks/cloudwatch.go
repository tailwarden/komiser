package mocks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/stretchr/testify/mock"
)

type CloudwatchClient struct {
	mock.Mock
}

func (_m *CloudwatchClient) GetMetricStatistics(ctx context.Context, input *cloudwatch.GetMetricStatisticsInput, opt ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricStatisticsOutput, error) {
	ret := _m.Called(ctx, input, opt)
	if ret.Get(1) == nil {
		return ret.Get(0).(*cloudwatch.GetMetricStatisticsOutput), nil
	}
	return nil, ret.Get(1).(error)
}
