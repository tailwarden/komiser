package kinesis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cloudwatchTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tailwarden/komiser/mocks"
	. "github.com/tailwarden/komiser/models"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func Test_getStreamConsumers(t *testing.T) {
	tests := []struct {
		name       string
		stream     types.StreamSummary
		setupMock  func(m *mocks.KinesisClient)
		clientName string
		region     string
		want       []Resource
		wantErr    bool
	}{
		{
			name: "Should return one EFO consumer",
			stream: types.StreamSummary{
				StreamARN:  aws.String("arn:aws:kinesis:us-east-1:0123456789:stream/kinesis-data-stream"),
				StreamName: aws.String("kinesis-data-stream"),
			},
			setupMock: func(m *mocks.KinesisClient) {
				m.On("ListStreamConsumers", mock.Anything, mock.Anything, mock.Anything).Return(&kinesis.ListStreamConsumersOutput{
					Consumers: []types.Consumer{
						{
							ConsumerARN:               aws.String("arn:aws:kinesis:us-east-1:0123456789:stream/kinesis-data-stream/consumer/kinesis-efo-consumer:1234567890"),
							ConsumerCreationTimestamp: aws.Time(time.UnixMilli(1234567890)),
							ConsumerName:              aws.String("kinesis-efo-consumer"),
							ConsumerStatus:            types.ConsumerStatusActive,
						},
					},
				}, nil).Once()
			},
			clientName: "sandbox",
			region:     "us-east-1",
			want: []Resource{
				{
					Provider:   "AWS",
					Account:    "sandbox",
					Service:    "Kinesis EFO Consumer",
					ResourceId: "arn:aws:kinesis:us-east-1:0123456789:stream/kinesis-data-stream/consumer/kinesis-efo-consumer:1234567890",
					Region:     "us-east-1",
					Name:       "kinesis-efo-consumer",
					Cost:       0,
					CreatedAt:  time.UnixMilli(1234567890),
					FetchedAt:  time.Now(),
					Link:       "https://us-east-1.console.aws.amazon.com/kinesis/home?region=us-east-1#/streams/details/kinesis-data-stream/registeredConsumers/kinesis-efo-consumer",
				},
			},
			wantErr: false,
		},
		{
			name: "Should paginate using next token",
			stream: types.StreamSummary{
				StreamARN:  aws.String("arn:aws:kinesis:us-east-1:0123456789:stream/kinesis-data-stream"),
				StreamName: aws.String("kinesis-data-stream"),
			},
			setupMock: func(m *mocks.KinesisClient) {
				m.On("ListStreamConsumers", mock.Anything, mock.Anything, mock.Anything).Return(&kinesis.ListStreamConsumersOutput{
					NextToken: aws.String("next-token"),
					Consumers: []types.Consumer{
						{
							ConsumerARN:               aws.String("arn:aws:kinesis:us-east-1:0123456789:stream/kinesis-data-stream/consumer/kinesis-efo-consumer-1:1234567890"),
							ConsumerCreationTimestamp: aws.Time(time.UnixMilli(1234567890)),
							ConsumerName:              aws.String("kinesis-efo-consumer-1"),
							ConsumerStatus:            types.ConsumerStatusActive,
						},
					},
				}, nil).Once()
				m.On("ListStreamConsumers", mock.Anything, mock.Anything, mock.Anything).Return(&kinesis.ListStreamConsumersOutput{
					Consumers: []types.Consumer{
						{
							ConsumerARN:               aws.String("arn:aws:kinesis:us-east-1:0123456789:stream/kinesis-data-stream/consumer/kinesis-efo-consumer-2:1234567890"),
							ConsumerCreationTimestamp: aws.Time(time.UnixMilli(1234567890)),
							ConsumerName:              aws.String("kinesis-efo-consumer-2"),
							ConsumerStatus:            types.ConsumerStatusActive,
						},
					},
				}, nil).Once()
			},
			clientName: "sandbox",
			region:     "us-east-1",
			want: []Resource{
				{
					Provider:   "AWS",
					Account:    "sandbox",
					Service:    "Kinesis EFO Consumer",
					ResourceId: "arn:aws:kinesis:us-east-1:0123456789:stream/kinesis-data-stream/consumer/kinesis-efo-consumer-1:1234567890",
					Region:     "us-east-1",
					Name:       "kinesis-efo-consumer-1",
					Cost:       0,
					CreatedAt:  time.UnixMilli(1234567890),
					FetchedAt:  time.Now(),
					Link:       "https://us-east-1.console.aws.amazon.com/kinesis/home?region=us-east-1#/streams/details/kinesis-data-stream/registeredConsumers/kinesis-efo-consumer-1",
				},
				{
					Provider:   "AWS",
					Account:    "sandbox",
					Service:    "Kinesis EFO Consumer",
					ResourceId: "arn:aws:kinesis:us-east-1:0123456789:stream/kinesis-data-stream/consumer/kinesis-efo-consumer-2:1234567890",
					Region:     "us-east-1",
					Name:       "kinesis-efo-consumer-2",
					Cost:       0,
					CreatedAt:  time.UnixMilli(1234567890),
					FetchedAt:  time.Now(),
					Link:       "https://us-east-1.console.aws.amazon.com/kinesis/home?region=us-east-1#/streams/details/kinesis-data-stream/registeredConsumers/kinesis-efo-consumer-2",
				},
			},
			wantErr: false,
		},
		{
			name: "Should return error if error with kinesis client",
			stream: types.StreamSummary{
				StreamARN:  aws.String("arn:aws:kinesis:us-east-1:0123456789:stream/kinesis-data-stream"),
				StreamName: aws.String("kinesis-data-stream"),
			},
			setupMock: func(m *mocks.KinesisClient) {
				m.On("ListStreamConsumers", mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("unit test error")).Once()
			},
			clientName: "sandbox",
			region:     "us-east-1",
			want:       []Resource{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			kinesisClient := &mocks.KinesisClient{}
			tt.setupMock(kinesisClient)

			got, err := getStreamConsumers(ctx, kinesisClient, tt.stream, tt.clientName, tt.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("getStreamConsumers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("getStreamConsumers() incorrect lenght of resources got = %v, want %v", len(got), len(tt.want))
			} else {
				for i := range got {
					assert.Equalf(t, tt.want[i].Link, got[i].Link, "incorrect Link for resources")
					assert.Equalf(t, tt.want[i].Provider, got[i].Provider, "incorrect Provider for resources")
					assert.Equalf(t, tt.want[i].Account, got[i].Account, "incorrect Account for resources")
					assert.Equalf(t, tt.want[i].Service, got[i].Service, "incorrect Service for resources")
					assert.Equalf(t, tt.want[i].ResourceId, got[i].ResourceId, "incorrect ResourceId for resources")
					assert.Equalf(t, tt.want[i].Region, got[i].Region, "incorrect Region for resources")
					assert.Equalf(t, tt.want[i].Name, got[i].Name, "incorrect Name for resources")
					assert.Equalf(t, tt.want[i].Cost, got[i].Cost, "incorrect Cost for resources")
				}
			}
			kinesisClient.AssertExpectations(t)
		})
	}
}

func TestCalculateCostOfKinesisDataStream(t *testing.T) {
	priceMap := map[string][]awsUtils.PriceDimensions{
		"Provisioned shard hour": {
			{
				EndRange:   "Inf",
				BeginRange: 0.0,
				PricePerUnit: struct {
					USD float64 `json:"USD,string"`
				}{USD: 0.03},
			},
		},
		"Payload Units": {
			{
				EndRange:   "Inf",
				BeginRange: 0.0,
				PricePerUnit: struct {
					USD float64 `json:"USD,string"`
				}{USD: 0.000000028},
			},
		},
	}
	tests := []struct {
		name            string
		stream          types.StreamDescriptionSummary
		totalPutRecords float64
		want            float64
		wantErr         assert.ErrorAssertionFunc
	}{
		{
			name: "Should return zero given on demand mode",
			stream: types.StreamDescriptionSummary{
				StreamStatus: types.StreamStatusActive,
				StreamModeDetails: &types.StreamModeDetails{
					StreamMode: types.StreamModeOnDemand,
				},
			},
			totalPutRecords: 0.0,
			want:            0.0,
			wantErr:         assert.NoError,
		},
		{
			name: "Should return shard cost given provisioned mode",
			stream: types.StreamDescriptionSummary{
				StreamStatus: types.StreamStatusActive,
				StreamModeDetails: &types.StreamModeDetails{
					StreamMode: types.StreamModeProvisioned,
				},
				OpenShardCount:          aws.Int32(4),
				StreamCreationTimestamp: aws.Time(time.Now().Add(-2 * time.Hour)),
			},
			totalPutRecords: 0.0,
			want:            0.24,
			wantErr:         assert.NoError,
		},
		{
			name: "Should return put record cost given provisioned mode",
			stream: types.StreamDescriptionSummary{
				StreamStatus: types.StreamStatusActive,
				StreamModeDetails: &types.StreamModeDetails{
					StreamMode: types.StreamModeProvisioned,
				},
				OpenShardCount:          aws.Int32(0),
				StreamCreationTimestamp: aws.Time(time.Now().Add(-2 * time.Hour)),
			},
			totalPutRecords: 50000000.0, // 50 million
			want:            1.4,
			wantErr:         assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateCostOfKinesisDataStream(&tt.stream, tt.totalPutRecords, priceMap)
			if !tt.wantErr(t, err, fmt.Sprintf("calculateCostOfKinesisDataStream(%v)", tt.stream)) {
				return
			}
			assert.Equalf(t, tt.want, got, "calculateCostOfKinesisDataStream(%v)", tt.stream)
		})
	}
}

func TestRetrievePriceMap(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocks.PricingClient)
		want      map[string][]awsUtils.PriceDimensions
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "Should return error giving error with pricing client",
			setupMock: func(m *mocks.PricingClient) {
				m.On("GetProducts", mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("unit test error")).Once()
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "Should return error given invalid json in pricing output",
			setupMock: func(m *mocks.PricingClient) {
				m.On("GetProducts", mock.Anything, mock.Anything, mock.Anything).Return(&pricing.GetProductsOutput{
					PriceList: []string{
						"invalid-json",
					},
				}, nil).Once()
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "Should return price map",
			setupMock: func(m *mocks.PricingClient) {
				m.On("GetProducts", mock.Anything, mock.Anything, mock.Anything).Return(&pricing.GetProductsOutput{
					PriceList: []string{
						`{
							"product": {
								"attributes": {
									"group": "Kinesis Data Streams",
									"operation": "Write Throughput Units",
									"groupDescription": "Kinesis Data Streams",
									"requestDescription": "Write Throughput Units (PUT records)",
									"instanceType": "Write Throughput Units",
									"instanceTypeFamily": "Write Throughput Units"
								}
							},
							"terms": {
								"OnDemand": {
									"1234567890": {
										"priceDimensions": {
											"1234567890": {
												"unit": "Write Throughput Unit-Hours",
												"pricePerUnit": {
													"USD": "0.014"
												},
												"appliesTo": []
											}
										},
										"sku": "1234567890",
										"effectiveDate": "2021-01-01T00:00:00Z",
										"offerTermCode": "1234567890",
										"termAttributes": {}
									}
								}
							}
						}`,
					},
				}, nil).Once()
			},
			want: map[string][]awsUtils.PriceDimensions{
				"Kinesis Data Streams": {
					{
						EndRange:   "",
						BeginRange: 0.0,
						PricePerUnit: struct {
							USD float64 `json:"USD,string"`
						}{USD: 0.014},
					},
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			pricingClient := &mocks.PricingClient{}
			tt.setupMock(pricingClient)

			got, err := retrievePriceMap(ctx, pricingClient, "us-east-1")
			if !tt.wantErr(t, err, fmt.Sprintf("retrievePriceMap(%v)", tt.name)) {
				return
			}
			assert.Equalf(t, tt.want, got, "retrievePriceMap(%v)", tt.name)
		})
	}
}

func TestRetrievePutRecords(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		setupMock func(m *mocks.CloudwatchClient)
		want      float64
	}{
		{
			name: "Should return zero given error with cloudwatch client",
			setupMock: func(m *mocks.CloudwatchClient) {
				m.On("GetMetricStatistics", ctx, mock.MatchedBy(func(input *cloudwatch.GetMetricStatisticsInput) bool {
					return *input.Namespace == "AWS/Kinesis" && *input.MetricName == "PutRecords.SuccessfulRecords" && input.Statistics[0] == cloudwatchTypes.StatisticSum && *input.Dimensions[0].Value == "test-data-stream"
				}), mock.Anything).Return(nil, fmt.Errorf("unit test error")).Once()
			},
			want: 0.0,
		},
		{
			name: "Should return zero given no datapoints",
			setupMock: func(m *mocks.CloudwatchClient) {
				m.On("GetMetricStatistics", ctx, mock.MatchedBy(func(input *cloudwatch.GetMetricStatisticsInput) bool {
					return *input.Namespace == "AWS/Kinesis" && *input.MetricName == "PutRecords.SuccessfulRecords" && input.Statistics[0] == cloudwatchTypes.StatisticSum && *input.Dimensions[0].Value == "test-data-stream"
				}), mock.Anything).Return(&cloudwatch.GetMetricStatisticsOutput{
					Datapoints: []cloudwatchTypes.Datapoint{},
				}, nil).Once()
			},
			want: 0.0,
		},
		{
			name: "Should return 1 million",
			setupMock: func(m *mocks.CloudwatchClient) {
				m.On("GetMetricStatistics", ctx, mock.MatchedBy(func(input *cloudwatch.GetMetricStatisticsInput) bool {
					return *input.Namespace == "AWS/Kinesis" && *input.MetricName == "PutRecords.SuccessfulRecords" && input.Statistics[0] == cloudwatchTypes.StatisticSum && *input.Dimensions[0].Value == "test-data-stream"
				}), mock.Anything).Return(&cloudwatch.GetMetricStatisticsOutput{
					Datapoints: []cloudwatchTypes.Datapoint{
						{
							Sum: aws.Float64(1000000.0),
						},
					},
				}, nil).Once()
			},
			want: 1000000.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			streamName := aws.String("test-data-stream")

			cloudwatchClient := &mocks.CloudwatchClient{}
			tt.setupMock(cloudwatchClient)

			assert.Equalf(t, tt.want, retrievePutRecords(ctx, cloudwatchClient, streamName), "retrievePutRecords(%v)", tt.name)
		})
	}
}
