package dynamodb

import (
	"bytes"
	"hash/crc32"
	"io"
	"io/ioutil"
	"math"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	client "github.com/aws/aws-sdk-go-v2/aws"
	request "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
)

type retryer struct {
	client.DefaultRetryer
}

func (d retryer) RetryRules(r *request.Request) time.Duration {
	delay := time.Duration(math.Pow(2, float64(r.RetryCount))) * 50
	return delay * time.Millisecond
}

func init() {
	initClient = func(c *DynamoDB) {
		if c.Config.Retryer == nil {
			// Only override the retryer with a custom one if the config
			// does not already contain a retryer
			setCustomRetryer(c)
		}

		c.Handlers.Build.PushBackNamed(disableCompressionHandler)
		c.Handlers.Unmarshal.PushFrontNamed(validateCRC32Handler)
	}

	initRequest = func(c *DynamoDB, req *aws.Request) {
		if c.DisableComputeChecksums {
			// Checksum validation is off, remove the validator.
			req.Handlers.Unmarshal.Remove(validateCRC32Handler)
		}
	}
}

func setCustomRetryer(c *DynamoDB) {
	c.Retryer = retryer{
		DefaultRetryer: client.DefaultRetryer{
			NumMaxRetries: 10,
		},
	}
}

func drainBody(b io.ReadCloser, length int64) (out *bytes.Buffer, err error) {
	if length < 0 {
		length = 0
	}
	buf := bytes.NewBuffer(make([]byte, 0, length))

	if _, err = buf.ReadFrom(b); err != nil {
		return nil, err
	}
	if err = b.Close(); err != nil {
		return nil, err
	}
	return buf, nil
}

var disableCompressionHandler = aws.NamedHandler{Name: "dynamodb.DisableCompression", Fn: disableCompression}

func disableCompression(r *request.Request) {
	r.HTTPRequest.Header.Set("Accept-Encoding", "identity")
}

var validateCRC32Handler = aws.NamedHandler{Name: "dynamodb.ValidateCRC32", Fn: validateCRC32}

func validateCRC32(r *request.Request) {
	if r.Error != nil {
		return // already have an error, no need to verify CRC
	}

	// Try to get CRC from response
	header := r.HTTPResponse.Header.Get("X-Amz-Crc32")
	if header == "" {
		return // No header, skip
	}

	expected, err := strconv.ParseUint(header, 10, 32)
	if err != nil {
		return // Could not determine CRC value, skip
	}

	buf, err := drainBody(r.HTTPResponse.Body, r.HTTPResponse.ContentLength)
	if err != nil { // failed to read the response body, skip
		return
	}

	// Reset body for subsequent reads
	r.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))

	// Compute the CRC checksum
	crc := crc32.ChecksumIEEE(buf.Bytes())

	if crc != uint32(expected) {
		// CRC does not match, set a retryable error
		r.Retryable = aws.Bool(true)
		r.Error = awserr.New("CRC32CheckFailed", "CRC32 integrity check failed", nil)
	}
}
