package slack

import (
	"context"
	"fmt"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	. "github.com/mlabouardy/komiser/handlers/aws"
	. "github.com/mlabouardy/komiser/handlers/azure"
	"github.com/mlabouardy/komiser/handlers/azure/config"
	. "github.com/mlabouardy/komiser/handlers/gcp"
	. "github.com/mlabouardy/komiser/services/ini"
	slackClient "github.com/slack-go/slack"
)

type Slack struct {
	Channel string
	Token   string
}

func (slack *Slack) SetCredentials(token string, channel string) {
	slack.Channel = channel
	slack.Token = token
}

func (slack *Slack) sendCostAlert(provider string, logo string, costs float64, label string) {
	slackApi := slackClient.New(slack.Token)

	attachment := slackClient.Attachment{
		Color: "good",
		Fields: []slackClient.AttachmentField{
			slackClient.AttachmentField{
				Title: "Provider",
				Value: provider,
			},
			slackClient.AttachmentField{
				Title: "Label",
				Value: label,
			},
			slackClient.AttachmentField{
				Title: "Monthly Cost",
				Value: fmt.Sprintf("%.2f $", costs),
			},
		},
		ThumbURL: logo,
		Footer:   "Get more information at https://docs.komiser.io",
	}
	slackApi.PostMessage(slack.Channel, slackClient.MsgOptionAttachments(attachment))
}

func (slack *Slack) SendDailyNotification(awsHandler *AWSHandler, gcpHandler *GCPHandler, azureHandler *AzureHandler) {
	// AWS Alert
	if awsHandler.HasMultipleEnvs() {
		profiles, _ := OpenFile(awsConfig.DefaultSharedCredentialsFilename())
		for _, profile := range profiles.List() {
			cfg, _ := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithSharedConfigProfile(profile))
			bill, err := awsHandler.GetAWSHandler().DescribeCostAndUsage(cfg)
			if err == nil {
				slack.sendCostAlert("Amazon Web Services", "https://cdn.komiser.io/images/aws.png", bill.Total, profile)
			}
		}
	} else {
		cfg, _ := awsConfig.LoadDefaultConfig(context.Background())
		bill, err := awsHandler.GetAWSHandler().DescribeCostAndUsage(cfg)
		if err == nil {
			slack.sendCostAlert("Amazon Web Services", "https://cdn.komiser.io/images/aws.png", bill.Total, "default")
		}
	}

	// GCP Alert
	costs, err := gcpHandler.GetGCPHandler().CostInLastSixMonths()
	if err == nil && len(costs) > 0 {
		slack.sendCostAlert("Google Cloud Platform", "https://cdn.komiser.io/images/gcp.png", costs[len(costs)-1].Cost, "default")
	}

	// Azure Alert
	err = config.ParseEnvironment()
	if err == nil {
		bill, err := azureHandler.GetAzureHandler().GetBilling(config.SubscriptionID())
		if err == nil {
			slack.sendCostAlert("Microsoft Azure", "https://swimburger.net/media/fbqnp2ie/azure.svg", bill.Amount, "default")
		}
	}
}
