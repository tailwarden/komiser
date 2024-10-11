package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlackConfig(t *testing.T) {
	// test when a value is only specified for webhook
	cfgText := `
	[slack]
		webhook = "https://example.com"
	`
	cfgBytes := []byte(cfgText)
	config, _ := loadConfigFromBytes(cfgBytes)

	assert.Equal(t, "https://example.com", config.Slack.Webhook)
	assert.Equal(t, false, config.Slack.Reporting)
	assert.Equal(t, "", config.Slack.Host)

	// test when a value is specified for reporting
	cfgText = `
	[slack]
		webhook = "https://example.com"
		reporting = true
	`
	cfgBytes = []byte(cfgText)
	config, _ = loadConfigFromBytes(cfgBytes)

	assert.Equal(t, "https://example.com", config.Slack.Webhook)
	assert.Equal(t, true, config.Slack.Reporting)
	assert.Equal(t, "", config.Slack.Host)

	// test when a value is specified for host
	cfgText = `
	[slack]
		webhook = "https://example.com"
		reporting = true
		host = "https://example.com/komiser"
	`
	cfgBytes = []byte(cfgText)
	config, _ = loadConfigFromBytes(cfgBytes)

	assert.Equal(t, "https://example.com", config.Slack.Webhook)
	assert.Equal(t, true, config.Slack.Reporting)
	assert.Equal(t, "https://example.com/komiser", config.Slack.Host)

}
