package integrations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	. "github.com/mlabouardy/komiser/handlers/aws"
	. "github.com/mlabouardy/komiser/handlers/azure"
	azureConfig "github.com/mlabouardy/komiser/handlers/azure/config"
	. "github.com/mlabouardy/komiser/handlers/digitalocean"
	. "github.com/mlabouardy/komiser/handlers/gcp"
	models "github.com/mlabouardy/komiser/models/aws"
	"github.com/mlabouardy/komiser/services/aws"
	. "github.com/mlabouardy/komiser/services/cache"
	. "github.com/mlabouardy/komiser/services/ini"
	. "github.com/mlabouardy/komiser/services/integrations/slack"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Region struct {
	Name      string `json:"name" bson:"name"`
	Label     string `json:"label" bson:"label"`
	Latitude  string `json:"latitude" bson:"latitude"`
	Longitude string `json:"longitude" bson:"longitude"`
}

type AccountHandler struct {
	cache               Cache
	awsHandler          *AWSHandler
	gcpHandler          *GCPHandler
	azureHandler        *AzureHandler
	digitaloceanHandler *DigitalOceanHandler
	slack               Slack
	services            map[string]interface{}
}

type View struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Tags []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"tags"`
}

func NewAccountHandler(cache Cache, awsHandler *AWSHandler, gcpHandler *GCPHandler, azureHandler *AzureHandler, digitaloceanHandler *DigitalOceanHandler, services map[string]interface{}) *AccountHandler {
	accountHandler := AccountHandler{
		cache:               cache,
		awsHandler:          awsHandler,
		gcpHandler:          gcpHandler,
		azureHandler:        azureHandler,
		digitaloceanHandler: digitaloceanHandler,
		services:            services,
	}
	return &accountHandler
}

func (handler *AccountHandler) ListCloudAccountsHandler(w http.ResponseWriter, r *http.Request) {
	accounts := make(map[string][]string, 0)

	sections, err := OpenFile(config.DefaultSharedCredentialsFilename())
	if err == nil {
		for _, section := range sections.List() {
			accounts["AWS"] = append(accounts["AWS"], section)
		}
	}

	_, err = google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err == nil {
		accounts["GCP"] = append(accounts["GCP"], "Default")
	}

	err = azureConfig.ParseEnvironment()
	if err == nil {
		subscriptionID := azureConfig.SubscriptionID()
		accounts["AZURE"] = append(accounts["AZURE"], subscriptionID)
	}

	if os.Getenv("DIGITALOCEAN_ACCESS_TOKEN") != "" {
		accounts["DIGITALOCEAN"] = append(accounts["DIGITALOCEAN"], "Default")
	}

	respondWithJSON(w, 200, accounts)
}

func (handler *AccountHandler) ListActiveRegionsHandler(w http.ResponseWriter, r *http.Request) {
	regions := make([]string, 0)

	response, found := handler.cache.Get("komiser.regions")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		sections, err := OpenFile(config.DefaultSharedCredentialsFilename())
		if err == nil {
			for _, section := range sections.List() {
				cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(section))
				awsInstances, err := handler.awsHandler.GetAWSHandler().DescribeInstances(cfg)
				if err == nil {
					for _, instance := range awsInstances {
						found := false
						for _, region := range regions {
							if region == instance.Region {
								found = true
							}
						}
						if !found {
							regions = append(regions, instance.Region)
						}
					}
				}
			}
		}

		gcpInstances, err := handler.gcpHandler.GetGCPHandler().GetComputeInstances()
		if err == nil {
			for _, instance := range gcpInstances {
				found := false
				for _, region := range regions {
					if region == instance.Region {
						found = true
					}
				}
				if !found {
					regions = append(regions, instance.Region)
				}
			}
		}

		err = azureConfig.ParseEnvironment()
		if err == nil {
			subscriptionID := azureConfig.SubscriptionID()
			azureInstances, err := handler.azureHandler.GetAzureHandler().DescribeVMs(subscriptionID)
			if err == nil {
				for _, instance := range azureInstances {
					found := false
					for _, region := range regions {
						if region == instance.Region {
							found = true
						}
					}
					if !found {
						regions = append(regions, instance.Region)
					}
				}
			}
		}

		listOfActiveRegions := make([]Region, 0)
		listOfSupportedRegions := getSupportedRegions()

		for _, region := range regions {
			for _, regionToCompareWith := range listOfSupportedRegions {
				if regionToCompareWith.Label == region {
					listOfActiveRegions = append(listOfActiveRegions, regionToCompareWith)
				}
			}
		}

		handler.cache.Set("komiser.regions", listOfActiveRegions)
		respondWithJSON(w, 200, regions)
	}
}

func (handler *AccountHandler) CostBreakdownByCloudProviderHandler(w http.ResponseWriter, r *http.Request) {
	costs := make(map[string]float64, 0)

	sections, err := OpenFile(config.DefaultSharedCredentialsFilename())
	if err == nil {
		for _, section := range sections.List() {
			cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(section))
			cost, err := handler.awsHandler.GetAWSHandler().DescribeCostAndUsage(cfg)
			if err == nil {
				costs["AWS"] += cost.Total
			}
		}
	}

	respondWithJSON(w, 200, costs)
}

func (handler *AccountHandler) CostBreakdownByCloudAccountHandler(w http.ResponseWriter, r *http.Request) {
	costs := make(map[string]float64, 0)

	sections, err := OpenFile(config.DefaultSharedCredentialsFilename())
	if err == nil {
		for _, section := range sections.List() {
			cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(section))
			cost, err := handler.awsHandler.GetAWSHandler().DescribeCostAndUsage(cfg)
			if err == nil {
				costs[fmt.Sprintf("AWS:%s", section)] += cost.Total
			}
		}
	}

	respondWithJSON(w, 200, costs)
}

func (handler *AccountHandler) CostBreakdownByCloudRegionHandler(w http.ResponseWriter, r *http.Request) {
	costs := make(map[string]float64, 0)

	sections, err := OpenFile(config.DefaultSharedCredentialsFilename())
	if err == nil {
		for _, section := range sections.List() {
			cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(section))
			cost, err := handler.awsHandler.GetAWSHandler().DescribeCostAndUsageByRegion(cfg)
			if err == nil {
				if len(cost.History) > 0 {
					for _, k := range cost.History[len(cost.History)-1].Groups {
						costs[fmt.Sprintf("AWS:%s", k.Key)] += k.Amount
					}
				}
			}
		}
	}

	respondWithJSON(w, 200, costs)
}

func (handler *AccountHandler) ListViewsHandler(w http.ResponseWriter, r *http.Request) {
	views := make([]View, 0)

	response, found := handler.cache.Get("komiser.views")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		handler.cache.Set("komiser.views", views)
		respondWithJSON(w, 200, views)
	}
}

func (handler *AccountHandler) NewViewHandler(w http.ResponseWriter, r *http.Request) {
	var view View
	json.NewDecoder(r.Body).Decode(&view)

	response, found := handler.cache.Get("komiser.views")
	if found {
		views := response.([]View)
		views = append(views, view)
		handler.cache.Set("komiser.views", views)
		respondWithJSON(w, 200, view)
	} else {
		views := make([]View, 0)
		views = append(views, view)
		handler.cache.Set("komiser.views", views)
		respondWithJSON(w, 200, view)
	}
}

func (handler *AccountHandler) FilterbyTagsHandler(w http.ResponseWriter, r *http.Request) {
	var tags []Tag
	json.NewDecoder(r.Body).Decode(&tags)

	listOfResources := make([]interface{}, 0)

	for key, resources := range handler.services {
		if key == "aws:instances" {
			instances := resources.([]aws.EC2Instance)
			for _, instance := range instances {
				if hasMatchingTags(instance.Tags, tags) {
					listOfResources = append(listOfResources, instance)
				}
			}
		} else if key == "aws:lambda" {
			instances := resources.([]models.Lambda)
			for _, instance := range instances {
				if hasMatchingTags(instance.Tags, tags) {
					listOfResources = append(listOfResources, instance)
				}
			}
		} else if key == "aws:buckets" {
			instances := resources.([]aws.S3Bucket)
			for _, instance := range instances {
				if hasMatchingTags(instance.Tags, tags) {
					listOfResources = append(listOfResources, instance)
				}
			}
		} else if key == "aws:tables" {
			instances := resources.([]aws.AWSDynamoDBTable)
			for _, instance := range instances {
				if hasMatchingTags(instance.Tags, tags) {
					listOfResources = append(listOfResources, instance)
				}
			}
		} else if key == "aws:vpcs" {
			instances := resources.([]aws.AWSVPC)
			for _, instance := range instances {
				if hasMatchingTags(instance.Tags, tags) {
					listOfResources = append(listOfResources, instance)
				}
			}
		} else if key == "aws:routetables" {
			instances := resources.([]aws.AWSRouteTable)
			for _, instance := range instances {
				if hasMatchingTags(instance.Tags, tags) {
					listOfResources = append(listOfResources, instance)
				}
			}
		} else if key == "aws:securitygroups" {
			instances := resources.([]models.SecurityGroup)
			for _, instance := range instances {
				if hasMatchingTags(instance.Tags, tags) {
					listOfResources = append(listOfResources, instance)
				}
			}
		} else if key == "aws:queues" {
			instances := resources.([]models.Queue)
			for _, instance := range instances {
				if hasMatchingTags(instance.Tags, tags) {
					listOfResources = append(listOfResources, instance)
				}
			}
		} else if key == "aws:ecs" {
			instances := resources.(aws.ECSData)
			for _, instance := range instances.Clusters {
				if hasMatchingTags(instance.Tags, tags) {
					listOfResources = append(listOfResources, instance)
				}
			}
		}
	}

	respondWithJSON(w, 200, listOfResources)
}

func hasMatchingTags(sources []string, values []Tag) bool {
	matches := 0
	for _, source := range sources {
		for _, value := range values {

			if strings.Contains(source, value.Key) || strings.Contains(source, value.Value) {
				matches++
			}
		}
	}
	return matches == len(sources)
}

func getSupportedRegions() []Region {
	return []Region{
		Region{
			Name:      "Ohio",
			Label:     "us-east-2",
			Latitude:  "40.367474",
			Longitude: "-82.996216",
		},
		Region{
			Name:      "N.Virginia",
			Label:     "us-east-1",
			Latitude:  "37.926868",
			Longitude: "-78.024902",
		},
		Region{
			Name:      "N.California",
			Label:     "us-west-1",
			Latitude:  "36.778261",
			Longitude: "-119.4179324",
		},
		Region{
			Name:      "Oregon",
			Label:     "us-west-2",
			Latitude:  "45.523062",
			Longitude: "-122.676482",
		},
		Region{
			Name:      "Cape Town",
			Label:     "af-south-1",
			Latitude:  "-33.924869",
			Longitude: "18.424055",
		},
		Region{
			Name:      "Hong Kong",
			Label:     "ap-east-1",
			Latitude:  "22.302711",
			Longitude: "114.177216",
		},
		Region{
			Name:      "Jakarta",
			Label:     "ap-southeast-3",
			Latitude:  "-6.2087634",
			Longitude: "106.816666",
		},
		Region{
			Name:      "Mumbai",
			Label:     "ap-south-1",
			Latitude:  "19.076090",
			Longitude: "72.877426",
		},
		Region{
			Name:      "Osaka",
			Label:     "ap-northeast-3",
			Latitude:  "34.6937378",
			Longitude: "135.5021651",
		},
		Region{
			Name:      "Seoul",
			Label:     "ap-northeast-2",
			Latitude:  "37.566535",
			Longitude: "126.9779692",
		},
		Region{
			Name:      "Singapore",
			Label:     "ap-southeast-1",
			Latitude:  "1.290270",
			Longitude: "103.851959",
		},
		Region{
			Name:      "Sydney",
			Label:     "ap-southeast-2",
			Latitude:  "-33.8667",
			Longitude: "151.206990",
		},
		Region{
			Name:      "Tokyo",
			Label:     "ap-northeast-1",
			Latitude:  "35.652832",
			Longitude: "139.839478",
		},
		Region{
			Name:      "Canada",
			Label:     "ca-central-1",
			Latitude:  "-79.347015",
			Longitude: "43.651070",
		},
		Region{
			Name:      "Frankfurt",
			Label:     "eu-central-1",
			Latitude:  "50.1109221",
			Longitude: "8.6821267",
		},
		Region{
			Name:      "Ireland",
			Label:     "eu-west-1",
			Latitude:  "53.350140",
			Longitude: "-6.266155",
		},
		Region{
			Name:      "London",
			Label:     "eu-west-2",
			Latitude:  "51.5073509",
			Longitude: "-0.1277583",
		},
		Region{
			Name:      "Milan",
			Label:     "eu-south-1",
			Latitude:  "45.4654219",
			Longitude: "9.1859243",
		},
		Region{
			Name:      "Paris",
			Label:     "eu-west-3",
			Latitude:  "2.349014",
			Longitude: "48.864716",
		},
		Region{
			Name:      "Stockholm",
			Label:     "eu-north-1",
			Latitude:  "59.334591",
			Longitude: "18.063240",
		},
		Region{
			Name:      "Bahrain",
			Label:     "me-south-1",
			Latitude:  "26.066700",
			Longitude: "50.557700",
		},
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
