package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeLambdaFunctionsPerRuntime(cfg aws.Config) (map[string]int, error) {
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
			output[lambda.Runtime]++
		}
	}
	return output, nil
}

func (aws AWS) getLambdaFunctions(cfg aws.Config, region string) ([]Lambda, error) {
	cfg.Region = region
	svc := lambda.New(cfg)
	req := svc.ListFunctionsRequest(&lambda.ListFunctionsInput{})
	result, err := req.Send()
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
