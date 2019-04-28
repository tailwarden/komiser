package aws

import (
	"context"
	"sort"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	. "github.com/mlabouardy/komiser/models/aws"
)

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

func (aws AWS) DescribeLambdaFunctions(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		functions, err := aws.getLambdaFunctions(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		for _, lambda := range functions {
			output[getRuntime(lambda.Runtime)]++
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

func (awsModel AWS) GetLambdaInvocationMetrics(cfg aws.Config) ([]LambdaTotalInvocationMetric, error) {
	datapoints := make([]LambdaInvocationMetric, 0)

	regions, err := awsModel.getRegions(cfg)
	if err != nil {
		return []LambdaTotalInvocationMetric{}, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := cloudwatch.New(cfg)
		req := svc.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/Lambda"),
			MetricName: aws.String("Invocations"),
			StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int64(86400),
			Statistics: []cloudwatch.Statistic{
				cloudwatch.StatisticSum,
			},
		})
		result, err := req.Send(context.Background())
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

func (awsModel AWS) GetLambdaErrorsMetrics(cfg aws.Config) ([]LambdaTotalInvocationMetric, error) {
	datapoints := make([]LambdaInvocationMetric, 0)

	regions, err := awsModel.getRegions(cfg)
	if err != nil {
		return []LambdaTotalInvocationMetric{}, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := cloudwatch.New(cfg)
		req := svc.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/Lambda"),
			MetricName: aws.String("Errors"),
			StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int64(86400),
			Statistics: []cloudwatch.Statistic{
				cloudwatch.StatisticSum,
			},
		})
		result, err := req.Send(context.Background())
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

func (aws AWS) getLambdaFunctions(cfg aws.Config, region string) ([]Lambda, error) {
	cfg.Region = region
	svc := lambda.New(cfg)
	req := svc.ListFunctionsRequest(&lambda.ListFunctionsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []Lambda{}, err
	}
	listOfFunctions := make([]Lambda, 0)
	for _, lambda := range result.Functions {
		runtime, _ := lambda.Runtime.MarshalValue()
		listOfFunctions = append(listOfFunctions, Lambda{
			Name:    *lambda.FunctionName,
			Memory:  *lambda.MemorySize,
			Runtime: runtime,
		})
	}
	return listOfFunctions, nil
}
