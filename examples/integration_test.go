// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main_test

import (
	"context"
	"errors"
	"github.com/stainless-sdks/phoebe-go"
	"github.com/stainless-sdks/phoebe-go/option"
	"os"
	"testing"
)

func TestDataObservationGeoRecentList(t *testing.T) {
	client := getTestClient(t)
	_, err := client.Data.Observations.Geo.Recent.List(context.TODO(), phoebe.DataObservationGeoRecentListParams{
		Lat: phoebe.F(40.77),
		Lng: phoebe.F(73.96),
	})
	if err != nil {
		reportError(t, err)
	}
}

func TestDataObservationRecentList(t *testing.T) {
	client := getTestClient(t)
	_, err := client.Data.Observations.Recent.List(context.TODO(), "KZ", phoebe.DataObservationRecentListParams{})
	if err != nil {
		reportError(t, err)
	}
}

func getTestClient(t *testing.T) *phoebe.Client {
	if apiKey, ok := os.LookupEnv("EBIRD_API_KEY"); !ok {
		t.Skipf("No API Key found in env: EBIRD_API_KEY. Skipping test.")
		return nil
	} else {
		return phoebe.NewClient(
			option.WithAPIKey(apiKey),
		)
	}
}

func reportError(t *testing.T, err error) {
	var apierr *phoebe.Error
	if errors.As(err, &apierr) {
		t.Log(string(apierr.DumpRequest(true)))
	}
	t.Fatalf("err should be nil: %s", err.Error())
}
