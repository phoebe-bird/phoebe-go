// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/phoebe-bird/phoebe-go"
	"github.com/phoebe-bird/phoebe-go/internal/testutil"
	"github.com/phoebe-bird/phoebe-go/option"
)

func TestDataObservationRecentNotableListWithOptionalParams(t *testing.T) {
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
	_, err := client.Data.Observations.Recent.Notable.List(
		context.TODO(),
		"regionCode",
		phoebe.DataObservationRecentNotableListParams{
			Back:       phoebe.F(int64(1)),
			Detail:     phoebe.F(phoebe.DataObservationRecentNotableListParamsDetailSimple),
			Hotspot:    phoebe.F(true),
			MaxResults: phoebe.F(int64(1)),
			R:          phoebe.F([]string{"string"}),
			SppLocale:  phoebe.F("sppLocale"),
		},
	)
	if err != nil {
		var apierr *phoebe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
