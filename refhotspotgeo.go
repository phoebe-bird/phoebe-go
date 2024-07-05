// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebebird

import (
	"context"
	"net/http"
	"net/url"

	"github.com/phoebe-bird/phoebe-go/internal/apijson"
	"github.com/phoebe-bird/phoebe-go/internal/apiquery"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
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
func (r *RefHotspotGeoService) Get(ctx context.Context, query RefHotspotGeoGetParams, opts ...option.RequestOption) (res *[]RefHotspotGeoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ref/hotspot/geo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RefHotspotGeoGetResponse struct {
	CountryCode       string                       `json:"countryCode"`
	Lat               float64                      `json:"lat"`
	LatestObsDt       string                       `json:"latestObsDt"`
	Lng               float64                      `json:"lng"`
	LocID             string                       `json:"locId"`
	LocName           string                       `json:"locName"`
	NumSpeciesAllTime int64                        `json:"numSpeciesAllTime"`
	Subnational1Code  string                       `json:"subnational1Code"`
	Subnational2Code  string                       `json:"subnational2Code"`
	JSON              refHotspotGeoGetResponseJSON `json:"-"`
}

// refHotspotGeoGetResponseJSON contains the JSON metadata for the struct
// [RefHotspotGeoGetResponse]
type refHotspotGeoGetResponseJSON struct {
	CountryCode       apijson.Field
	Lat               apijson.Field
	LatestObsDt       apijson.Field
	Lng               apijson.Field
	LocID             apijson.Field
	LocName           apijson.Field
	NumSpeciesAllTime apijson.Field
	Subnational1Code  apijson.Field
	Subnational2Code  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RefHotspotGeoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refHotspotGeoGetResponseJSON) RawJSON() string {
	return r.raw
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
