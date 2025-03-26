// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe_test

import (
	"context"
	"os"
	"testing"

	"github.com/phoebe-bird/phoebe-go"
	"github.com/phoebe-bird/phoebe-go/internal/testutil"
	"github.com/phoebe-bird/phoebe-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := phoebe.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	info, err := client.Ref.Hotspot.Info.Get(context.TODO(), "L99381")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", info.CountryCode)
}
