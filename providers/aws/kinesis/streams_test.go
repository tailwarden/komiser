package kinesis

import (
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tailwarden/komiser/mocks"
	. "github.com/tailwarden/komiser/models"
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
			kinesisClient := &mocks.KinesisClient{}
			tt.setupMock(kinesisClient)

			got, err := getStreamConsumers(kinesisClient, tt.stream, tt.clientName, tt.region)
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
