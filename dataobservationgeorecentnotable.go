// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/phoebe-bird/phoebe-go/internal/apiquery"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// DataObservationGeoRecentNotableService contains methods and other services that
// help with interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationGeoRecentNotableService] method instead.
type DataObservationGeoRecentNotableService struct {
	Options []option.RequestOption
}

// NewDataObservationGeoRecentNotableService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDataObservationGeoRecentNotableService(opts ...option.RequestOption) (r *DataObservationGeoRecentNotableService) {
	r = &DataObservationGeoRecentNotableService{}
	r.Options = opts
	return
}

// Get the list of notable observations (up to 30 days ago) of birds seen at
// locations within a radius of up to 50 kilometers, from a given set of
// coordinates. Notable observations can be for locally or nationally rare species
// or are otherwise unusual, for example over-wintering birds in a species which is
// normally only a summer visitor.
func (r *DataObservationGeoRecentNotableService) List(ctx context.Context, query DataObservationGeoRecentNotableListParams, opts ...option.RequestOption) (res *[]Observation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "data/obs/geo/recent/notable"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObservationGeoRecentNotableListParams struct {
	Lat param.Field[float64] `query:"lat" api:"required"`
	Lng param.Field[float64] `query:"lng" api:"required"`
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Include a subset (simple), or all (full), of the fields available.
	Detail param.Field[DataObservationGeoRecentNotableListParamsDetail] `query:"detail"`
	// The search radius from the given position, in kilometers.
	Dist param.Field[int64] `query:"dist"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObservationGeoRecentNotableListParams]'s query
// parameters as `url.Values`.
func (r DataObservationGeoRecentNotableListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Include a subset (simple), or all (full), of the fields available.
type DataObservationGeoRecentNotableListParamsDetail string

const (
	DataObservationGeoRecentNotableListParamsDetailSimple DataObservationGeoRecentNotableListParamsDetail = "simple"
	DataObservationGeoRecentNotableListParamsDetailFull   DataObservationGeoRecentNotableListParamsDetail = "full"
)

func (r DataObservationGeoRecentNotableListParamsDetail) IsKnown() bool {
	switch r {
	case DataObservationGeoRecentNotableListParamsDetailSimple, DataObservationGeoRecentNotableListParamsDetailFull:
		return true
	}
	return false
}
