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

func TestDataObservationGeoRecentNotableListWithOptionalParams(t *testing.T) {
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
	_, err := client.Data.Observations.Geo.Recent.Notable.List(context.TODO(), phoebe.DataObservationGeoRecentNotableListParams{
		Lat:        phoebe.F(-90.000000),
		Lng:        phoebe.F(-180.000000),
		Back:       phoebe.F(int64(1)),
		Detail:     phoebe.F(phoebe.DataObservationGeoRecentNotableListParamsDetailSimple),
		Dist:       phoebe.F(int64(0)),
		Hotspot:    phoebe.F(true),
		MaxResults: phoebe.F(int64(1)),
		SppLocale:  phoebe.F("sppLocale"),
	})
	if err != nil {
		var apierr *phoebe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
