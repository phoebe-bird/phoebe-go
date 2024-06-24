// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/phoebe-go"
	"github.com/stainless-sdks/phoebe-go/internal/testutil"
	"github.com/stainless-sdks/phoebe-go/option"
)

func TestDataObRecentNotableListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := phoebe.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Data.Obs.Recent.Notable.List(
		context.TODO(),
		"string",
		phoebe.DataObRecentNotableListParams{
			Back:       phoebe.F(int64(1)),
			Detail:     phoebe.F(phoebe.DataObRecentNotableListParamsDetailSimple),
			Hotspot:    phoebe.F(true),
			MaxResults: phoebe.F(int64(1)),
			R:          phoebe.F([]string{"string"}),
			SppLocale:  phoebe.F("string"),
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
