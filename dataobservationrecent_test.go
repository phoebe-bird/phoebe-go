// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebebird_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/phoebe-bird/phoebe-go"
	"github.com/phoebe-bird/phoebe-go/internal/testutil"
	"github.com/phoebe-bird/phoebe-go/option"
)

func TestDataObservationRecentListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := phoebebird.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Data.Observations.Recent.List(
		context.TODO(),
		"string",
		phoebebird.DataObservationRecentListParams{
			Back:               phoebebird.F(int64(1)),
			Cat:                phoebebird.F(phoebebird.DataObservationRecentListParamsCatSpecies),
			Hotspot:            phoebebird.F(true),
			IncludeProvisional: phoebebird.F(true),
			MaxResults:         phoebebird.F(int64(1)),
			R:                  phoebebird.F([]string{"string"}),
			SppLocale:          phoebebird.F("string"),
		},
	)
	if err != nil {
		var apierr *phoebebird.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
