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

// RefHotspotGeoService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefHotspotGeoService] method instead.
type RefHotspotGeoService struct {
	Options []option.RequestOption
}

// NewRefHotspotGeoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRefHotspotGeoService(opts ...option.RequestOption) (r *RefHotspotGeoService) {
	r = &RefHotspotGeoService{}
	r.Options = opts
	return
}

// Get the list of hotspots, within a radius of up to 50 kilometers, from a given
// set of coordinates.
func (r *RefHotspotGeoService) Get(ctx context.Context, query RefHotspotGeoGetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ref/hotspot/geo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type RefHotspotGeoGetParams struct {
	Lat param.Field[float64] `query:"lat,required"`
	Lng param.Field[float64] `query:"lng,required"`
	// The number of days back to fetch hotspots.
	Back param.Field[int64] `query:"back"`
	// The search radius from the given position, in kilometers.
	Dist param.Field[int64] `query:"dist"`
	// Fetch the records in CSV or JSON format.
	Fmt param.Field[RefHotspotGeoGetParamsFmt] `query:"fmt"`
}

// URLQuery serializes [RefHotspotGeoGetParams]'s query parameters as `url.Values`.
func (r RefHotspotGeoGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Fetch the records in CSV or JSON format.
type RefHotspotGeoGetParamsFmt string

const (
	RefHotspotGeoGetParamsFmtCsv  RefHotspotGeoGetParamsFmt = "csv"
	RefHotspotGeoGetParamsFmtJson RefHotspotGeoGetParamsFmt = "json"
)

func (r RefHotspotGeoGetParamsFmt) IsKnown() bool {
	switch r {
	case RefHotspotGeoGetParamsFmtCsv, RefHotspotGeoGetParamsFmtJson:
		return true
	}
	return false
}
