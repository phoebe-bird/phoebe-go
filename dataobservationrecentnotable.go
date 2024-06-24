// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/phoebe-bird/phoebe-go/internal/apiquery"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// DataObservationRecentNotableService contains methods and other services that
// help with interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationRecentNotableService] method instead.
type DataObservationRecentNotableService struct {
	Options []option.RequestOption
}

// NewDataObservationRecentNotableService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDataObservationRecentNotableService(opts ...option.RequestOption) (r *DataObservationRecentNotableService) {
	r = &DataObservationRecentNotableService{}
	r.Options = opts
	return
}

// Get the list of recent, notable observations (up to 30 days ago) of birds seen
// in a country, region or location. Notable observations can be for locally or
// nationally rare species or are otherwise unusual, e.g. over-wintering birds in a
// species which is normally only a summer visitor.
func (r *DataObservationRecentNotableService) List(ctx context.Context, regionCode string, query DataObservationRecentNotableListParams, opts ...option.RequestOption) (res *[]Observation, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("data/obs/%s/recent/notable", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObservationRecentNotableListParams struct {
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Include a subset (simple), or all (full), of the fields available.
	Detail param.Field[DataObservationRecentNotableListParamsDetail] `query:"detail"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Fetch observations from up to 10 locations
	R param.Field[[]string] `query:"r"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObservationRecentNotableListParams]'s query parameters
// as `url.Values`.
func (r DataObservationRecentNotableListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Include a subset (simple), or all (full), of the fields available.
type DataObservationRecentNotableListParamsDetail string

const (
	DataObservationRecentNotableListParamsDetailSimple DataObservationRecentNotableListParamsDetail = "simple"
	DataObservationRecentNotableListParamsDetailFull   DataObservationRecentNotableListParamsDetail = "full"
)

func (r DataObservationRecentNotableListParamsDetail) IsKnown() bool {
	switch r {
	case DataObservationRecentNotableListParamsDetailSimple, DataObservationRecentNotableListParamsDetailFull:
		return true
	}
	return false
}
