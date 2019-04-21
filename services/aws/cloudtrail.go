package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

func (awsClient AWS) CloudTrailConsoleSignInEvents(cfg aws.Config) (map[string]map[string]int64, error) {
	events := make(map[string]map[string]int64, 0)

	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return events, err
	}

	for _, region := range regions {
		cfg.Region = region.Name
		cloudtrailClient := cloudtrail.New(cfg)
		req := cloudtrailClient.LookupEventsRequest(&cloudtrail.LookupEventsInput{
			LookupAttributes: []cloudtrail.LookupAttribute{
				cloudtrail.LookupAttribute{
					AttributeKey:   cloudtrail.LookupAttributeKeyEventName,
					AttributeValue: aws.String("ConsoleLogin"),
				},
			},
			StartTime: aws.Time(time.Now().AddDate(0, 0, -7)),
			EndTime:   aws.Time(time.Now()),
		})
		res, err := req.Send(context.Background())
		if err != nil {
			return events, err
		}

		for _, event := range res.Events {
			timestamp := (*event.EventTime).Format("2006-01-02")
			username := *event.Username

			if events[timestamp] == nil {
				events[timestamp] = make(map[string]int64, 1)
				events[timestamp][username] = 0
			} else {
				events[timestamp][username]++
			}
		}
	}

	return events, nil
}

type CloudTrailEvent struct {
	SourceIPAddress string `json:"sourceIPAddress"`
}

type SourceIp struct {
	Coordinate Coordinate `json:"coordinate"`
	Total      int64      `json:"total"`
}

func (awsClient AWS) CloudTrailConsoleSignInSourceIpEvents(cfg aws.Config) (map[string]SourceIp, error) {
	events := make(map[string]SourceIp, 0)

	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return events, err
	}

	for _, region := range regions {
		cfg.Region = region.Name
		cloudtrailClient := cloudtrail.New(cfg)
		req := cloudtrailClient.LookupEventsRequest(&cloudtrail.LookupEventsInput{
			LookupAttributes: []cloudtrail.LookupAttribute{
				cloudtrail.LookupAttribute{
					AttributeKey:   cloudtrail.LookupAttributeKeyEventName,
					AttributeValue: aws.String("ConsoleLogin"),
				},
			},
			StartTime: aws.Time(time.Now().AddDate(0, 0, -7)),
			EndTime:   aws.Time(time.Now()),
		})
		res, err := req.Send(context.Background())
		if err != nil {
			return events, err
		}

		for _, event := range res.Events {
			cloudtrailEvent := CloudTrailEvent{}

			json.Unmarshal([]byte(*event.CloudTrailEvent), &cloudtrailEvent)

			if events[cloudtrailEvent.SourceIPAddress] == (SourceIp{}) {
				coordinate, err := getCoordinates(cloudtrailEvent.SourceIPAddress)
				if err != nil {
					return events, err
				}

				events[cloudtrailEvent.SourceIPAddress] = SourceIp{
					Coordinate: coordinate,
					Total:      0,
				}
			} else {
				events[cloudtrailEvent.SourceIPAddress] = SourceIp{
					Coordinate: events[cloudtrailEvent.SourceIPAddress].Coordinate,
					Total:      events[cloudtrailEvent.SourceIPAddress].Total + 1,
				}
			}
		}
	}

	return events, nil
}

type Coordinate struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func getCoordinates(ip string) (Coordinate, error) {
	coordinate := Coordinate{}

	url := fmt.Sprintf(`http://ip-api.com/json/%s`, ip)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return coordinate, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return coordinate, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(data, &coordinate)

	return coordinate, nil
}
