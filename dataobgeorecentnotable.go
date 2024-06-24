// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/phoebe-go/internal/apiquery"
	"github.com/stainless-sdks/phoebe-go/internal/param"
	"github.com/stainless-sdks/phoebe-go/internal/requestconfig"
	"github.com/stainless-sdks/phoebe-go/option"
)

// DataObGeoRecentNotableService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObGeoRecentNotableService] method instead.
type DataObGeoRecentNotableService struct {
	Options []option.RequestOption
}

// NewDataObGeoRecentNotableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDataObGeoRecentNotableService(opts ...option.RequestOption) (r *DataObGeoRecentNotableService) {
	r = &DataObGeoRecentNotableService{}
	r.Options = opts
	return
}

// Get the list of notable observations (up to 30 days ago) of birds seen at
// locations within a radius of up to 50 kilometers, from a given set of
// coordinates. Notable observations can be for locally or nationally rare species
// or are otherwise unusual, for example over-wintering birds in a species which is
// normally only a summer visitor.
func (r *DataObGeoRecentNotableService) List(ctx context.Context, query DataObGeoRecentNotableListParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "data/obs/geo/recent/notable"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type DataObGeoRecentNotableListParams struct {
	Lat param.Field[float64] `query:"lat,required"`
	Lng param.Field[float64] `query:"lng,required"`
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Include a subset (simple), or all (full), of the fields available.
	Detail param.Field[DataObGeoRecentNotableListParamsDetail] `query:"detail"`
	// The search radius from the given position, in kilometers.
	Dist param.Field[int64] `query:"dist"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObGeoRecentNotableListParams]'s query parameters as
// `url.Values`.
func (r DataObGeoRecentNotableListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Include a subset (simple), or all (full), of the fields available.
type DataObGeoRecentNotableListParamsDetail string

const (
	DataObGeoRecentNotableListParamsDetailSimple DataObGeoRecentNotableListParamsDetail = "simple"
	DataObGeoRecentNotableListParamsDetailFull   DataObGeoRecentNotableListParamsDetail = "full"
)

func (r DataObGeoRecentNotableListParamsDetail) IsKnown() bool {
	switch r {
	case DataObGeoRecentNotableListParamsDetailSimple, DataObGeoRecentNotableListParamsDetailFull:
		return true
	}
	return false
}
