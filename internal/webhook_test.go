package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlackAttachment(t *testing.T) {
	// default - no slack host specified in config
	attachment := createSlackAttachment("testViewName", 3000, 101, 99, 99.99, "", "RESOURCES")
	expectedActionUrl := "http://localhost:3000/inventory?view=101"
	assert.Equal(t, expectedActionUrl, attachment.Actions[0].URL)

	// explicit - slack host specified in config
	attachment = createSlackAttachment("testViewName", 3000, 101, 99, 99.99, "https://example.com", "RESOURCES")
	expectedActionUrl = "https://example.com/inventory?view=101"
	assert.Equal(t, expectedActionUrl, attachment.Actions[0].URL)
}
