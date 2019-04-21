package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) OpenSupportTickets(cfg aws.Config) ([]Ticket, error) {
	tickets := make([]Ticket, 0)

	cfg.Region = "us-east-1"
	svc := support.New(cfg)
	req := svc.DescribeCasesRequest(&support.DescribeCasesInput{})
	res, err := req.Send(context.Background())
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

func (awsClient AWS) TicketsInLastSixMonthsTickets(cfg aws.Config) ([]Ticket, error) {
	tickets := make([]Ticket, 0)

	cfg.Region = "us-east-1"
	svc := support.New(cfg)
	req := svc.DescribeCasesRequest(&support.DescribeCasesInput{
		IncludeResolvedCases: aws.Bool(true),
		AfterTime:            aws.String(aws.Time(time.Now().AddDate(0, -6, 0)).Format("2006-01-02")),
	})
	res, err := req.Send(context.Background())
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

func (awsClient AWS) DescribeServiceLimitsChecks(cfg aws.Config) ([]ServiceLimit, error) {
	limits := make([]ServiceLimit, 0)

	cfg.Region = "us-east-1"
	svc := support.New(cfg)
	req := svc.DescribeTrustedAdvisorChecksRequest(&support.DescribeTrustedAdvisorChecksInput{
		Language: aws.String("en"),
	})
	res, err := req.Send(context.Background())
	if err != nil {
		return limits, err
	}

	for _, check := range res.Checks {
		if *check.Category == "service_limits" {
			reqCheckResult := svc.DescribeTrustedAdvisorCheckResultRequest(&support.DescribeTrustedAdvisorCheckResultInput{
				CheckId: check.Id,
			})

			res, err := reqCheckResult.Send(context.Background())
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
