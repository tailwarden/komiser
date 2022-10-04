package aws

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go/aws"
	. "github.com/mlabouardy/komiser/models/aws"
)

type LambdaFunction struct {
}

func getRuntime(input string) string {
	if strings.HasPrefix(input, "go") {
		return "golang"
	}
	if strings.HasPrefix(input, "java") {
		return "java"
	}
	if strings.HasPrefix(input, "python") {
		return "python"
	}
	if strings.HasPrefix(input, "node") {
		return "node"
	}
	return "custom"
}

func (aws AWS) DescribeLambdaFunctions(cfg awsConfig.Config) ([]Lambda, error) {
	output := make([]Lambda, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return output, err
	}
	for _, region := range regions {
		functions, err := aws.getLambdaFunctions(cfg, region.Name)
		if err != nil {
			return output, err
		}

		for _, f := range functions {
			output = append(output, f)
		}
	}
	return output, nil
}

type LambdaTotalInvocationMetric struct {
	Timestamp time.Time       `json:"timestamp"`
	Metrics   []MetricKeyPair `json:"metrics"`
}

type MetricKeyPair struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

func (awsModel AWS) GetLambdaInvocationMetrics(cfg awsConfig.Config) ([]LambdaTotalInvocationMetric, error) {
	datapoints := make([]LambdaInvocationMetric, 0)

	regions, err := awsModel.getRegions(cfg)
	if err != nil {
		return []LambdaTotalInvocationMetric{}, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := cloudwatch.NewFromConfig(cfg)
		result, err := svc.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/Lambda"),
			MetricName: aws.String("Invocations"),
			StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int32(86400),
			Statistics: []types.Statistic{
				types.StatisticSum,
			},
		})
		if err != nil {
			return []LambdaTotalInvocationMetric{}, err
		}
		metrics := make([]Datapoint, 0)
		for _, datapoint := range result.Datapoints {
			metrics = append(metrics, Datapoint{
				Value:     *datapoint.Sum,
				Timestamp: *datapoint.Timestamp,
			})
		}

		datapoints = append(datapoints, LambdaInvocationMetric{
			Region:     region.Name,
			Datapoints: metrics,
		})
	}

	metrics := make(map[string]map[string]float64, 0)
	for _, datapoint := range datapoints {
		for _, dt := range datapoint.Datapoints {
			if metrics[dt.Timestamp.Format("2006-01")] == nil {
				metrics[dt.Timestamp.Format("2006-01")] = make(map[string]float64, 1)
				metrics[dt.Timestamp.Format("2006-01")][datapoint.Region] = dt.Value
			} else {
				metrics[dt.Timestamp.Format("2006-01")][datapoint.Region] += dt.Value
			}
		}
	}

	output := make([]LambdaTotalInvocationMetric, 0)
	for timestamp, metric := range metrics {
		time, _ := time.Parse("2006-01", timestamp)
		series := make([]MetricKeyPair, 0)
		for region, value := range metric {
			series = append(series, MetricKeyPair{
				Label: region,
				Value: value,
			})
		}
		output = append(output, LambdaTotalInvocationMetric{
			Timestamp: time,
			Metrics:   series,
		})
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Timestamp.Sub(output[j].Timestamp) < 0
	})

	return output, err
}

func (awsModel AWS) GetLambdaErrorsMetrics(cfg awsConfig.Config) ([]LambdaTotalInvocationMetric, error) {
	datapoints := make([]LambdaInvocationMetric, 0)

	regions, err := awsModel.getRegions(cfg)
	if err != nil {
		return []LambdaTotalInvocationMetric{}, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := cloudwatch.NewFromConfig(cfg)
		result, err := svc.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/Lambda"),
			MetricName: aws.String("Errors"),
			StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int32(86400),
			Statistics: []types.Statistic{
				types.StatisticSum,
			},
		})
		if err != nil {
			return []LambdaTotalInvocationMetric{}, err
		}
		metrics := make([]Datapoint, 0)
		for _, datapoint := range result.Datapoints {
			metrics = append(metrics, Datapoint{
				Value:     *datapoint.Sum,
				Timestamp: *datapoint.Timestamp,
			})
		}

		datapoints = append(datapoints, LambdaInvocationMetric{
			Region:     region.Name,
			Datapoints: metrics,
		})
	}

	metrics := make(map[string]map[string]float64, 0)
	for _, datapoint := range datapoints {
		for _, dt := range datapoint.Datapoints {
			if metrics[dt.Timestamp.Format("2006-01-02")] == nil {
				metrics[dt.Timestamp.Format("2006-01-02")] = make(map[string]float64, 1)
				metrics[dt.Timestamp.Format("2006-01-02")][datapoint.Region] = dt.Value
			} else {
				metrics[dt.Timestamp.Format("2006-01-02")][datapoint.Region] += dt.Value
			}
		}
	}

	output := make([]LambdaTotalInvocationMetric, 0)
	for timestamp, metric := range metrics {
		time, _ := time.Parse("2006-01-02", timestamp)
		series := make([]MetricKeyPair, 0)
		for region, value := range metric {
			series = append(series, MetricKeyPair{
				Label: region,
				Value: value,
			})
		}
		output = append(output, LambdaTotalInvocationMetric{
			Timestamp: time,
			Metrics:   series,
		})
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Timestamp.Sub(output[j].Timestamp) < 0
	})

	return output, err
}

func (aws AWS) getLambdaFunctions(cfg awsConfig.Config, region string) ([]Lambda, error) {
	cfg.Region = region
	svc := lambda.NewFromConfig(cfg)
	params := &lambda.ListFunctionsInput{}
	listOfFunctions := make([]Lambda, 0)
	for {
		result, err := svc.ListFunctions(context.Background(), params)
		if err != nil {
			return []Lambda{}, err
		}
		for _, l := range result.Functions {
			tagsResp, err := svc.ListTags(context.Background(), &lambda.ListTagsInput{
				Resource: *&l.FunctionArn,
			})
			if err != nil {
				return []Lambda{}, err
			}

			tags := make([]string, 0)
			for key, value := range tagsResp.Tags {
				tags = append(tags, fmt.Sprintf("%s:%s", key, value))
			}

			listOfFunctions = append(listOfFunctions, Lambda{
				Name:        *l.FunctionName,
				Memory:      int64(*l.MemorySize),
				Runtime:     string(l.Runtime),
				FunctionArn: *l.FunctionArn,
				Tags:        tags,
				Region:      region,
			})
		}
		if result.NextMarker == nil {
			break
		}
		params = &lambda.ListFunctionsInput{
			Marker: result.NextMarker,
		}
	}
	return listOfFunctions, nil
}
