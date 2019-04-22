package kinesis

import (
	"time"

	request "github.com/aws/aws-sdk-go-v2/aws"
)

var readDuration = 5 * time.Second

func init() {
	ops := []string{
		opGetRecords,
	}
	initRequest = func(c *Kinesis, r *request.Request) {
		for _, operation := range ops {
			if r.Operation.Name == operation {
				r.ApplyOptions(request.WithResponseReadTimeout(readDuration))
			}
		}
	}
}
