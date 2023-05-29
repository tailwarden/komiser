package utils

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

type Location struct {
	Name      string `json:"name"`
	Label     string `json:"label"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func getAWSRegions() []Location {
	return []Location{
		{
			Name:      "Ohio",
			Label:     "us-east-2",
			Latitude:  "40.367474",
			Longitude: "-82.996216",
		},
		{
			Name:      "N.Virginia",
			Label:     "us-east-1",
			Latitude:  "37.926868",
			Longitude: "-78.024902",
		},
		{
			Name:      "N.California",
			Label:     "us-west-1",
			Latitude:  "36.778261",
			Longitude: "-119.4179324",
		},
		{
			Name:      "Oregon",
			Label:     "us-west-2",
			Latitude:  "45.523062",
			Longitude: "-122.676482",
		},
		{
			Name:      "Cape Town",
			Label:     "af-south-1",
			Latitude:  "-33.924869",
			Longitude: "18.424055",
		},
		{
			Name:      "Hong Kong",
			Label:     "ap-east-1",
			Latitude:  "22.302711",
			Longitude: "114.177216",
		},
		{
			Name:      "Jakarta",
			Label:     "ap-southeast-3",
			Latitude:  "-6.2087634",
			Longitude: "106.816666",
		},
		{
			Name:      "Mumbai",
			Label:     "ap-south-1",
			Latitude:  "19.076090",
			Longitude: "72.877426",
		},
		{
			Name:      "Osaka",
			Label:     "ap-northeast-3",
			Latitude:  "34.6937378",
			Longitude: "135.5021651",
		},
		{
			Name:      "Seoul",
			Label:     "ap-northeast-2",
			Latitude:  "37.566535",
			Longitude: "126.9779692",
		},
		{
			Name:      "Singapore",
			Label:     "ap-southeast-1",
			Latitude:  "1.290270",
			Longitude: "103.851959",
		},
		{
			Name:      "Sydney",
			Label:     "ap-southeast-2",
			Latitude:  "-33.8667",
			Longitude: "151.206990",
		},
		{
			Name:      "Tokyo",
			Label:     "ap-northeast-1",
			Latitude:  "35.652832",
			Longitude: "139.839478",
		},
		{
			Name:      "Canada",
			Label:     "ca-central-1",
			Latitude:  "56.130367",
			Longitude: "-106.346771",
		},
		{
			Name:      "Frankfurt",
			Label:     "eu-central-1",
			Latitude:  "50.1109221",
			Longitude: "8.6821267",
		},
		{
			Name:      "Ireland",
			Label:     "eu-west-1",
			Latitude:  "53.350140",
			Longitude: "-6.266155",
		},
		{
			Name:      "London",
			Label:     "eu-west-2",
			Latitude:  "51.5073509",
			Longitude: "-0.1277583",
		},
		{
			Name:      "Milan",
			Label:     "eu-south-1",
			Latitude:  "45.4654219",
			Longitude: "9.1859243",
		},
		{
			Name:      "Paris",
			Label:     "eu-west-3",
			Latitude:  "48.864716",
			Longitude: "2.352222",
		},
		{
			Name:      "Stockholm",
			Label:     "eu-north-1",
			Latitude:  "59.334591",
			Longitude: "18.063240",
		},
		{
			Name:      "Bahrain",
			Label:     "me-south-1",
			Latitude:  "26.066700",
			Longitude: "50.557700",
		},
		{
			Name:      "Sao Paulo",
			Label:     "sa-east-1",
			Latitude:  "-23.533773",
			Longitude: "-46.625290",
		},
	}
}

func getDigitalOceanRegions() []Location {
	return []Location{
		{
			Name:      "New York City",
			Label:     "NYC1",
			Latitude:  "40.712776",
			Longitude: "-74.005974",
		},
		{
			Name:      "New York City",
			Label:     "NYC2",
			Latitude:  "40.712776",
			Longitude: "-74.005974",
		},
		{
			Name:      "New York City",
			Label:     "NYC3",
			Latitude:  "40.712776",
			Longitude: "-74.005974",
		},
		{
			Name:      "San Francisco",
			Label:     "SFO1",
			Latitude:  "37.774929",
			Longitude: "-122.419418",
		},
		{
			Name:      "San Francisco",
			Label:     "SFO2",
			Latitude:  "37.774929",
			Longitude: "-122.419418",
		},
		{
			Name:      "Toronto",
			Label:     "TOR1",
			Latitude:  "43.653225",
			Longitude: "-79.383186",
		},
		{
			Name:      "London",
			Label:     "LON1",
			Latitude:  "51.507351",
			Longitude: "-0.127758",
		},
		{
			Name:      "Frankfurt",
			Label:     "FRA1",
			Latitude:  "50.110924",
			Longitude: "8.682127",
		},
		{
			Name:      "Amsterdam",
			Label:     "AMS2",
			Latitude:  "52.377956",
			Longitude: "4.897070",
		},
		{
			Name:      "Amsterdam",
			Label:     "AMS3",
			Latitude:  "52.377956",
			Longitude: "4.897070",
		},
		{
			Name:      "Singapore",
			Label:     "SGP1",
			Latitude:  "1.290270",
			Longitude: "103.851959",
		},
		{
			Name:      "Bangalore",
			Label:     "BLR1",
			Latitude:  "12.972442",
			Longitude: "77.580643",
		},
	}
}

func getGCPZones() []Location {
	return []Location{
		// asia-east1
		{
			Name:      "Changhua County, Taiwan",
			Label:     "asia-east1-a",
			Latitude:  "24.075",
			Longitude: "120.513",
		},
		{
			Name:      "Changhua County, Taiwan",
			Label:     "asia-east1-b",
			Latitude:  "24.075",
			Longitude: "120.513",
		},
		{
			Name:      "Changhua County, Taiwan",
			Label:     "asia-east1-c",
			Latitude:  "24.075",
			Longitude: "120.513",
		},
		// asia-east2
		{
			Name:      "Hong Kong",
			Label:     "asia-east2-a",
			Latitude:  "22.3964",
			Longitude: "114.1095",
		},
		{
			Name:      "Hong Kong",
			Label:     "asia-east2-b",
			Latitude:  "22.3964",
			Longitude: "114.1095",
		},
		{
			Name:      "Hong Kong",
			Label:     "asia-east2-c",
			Latitude:  "22.3964",
			Longitude: "114.1095",
		},
		// asia-northeast1
		{
			Name:      "Osaka, Japan",
			Label:     "asia-northeast1-a",
			Latitude:  "34.6937",
			Longitude: "135.5022",
		},
		{
			Name:      "Osaka, Japan",
			Label:     "asia-northeast1-b",
			Latitude:  "34.6937",
			Longitude: "135.5022",
		},
		{
			Name:      "Osaka, Japan",
			Label:     "asia-northeast1-c",
			Latitude:  "34.6937",
			Longitude: "135.5022",
		},
		// asia-northeast2
		{
			Name:      "Seoul, South Korea",
			Label:     "asia-northeast2-a",
			Latitude:  "37.5665",
			Longitude: "126.9780",
		},
		{
			Name:      "Seoul, South Korea",
			Label:     "asia-northeast2-b",
			Latitude:  "37.5665",
			Longitude: "126.9780",
		},
		{
			Name:      "Seoul, South Korea",
			Label:     "asia-northeast2-c",
			Latitude:  "37.5665",
			Longitude: "126.9780",
		},
		// asia-southeast1
		{
			Name:      "Jurong West, Singapore",
			Label:     "asia-southeast1-a",
			Latitude:  "1.2931",
			Longitude: "103.8558",
		},
		{
			Name:      "Jurong West, Singapore",
			Label:     "asia-southeast1-b",
			Latitude:  "1.2931",
			Longitude: "103.8558",
		},
		{
			Name:      "Jurong West, Singapore",
			Label:     "asia-southeast1-c",
			Latitude:  "1.2931",
			Longitude: "103.8558",
		},
		// australia-southeast1
		{
			Name:      "Sydney, Australia",
			Label:     "australia-southeast1-a",
			Latitude:  "-33.8688",
			Longitude: "151.2093",
		},
		{
			Name:      "Sydney, Australia",
			Label:     "australia-southeast1-b",
			Latitude:  "-33.8688",
			Longitude: "151.2093",
		},
		{
			Name:      "Sydney, Australia",
			Label:     "australia-southeast1-c",
			Latitude:  "-33.8688",
			Longitude: "151.2093",
		},
		// australia-southeast2
		{
			Name:      "Melbourne, Australia",
			Label:     "australia-southeast2-a",
			Latitude:  "-37.8136",
			Longitude: "144.9631",
		},
		{
			Name:      "Melbourne, Australia",
			Label:     "australia-southeast2-b",
			Latitude:  "-37.8136",
			Longitude: "144.9631",
		},
		{
			Name:      "Melbourne, Australia",
			Label:     "australia-southeast2-c",
			Latitude:  "-37.8136",
			Longitude: "144.9631",
		},
		// us-central1
		{
			Name:      "Council Bluffs, Iowa, USA",
			Label:     "us-central1-a",
			Latitude:  "41.87801",
			Longitude: "-93.0977",
		},
		{
			Name:      "Council Bluffs, Iowa, USA",
			Label:     "us-central1-b",
			Latitude:  "41.87801",
			Longitude: "-93.0977",
		},
		{
			Name:      "Council Bluffs, Iowa, USA",
			Label:     "us-central1-c",
			Latitude:  "41.87801",
			Longitude: "-93.0977",
		},
		{
			Name:      "Council Bluffs, Iowa, USA",
			Label:     "us-central1-f",
			Latitude:  "41.87801",
			Longitude: "-93.0977",
		},
		// us-east1
		{
			Name:      "Moncks Corner, South Carolina, USA",
			Label:     "us-east1-b",
			Latitude:  "33.2009",
			Longitude: "-80.0076",
		},
		{
			Name:      "Moncks Corner, South Carolina, USA",
			Label:     "us-east1-c",
			Latitude:  "33.2009",
			Longitude: "-80.0076",
		},
		{
			Name:      "Moncks Corner, South Carolina, USA",
			Label:     "us-east1-d",
			Latitude:  "33.2009",
			Longitude: "-80.0076",
		},
		// us-east4
		{
			Name:      "South Carolina, USA",
			Label:     "us-east4-a",
			Latitude:  "33.2009",
			Longitude: "-80.0076",
		},
		{
			Name:      "South Carolina, USA",
			Label:     "us-east4-b",
			Latitude:  "33.2009",
			Longitude: "-80.0076",
		},
		{
			Name:      "South Carolina, USA",
			Label:     "us-east4-c",
			Latitude:  "33.2009",
			Longitude: "-80.0076",
		},
		// us-east5
		{
			Name:      "Columbus, Ohio, USA",
			Label:     "us-east5-a",
			Latitude:  "39.9612",
			Longitude: "-82.9988",
		},
		{
			Name:      "Columbus, Ohio, USA",
			Label:     "us-east5-b",
			Latitude:  "39.9612",
			Longitude: "-82.9988",
		},
		{
			Name:      "Columbus, Ohio, USA",
			Label:     "us-east5-c",
			Latitude:  "39.9612",
			Longitude: "-82.9988",
		},
		// us-west1
		{
			Name:      "Los Angeles, California, USA",
			Label:     "us-west1-a",
			Latitude:  "34.0522",
			Longitude: "-118.2437",
		},
		{
			Name:      "Los Angeles, California, USA",
			Label:     "us-west1-b",
			Latitude:  "34.0522",
			Longitude: "-118.2437",
		},
		{
			Name:      "Los Angeles, California, USA",
			Label:     "us-west1-c",
			Latitude:  "34.0522",
			Longitude: "-118.2437",
		},
		// us-west2
		{
			Name:      "Salt Lake City, Utah, USA",
			Label:     "us-west2-a",
			Latitude:  "40.7608",
			Longitude: "-111.8910",
		},
		{
			Name:      "Salt Lake City, Utah, USA",
			Label:     "us-west2-b",
			Latitude:  "40.7608",
			Longitude: "-111.8910",
		},
		{
			Name:      "Salt Lake City, Utah, USA",
			Label:     "us-west2-c",
			Latitude:  "40.7608",
			Longitude: "-111.8910",
		},
		// us-west3
		{
			Name:      "Las Vegas, Nevada, USA",
			Label:     "us-west3-a",
			Latitude:  "36.1699",
			Longitude: "-115.1398",
		},
		{
			Name:      "Las Vegas, Nevada, USA",
			Label:     "us-west3-b",
			Latitude:  "36.1699",
			Longitude: "-115.1398",
		},
		{
			Name:      "Las Vegas, Nevada, USA",
			Label:     "us-west3-c",
			Latitude:  "36.1699",
			Longitude: "-115.1398",
		},
		// us-west4
		{
			Name:      "Oregon, USA",
			Label:     "us-west4-a",
			Latitude:  "44.1419",
			Longitude: "-120.5381",
		},
		{
			Name:      "Oregon, USA",
			Label:     "us-west4-b",
			Latitude:  "44.1419",
			Longitude: "-120.5381",
		},
		{
			Name:      "Oregon, USA",
			Label:     "us-west4-c",
			Latitude:  "44.1419",
			Longitude: "-120.5381",
		},
		// southamerica-east1
		{
			Name:      "São Paulo, Brazil",
			Label:     "southamerica-east1-a",
			Latitude:  "-23.5505",
			Longitude: "-46.6333",
		},
		{
			Name:      "São Paulo, Brazil",
			Label:     "southamerica-east1-b",
			Latitude:  "-23.5505",
			Longitude: "-46.6333",
		},
		{
			Name:      "São Paulo, Brazil",
			Label:     "southamerica-east1-c",
			Latitude:  "-23.5505",
			Longitude: "-46.6333",
		},
		// southamerica-west1
		{
			Name:      "Santiago, Chile, South America",
			Label:     "southamerica-west1-a",
			Latitude:  "-33.4489",
			Longitude: "-70.6693",
		},
		{
			Name:      "Santiago, Chile, South America",
			Label:     "southamerica-west1-b",
			Latitude:  "-33.4489",
			Longitude: "-70.6693",
		},
		{
			Name:      "Santiago, Chile, South America",
			Label:     "southamerica-west1-c",
			Latitude:  "-33.4489",
			Longitude: "-70.6693",
		},
		// us-south1
		{
			Name:      "Dallas, Texas, USA",
			Label:     "us-south1-a",
			Latitude:  "32.7767",
			Longitude: "-96.7970",
		},
		{
			Name:      "Dallas, Texas, USA",
			Label:     "us-south1-b",
			Latitude:  "32.7767",
			Longitude: "-96.7970",
		},
		{
			Name:      "Dallas, Texas, USA",
			Label:     "us-south1-c",
			Latitude:  "32.7767",
			Longitude: "-96.7970",
		},
		// northamerica-northeast1
		{
			Name:      "Montréal, Canada",
			Label:     "northamerica-northeast1-a",
			Latitude:  "45.5017",
			Longitude: "-73.5673",
		},
		{
			Name:      "Montréal, Canada",
			Label:     "northamerica-northeast1-b",
			Latitude:  "45.5017",
			Longitude: "-73.5673",
		},
		{
			Name:      "Montréal, Canada",
			Label:     "northamerica-northeast1-c",
			Latitude:  "45.5017",
			Longitude: "-73.5673",
		},
		// northamerica-northeast2
		{
			Name:      "Toronto, Canada",
			Label:     "northamerica-northeast2-a",
			Latitude:  "43.6532",
			Longitude: "-79.3832",
		},
		{
			Name:      "Toronto, Canada",
			Label:     "northamerica-northeast2-b",
			Latitude:  "43.6532",
			Longitude: "-79.3832",
		},
		{
			Name:      "Toronto, Canada",
			Label:     "northamerica-northeast2-c",
			Latitude:  "43.6532",
			Longitude: "-79.3832",
		},
		// me-west1
		{
			Name:      "Tel Aviv, Israel",
			Label:     "me-west1-a",
			Latitude:  "32.0853",
			Longitude: "34.7818",
		},
		{
			Name:      "Tel Aviv, Israel",
			Label:     "me-west1-b",
			Latitude:  "32.0853",
			Longitude: "34.7818",
		},
		{
			Name:      "Tel Aviv, Israel",
			Label:     "me-west1-c",
			Latitude:  "32.0853",
			Longitude: "34.7818",
		},
		// europe-north1
		{
			Name:      "Hamina, Finland",
			Label:     "europe-north1-a",
			Latitude:  "60.5682",
			Longitude: "27.1990",
		},
		{
			Name:      "Hamina, Finland",
			Label:     "europe-north1-b",
			Latitude:  "60.5682",
			Longitude: "27.1990",
		},
		{
			Name:      "Hamina, Finland",
			Label:     "europe-north1-c",
			Latitude:  "60.5682",
			Longitude: "27.1990",
		},
		// europe-central2
		{
			Name:      "Warsaw, Poland",
			Label:     "europe-central2-a",
			Latitude:  "52.2297",
			Longitude: "21.0122",
		},

		{
			Name:      "Warsaw, Poland",
			Label:     "europe-central2-b",
			Latitude:  "52.2297",
			Longitude: "21.0122",
		},
		{
			Name:      "Warsaw, Poland",
			Label:     "europe-central2-c",
			Latitude:  "52.2297",
			Longitude: "21.0122",
		},
		// europe-southwest1
		{
			Name:      "Madrid, Spain",
			Label:     "europe-southwest1-a",
			Latitude:  "40.4168",
			Longitude: "-3.7038",
		},
		{
			Name:      "Madrid, Spain",
			Label:     "europe-southwest1-b",
			Latitude:  "40.4168",
			Longitude: "-3.7038",
		},
		{
			Name:      "Madrid, Spain",
			Label:     "europe-southwest1-c",
			Latitude:  "40.4168",
			Longitude: "-3.7038",
		},
		// europe-west1
		{
			Name:      "St. Ghislain, Belgium",
			Label:     "europe-west1-b",
			Latitude:  "50.6333",
			Longitude: "3.0667",
		},
		{
			Name:      "St. Ghislain, Belgium",
			Label:     "europe-west1-c",
			Latitude:  "50.6333",
			Longitude: "3.0667",
		},
		{
			Name:      "St. Ghislain, Belgium",
			Label:     "europe-west1-d",
			Latitude:  "50.6333",
			Longitude: "3.0667",
		},
		// europe-west2
		{
			Name:      "London, England",
			Label:     "europe-west2-a",
			Latitude:  "51.5074",
			Longitude: "-0.1278",
		},
		{
			Name:      "London, England",
			Label:     "europe-west2-b",
			Latitude:  "51.5074",
			Longitude: "-0.1278",
		},
		{
			Name:      "London, England",
			Label:     "europe-west2-c",
			Latitude:  "51.5074",
			Longitude: "-0.1278",
		},
		// europe-west3
		{
			Name:      "Frankfurt, Germany",
			Label:     "europe-west3-a",
			Latitude:  "50.1109",
			Longitude: "8.6821",
		},
		{
			Name:      "Frankfurt, Germany",
			Label:     "europe-west3-b",
			Latitude:  "50.1109",
			Longitude: "8.6821",
		},
		{
			Name:      "Frankfurt, Germany",
			Label:     "europe-west3-c",
			Latitude:  "50.1109",
			Longitude: "8.6821",
		},
		// europe-west4
		{
			Name:      "Eemshaven, Netherlands",
			Label:     "europe-west4-a",
			Latitude:  "53.4326",
			Longitude: "6.8586",
		},
		{
			Name:      "Eemshaven, Netherlands",
			Label:     "europe-west4-b",
			Latitude:  "53.4326",
			Longitude: "6.8586",
		},
		{
			Name:      "Eemshaven, Netherlands",
			Label:     "europe-west4-c",
			Latitude:  "53.4326",
			Longitude: "6.8586",
		},
		// europe-west6
		{
			Name:      "Zürich, Switzerland",
			Label:     "europe-west6-a",
			Latitude:  "47.3769",
			Longitude: "8.5417",
		},
		{
			Name:      "Zürich, Switzerland",
			Label:     "europe-west6-b",
			Latitude:  "47.3769",
			Longitude: "8.5417",
		},
		{
			Name:      "Zürich, Switzerland",
			Label:     "europe-west6-c",
			Latitude:  "47.3769",
			Longitude: "8.5417",
		},
		// europe-west8
		{
			Name:      "Amsterdam, Netherlands",
			Label:     "europe-west8-a",
			Latitude:  "52.3667",
			Longitude: "4.8945",
		},
		{
			Name:      "Amsterdam, Netherlands",
			Label:     "europe-west8-b",
			Latitude:  "52.3667",
			Longitude: "4.8945",
		},
		{
			Name:      "Amsterdam, Netherlands",
			Label:     "europe-west8-c",
			Latitude:  "52.3667",
			Longitude: "4.8945",
		},
		// europe-west9
		{
			Name:      "Paris, France",
			Label:     "europe-west9-a",
			Latitude:  "48.8566",
			Longitude: "2.3522",
		},
		{
			Name:      "Paris, France",
			Label:     "europe-west9-b",
			Latitude:  "48.8566",
			Longitude: "2.3522",
		},
		{
			Name:      "Paris, France",
			Label:     "europe-west9-c",
			Latitude:  "48.8566",
			Longitude: "2.3522",
		},
	}
}

func getGCPRegions() []Location {
	return []Location{
		{
			Name:      "Ashburn",
			Label:     "us-east1",
			Latitude:  "39.04372",
			Longitude: "-77.48749",
		},
		{
			Name:      "Moncks Corner",
			Label:     "us-east4",
			Latitude:  "33.20087",
			Longitude: "-80.00756",
		},
		{
			Name:      "Los Angeles",
			Label:     "us-west2",
			Latitude:  "34.05223",
			Longitude: "-118.24368",
		},
		{
			Name:      "Salt Lake City",
			Label:     "us-west3",
			Latitude:  "40.76078",
			Longitude: "-111.89105",
		},
		{
			Name:      "Las Vegas",
			Label:     "us-west4",
			Latitude:  "36.16994",
			Longitude: "-115.13983",
		},
		{
			Name:      "South Carolina",
			Label:     "us-central1",
			Latitude:  "34.00071",
			Longitude: "-81.03481",
		},
		{
			Name:      "Iowa",
			Label:     "us-central2",
			Latitude:  "41.87801",
			Longitude: "-93.0977",
		},
		{
			Name:      "Oregon",
			Label:     "us-west1",
			Latitude:  "45.52345",
			Longitude: "-122.67621",
		},
		{
			Name:      "South Carolina",
			Label:     "us-east2",
			Latitude:  "34.00071",
			Longitude: "-81.03481",
		},
		{
			Name:      "South Carolina",
			Label:     "us-east3",
			Latitude:  "34.00071",
			Longitude: "-81.03481",
		},
		{
			Name:      "Finland",
			Label:     "europe-north1",
			Latitude:  "60.16952",
			Longitude: "24.93838",
		},
		{
			Name:      "Belgium",
			Label:     "europe-west1",
			Latitude:  "50.85034",
			Longitude: "4.35171",
		},
		// Regions below are used for multi-region buckets
		{
			Name:      "EU",
			Label:     "eu",
			Latitude:  "47.751569",
			Longitude: "1.675063",
		},
		{
			Name:      "US",
			Label:     "us",
			Latitude:  "44.967243",
			Longitude: "-103.771556",
		},
		{
			Name:      "ASIA",
			Label:     "asia",
			Latitude:  "52.483333",
			Longitude: "96.085833",
		},
	}
}

func GetLocationFromRegion(label string) Location {
	awsRegions := getAWSRegions()
	doRegions := getDigitalOceanRegions()
	gcpRegions := getGCPRegions()
	gcpZones := getGCPZones()

	for _, region := range awsRegions {
		if region.Label == label {
			return region
		}
	}

	for _, region := range doRegions {
		if region.Label == label {
			return region
		}
	}

	for _, region := range gcpRegions {
		if region.Label == label {
			return region
		}
	}

	for _, zone := range gcpZones {
		if zone.Label == label {
			return zone
		}
	}

	return Location{}
}

// MongoDB Atlas returns the region names of AWS, GCP and Azure.
// The names are written as "EU_CENTRAL_1" instead of "eu-central-1", which
// this function fixes.
func NormalizeRegionName(regionName string) string {
	lowercased := strings.ToLower(regionName)
	dashReplaced := strings.Replace(lowercased, "_", "-", -1)
	return dashReplaced
}

func GcpGetRegionFromZone(zone string) string {
	p := strings.Split(zone, "-")
	return strings.Join(p[:len(p)-1], "-")
}

func GcpExtractZoneFromURL(url string) string {
	return url[strings.LastIndex(url, "/")+1:]
}

// Fetch all the available GCP Regions for the given ProjectID
func FetchGCPRegionsInRealtime(projectId string, creds option.ClientOption) ([]string, error) {
	var regions []string

	ctx := context.Background()
	computeService, err := compute.NewService(ctx, creds)
	if err != nil {
		logrus.WithError(err).Errorf("failed to fetch GCP regions")
		return nil, err
	}

	regionList, err := computeService.Regions.List(projectId).Do()
	if err != nil {
		logrus.WithError(err).Errorf("failed to list GCP regions")
		return nil, err
	}

	for _, region := range regionList.Items {
		regions = append(regions, region.Name)
	}
	return regions, nil
}
