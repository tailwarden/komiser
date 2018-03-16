package services

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeCostAndUsage(cfg aws.Config) ([]Cost, error) {
	currentTime := time.Now().Local()
	start := currentTime.AddDate(0, -6, 0).Format("2006-01-02")
	end := currentTime.Format("2006-01-02")
	svc := costexplorer.New(cfg)
	req := svc.GetCostAndUsageRequest(&costexplorer.GetCostAndUsageInput{
		Metrics:     []string{"BlendedCost"},
		Granularity: costexplorer.GranularityMonthly,
		TimePeriod: &costexplorer.DateInterval{
			Start: &start,
			End:   &end,
		},
	})
	result, err := req.Send()
	if err != nil {
		return []Cost{}, err
	}
	costs := make([]Cost, 0)
	for _, res := range result.ResultsByTime {
		start, _ := time.Parse("2006-01-02", *res.TimePeriod.Start)
		end, _ := time.Parse("2006-01-02", *res.TimePeriod.End)
		amount, _ := strconv.ParseFloat(*res.Total["BlendedCost"].Amount, 64)
		costs = append(costs, Cost{
			Start:  start,
			End:    end,
			Amount: amount,
			Unit:   *res.Total["BlendedCost"].Unit,
		})
	}
	return costs, nil
}
