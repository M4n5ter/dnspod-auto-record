package cloudflare_test

import (
	"dyndns/cloudflare"
	"testing"
)

func TestNewApiUrl(t *testing.T) {
	api := cloudflare.newApiUrl("123", "456")
	if api != "https://api.cloudflare.com/client/v4/zones/123/dns_records/456" {
		t.Errorf("api url error: %v", api)
	}

	api = cloudflare.newApiUrl("123", "")
	if api != "https://api.cloudflare.com/client/v4/zones/123/dns_records" {
		t.Errorf("api url error: %v", api)
	}

}
