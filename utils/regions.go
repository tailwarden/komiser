package utils

type Location struct {
	Name      string `json:"name"`
	Label     string `json:"label"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func getAWSRegions() []Location {
	return []Location{
		Location{
			Name:      "Ohio",
			Label:     "us-east-2",
			Latitude:  "40.367474",
			Longitude: "-82.996216",
		},
		Location{
			Name:      "N.Virginia",
			Label:     "us-east-1",
			Latitude:  "37.926868",
			Longitude: "-78.024902",
		},
		Location{
			Name:      "N.California",
			Label:     "us-west-1",
			Latitude:  "36.778261",
			Longitude: "-119.4179324",
		},
		Location{
			Name:      "Oregon",
			Label:     "us-west-2",
			Latitude:  "45.523062",
			Longitude: "-122.676482",
		},
		Location{
			Name:      "Cape Town",
			Label:     "af-south-1",
			Latitude:  "-33.924869",
			Longitude: "18.424055",
		},
		Location{
			Name:      "Hong Kong",
			Label:     "ap-east-1",
			Latitude:  "22.302711",
			Longitude: "114.177216",
		},
		Location{
			Name:      "Jakarta",
			Label:     "ap-southeast-3",
			Latitude:  "-6.2087634",
			Longitude: "106.816666",
		},
		Location{
			Name:      "Mumbai",
			Label:     "ap-south-1",
			Latitude:  "19.076090",
			Longitude: "72.877426",
		},
		Location{
			Name:      "Osaka",
			Label:     "ap-northeast-3",
			Latitude:  "34.6937378",
			Longitude: "135.5021651",
		},
		Location{
			Name:      "Seoul",
			Label:     "ap-northeast-2",
			Latitude:  "37.566535",
			Longitude: "126.9779692",
		},
		Location{
			Name:      "Singapore",
			Label:     "ap-southeast-1",
			Latitude:  "1.290270",
			Longitude: "103.851959",
		},
		Location{
			Name:      "Sydney",
			Label:     "ap-southeast-2",
			Latitude:  "-33.8667",
			Longitude: "151.206990",
		},
		Location{
			Name:      "Tokyo",
			Label:     "ap-northeast-1",
			Latitude:  "35.652832",
			Longitude: "139.839478",
		},
		Location{
			Name:      "Canada",
			Label:     "ca-central-1",
			Latitude:  "56.130367",
			Longitude: "-106.346771",
		},
		Location{
			Name:      "Frankfurt",
			Label:     "eu-central-1",
			Latitude:  "50.1109221",
			Longitude: "8.6821267",
		},
		Location{
			Name:      "Ireland",
			Label:     "eu-west-1",
			Latitude:  "53.350140",
			Longitude: "-6.266155",
		},
		Location{
			Name:      "London",
			Label:     "eu-west-2",
			Latitude:  "51.5073509",
			Longitude: "-0.1277583",
		},
		Location{
			Name:      "Milan",
			Label:     "eu-south-1",
			Latitude:  "45.4654219",
			Longitude: "9.1859243",
		},
		Location{
			Name:      "Paris",
			Label:     "eu-west-3",
			Latitude:  "48.864716",
			Longitude: "2.352222",
		},
		Location{
			Name:      "Stockholm",
			Label:     "eu-north-1",
			Latitude:  "59.334591",
			Longitude: "18.063240",
		},
		Location{
			Name:      "Bahrain",
			Label:     "me-south-1",
			Latitude:  "26.066700",
			Longitude: "50.557700",
		},
	}
}

func getDigitalOceanRegions() []Location {
	return []Location{
		Location{
			Name:      "New York City",
			Label:     "NYC1",
			Latitude:  "40.712776",
			Longitude: "-74.005974",
		},
		Location{
			Name:      "New York City",
			Label:     "NYC2",
			Latitude:  "40.712776",
			Longitude: "-74.005974",
		},
		Location{
			Name:      "New York City",
			Label:     "NYC3",
			Latitude:  "40.712776",
			Longitude: "-74.005974",
		},
		Location{
			Name:      "San Francisco",
			Label:     "SFO1",
			Latitude:  "37.774929",
			Longitude: "-122.419418",
		},
		Location{
			Name:      "San Francisco",
			Label:     "SFO2",
			Latitude:  "37.774929",
			Longitude: "-122.419418",
		},
		Location{
			Name:      "Toronto",
			Label:     "TOR1",
			Latitude:  "43.653225",
			Longitude: "-79.383186",
		},
		Location{
			Name:      "London",
			Label:     "LON1",
			Latitude:  "51.507351",
			Longitude: "-0.127758",
		},
		Location{
			Name:      "Frankfurt",
			Label:     "FRA1",
			Latitude:  "50.110924",
			Longitude: "8.682127",
		},
		Location{
			Name:      "Amsterdam",
			Label:     "AMS2",
			Latitude:  "52.377956",
			Longitude: "4.897070",
		},
		Location{
			Name:      "Amsterdam",
			Label:     "AMS3",
			Latitude:  "52.377956",
			Longitude: "4.897070",
		},
		Location{
			Name:      "Singapore",
			Label:     "SGP1",
			Latitude:  "1.290270",
			Longitude: "103.851959",
		},
		Location{
			Name:      "Bangalore",
			Label:     "BLR1",
			Latitude:  "12.972442",
			Longitude: "77.580643",
		},
	}
}

func GetLocationFromRegion(label string) Location {
	awsRegions := getAWSRegions()
	doRegions := getDigitalOceanRegions()

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

	return Location{}
}
