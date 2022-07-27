package aws

import (
	"context"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go/aws"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) OpenSupportTickets(cfg awsConfig.Config) ([]Ticket, error) {
	tickets := make([]Ticket, 0)

	cfg.Region = "us-east-1"
	svc := support.NewFromConfig(cfg)
	res, err := svc.DescribeCases(context.Background(), &support.DescribeCasesInput{})
	if err != nil {
		return tickets, err
	}

	for _, ticket := range res.Cases {
		timestamp, _ := time.Parse("2014-11-12T11:45:26.371Z", *ticket.TimeCreated)
		tickets = append(tickets, Ticket{
			Timestamp:    timestamp,
			CategoryCode: *ticket.CategoryCode,
			ServiceCode:  *ticket.ServiceCode,
			SeverityCode: *ticket.SeverityCode,
			Status:       *ticket.Status,
		})
	}

	return tickets, nil
}

func (awsClient AWS) TicketsInLastSixMonthsTickets(cfg awsConfig.Config) ([]Ticket, error) {
	tickets := make([]Ticket, 0)

	cfg.Region = "us-east-1"
	svc := support.NewFromConfig(cfg)
	res, err := svc.DescribeCases(context.Background(), &support.DescribeCasesInput{
		IncludeResolvedCases: true,
		AfterTime:            aws.String(aws.Time(time.Now().AddDate(0, -6, 0)).Format("2006-01-02")),
	})
	if err != nil {
		return tickets, err
	}

	for _, ticket := range res.Cases {
		timestamp, _ := time.Parse("2006-01-02T15:04:05.000Z", *ticket.TimeCreated)
		tickets = append(tickets, Ticket{
			Timestamp:    timestamp,
			CategoryCode: *ticket.CategoryCode,
			ServiceCode:  *ticket.ServiceCode,
			SeverityCode: *ticket.SeverityCode,
			Status:       *ticket.Status,
		})
	}

	return tickets, nil
}

type ServiceLimit struct {
	CheckId string `json:"checkId"`
	Name    string `json:"name"`
	Status  string `json:"status"`
}

func (awsClient AWS) DescribeServiceLimitsChecks(cfg awsConfig.Config) ([]ServiceLimit, error) {
	limits := make([]ServiceLimit, 0)

	cfg.Region = "us-east-1"
	svc := support.NewFromConfig(cfg)
	res, err := svc.DescribeTrustedAdvisorChecks(context.Background(), &support.DescribeTrustedAdvisorChecksInput{
		Language: aws.String("en"),
	})
	if err != nil {
		return limits, err
	}

	for _, check := range res.Checks {
		if *check.Category == "service_limits" {
			res, err := svc.DescribeTrustedAdvisorCheckResult(context.Background(), &support.DescribeTrustedAdvisorCheckResultInput{
				CheckId: check.Id,
			})

			if err != nil {
				return limits, err
			}

			limits = append(limits, ServiceLimit{
				CheckId: *check.Id,
				Name:    *check.Name,
				Status:  *res.Result.Status,
			})
		}
	}

	return limits, nil
}
