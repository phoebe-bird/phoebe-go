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

func TestDataObservationGeoRecentNotableListWithOptionalParams(t *testing.T) {
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
	_, err := client.Data.Observations.Geo.Recent.Notable.List(context.TODO(), phoebebird.DataObservationGeoRecentNotableListParams{
		Lat:        phoebebird.F(-90.000000),
		Lng:        phoebebird.F(-180.000000),
		Back:       phoebebird.F(int64(1)),
		Detail:     phoebebird.F(phoebebird.DataObservationGeoRecentNotableListParamsDetailSimple),
		Dist:       phoebebird.F(int64(0)),
		Hotspot:    phoebebird.F(true),
		MaxResults: phoebebird.F(int64(1)),
		SppLocale:  phoebebird.F("string"),
	})
	if err != nil {
		var apierr *phoebebird.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
