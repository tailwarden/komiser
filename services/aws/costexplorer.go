package aws

import (
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	models "github.com/mlabouardy/komiser/models/aws"
)

type Recommendation struct {
	Currency   string                 `json:"currency"`
	Total      string                 `json:"total"`
	Percentage string                 `json:"percentage"`
	Details    []RecommendationDetail `json:"details"`
}

type RecommendationDetail struct {
	AccountId                              string `json:"accountId"`
	Family                                 string `json:"instanceFamily"`
	InstanceType                           string `json:"instanceType"`
	Region                                 string `json:"region"`
	Platform                               string `json:"platform"`
	RecommendedNumberOfInstancesToPurchase string `json:"numberOfInstancesToPurchase"`
	EstimatedMonthlySavingsAmount          string `json:"estimatedMonthlySavings"`
	EstimatedMonthlyOnDemandCost           string `json:"estimatedMonthlyOnDemandCost"`
	UpfrontCost                            string `json:"upfrontCost"`
	RecurringStandardMonthlyCost           string `json:"recurringStandardMonthlyCost"`
}

func (awsClient AWS) DescribeCostAndUsage(cfg aws.Config) (models.Bill, error) {
	currentTime := time.Now().Local()
	start := currentTime.AddDate(0, -6, 0).Format("2006-01-02")
	end := currentTime.Format("2006-01-02")
	cfg.Region = "us-east-1"
	svc := costexplorer.NewFromConfig(cfg)
	result, err := svc.GetCostAndUsage(context.Background(), &costexplorer.GetCostAndUsageInput{
		Metrics:     []string{"BlendedCost"},
		Granularity: types.GranularityMonthly,
		TimePeriod: &types.DateInterval{
			Start: &start,
			End:   &end,
		},
		GroupBy: []types.GroupDefinition{
			types.GroupDefinition{
				Key:  aws.String("SERVICE"),
				Type: types.GroupDefinitionTypeDimension,
			},
		},
	})
	if err != nil {
		return models.Bill{}, err
	}

	costs := make([]models.Cost, 0)
	for _, res := range result.ResultsByTime {
		start, _ := time.Parse("2006-01-02", *res.TimePeriod.Start)
		end, _ := time.Parse("2006-01-02", *res.TimePeriod.End)

		unit := "USD"

		groups := make([]models.Group, 0)
		for _, group := range res.Groups {
			amount, _ := strconv.ParseFloat(*group.Metrics["BlendedCost"].Amount, 64)
			groups = append(groups, models.Group{
				Key:    group.Keys[0],
				Amount: amount,
			})
			unit = *group.Metrics["BlendedCost"].Unit
		}

		sort.Slice(groups, func(i, j int) bool {
			return groups[i].Amount > groups[j].Amount
		})

		costs = append(costs, models.Cost{
			Start:  start,
			End:    end,
			Unit:   unit,
			Groups: groups,
		})
	}

	var currentBill float64
	for _, group := range costs[len(costs)-1].Groups {
		currentBill += group.Amount
	}

	return models.Bill{
		Total:   currentBill,
		History: costs,
	}, nil
}

func (awsClient AWS) DescribeCostAndUsagePerInstanceType(cfg aws.Config) (models.Bill, error) {
	currentTime := time.Now().Local()
	start := currentTime.AddDate(0, -6, 0).Format("2006-01-02")
	end := currentTime.Format("2006-01-02")
	cfg.Region = "us-east-1"
	svc := costexplorer.NewFromConfig(cfg)
	result, err := svc.GetCostAndUsage(context.Background(), &costexplorer.GetCostAndUsageInput{
		Metrics:     []string{"BlendedCost"},
		Granularity: types.GranularityMonthly,
		TimePeriod: &types.DateInterval{
			Start: &start,
			End:   &end,
		},
		GroupBy: []types.GroupDefinition{
			types.GroupDefinition{
				Key:  aws.String("INSTANCE_TYPE"),
				Type: types.GroupDefinitionTypeDimension,
			},
		},
		Filter: &types.Expression{
			Dimensions: &types.DimensionValues{
				Key:    types.DimensionService,
				Values: []string{"Amazon Elastic Compute Cloud - Compute"},
			},
		},
	})
	if err != nil {
		return models.Bill{}, err
	}

	costs := make([]models.Cost, 0)
	for _, res := range result.ResultsByTime {
		start, _ := time.Parse("2006-01-02", *res.TimePeriod.Start)
		end, _ := time.Parse("2006-01-02", *res.TimePeriod.End)

		unit := "USD"

		groups := make([]models.Group, 0)
		for _, group := range res.Groups {
			amount, _ := strconv.ParseFloat(*group.Metrics["BlendedCost"].Amount, 64)
			groups = append(groups, models.Group{
				Key:    group.Keys[0],
				Amount: amount,
			})
			unit = *group.Metrics["BlendedCost"].Unit
		}

		sort.Slice(groups, func(i, j int) bool {
			return groups[i].Amount > groups[j].Amount
		})

		costs = append(costs, models.Cost{
			Start:  start,
			End:    end,
			Unit:   unit,
			Groups: groups,
		})
	}

	var currentBill float64
	for _, group := range costs[len(costs)-1].Groups {
		currentBill += group.Amount
	}

	return models.Bill{
		Total:   currentBill,
		History: costs,
	}, nil
}

func (awsClient AWS) DescribeForecastPrice(cfg aws.Config) (string, error) {
	currentTime := time.Now().Local()
	start := currentTime.AddDate(0, 0, 1).Format("2006-01-02")
	end := currentTime.AddDate(0, 1, -currentTime.Day()).Format("2006-01-02")
	cfg.Region = "us-east-1"
	svc := costexplorer.NewFromConfig(cfg)
	result, err := svc.GetCostForecast(context.Background(), &costexplorer.GetCostForecastInput{
		Metric:      types.MetricBlendedCost,
		Granularity: types.GranularityMonthly,
		TimePeriod: &types.DateInterval{
			Start: &start,
			End:   &end,
		},
	})
	if err != nil {
		return "", nil
	}

	return *result.Total.Amount, nil
}

func (awsClient AWS) DescribeReservationRecommendations(period types.LookbackPeriodInDays, payment types.PaymentOption, terms types.TermInYears, cfg aws.Config) ([]Recommendation, error) {
	recommendations := make([]Recommendation, 0)

	cfg.Region = "us-east-1"
	svc := costexplorer.NewFromConfig(cfg)
	res, err := svc.GetReservationPurchaseRecommendation(context.Background(), &costexplorer.GetReservationPurchaseRecommendationInput{
		LookbackPeriodInDays: period,
		PaymentOption:        payment,
		Service:              aws.String("Amazon Elastic Compute Cloud - Compute"),
		TermInYears:          terms,
	})
	if err != nil {
		return recommendations, err
	}

	for _, rec := range res.Recommendations {
		recommendation := Recommendation{
			Currency:   *rec.RecommendationSummary.CurrencyCode,
			Total:      *rec.RecommendationSummary.TotalEstimatedMonthlySavingsAmount,
			Percentage: *rec.RecommendationSummary.TotalEstimatedMonthlySavingsPercentage,
		}

		details := make([]RecommendationDetail, 0)

		for _, d := range rec.RecommendationDetails {
			details = append(details, RecommendationDetail{
				AccountId:                              *d.AccountId,
				Family:                                 *d.InstanceDetails.EC2InstanceDetails.Family,
				InstanceType:                           *d.InstanceDetails.EC2InstanceDetails.InstanceType,
				Region:                                 *d.InstanceDetails.EC2InstanceDetails.Region,
				Platform:                               *d.InstanceDetails.EC2InstanceDetails.Platform,
				RecommendedNumberOfInstancesToPurchase: *d.RecommendedNumberOfInstancesToPurchase,
				EstimatedMonthlySavingsAmount:          *d.EstimatedMonthlySavingsAmount,
				EstimatedMonthlyOnDemandCost:           *d.EstimatedMonthlyOnDemandCost,
				UpfrontCost:                            *d.UpfrontCost,
				RecurringStandardMonthlyCost:           *d.RecurringStandardMonthlyCost,
			})
		}

		recommendation.Details = details

		recommendations = append(recommendations, recommendation)
	}

	return recommendations, nil
}
