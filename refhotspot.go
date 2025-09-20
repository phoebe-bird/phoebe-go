// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/phoebe-bird/phoebe-go/internal/apijson"
	"github.com/phoebe-bird/phoebe-go/internal/apiquery"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefHotspotService contains methods and other services that help with interacting
// with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefHotspotService] method instead.
type RefHotspotService struct {
	Options []option.RequestOption
	Geo     *RefHotspotGeoService
	Info    *RefHotspotInfoService
}

// NewRefHotspotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRefHotspotService(opts ...option.RequestOption) (r *RefHotspotService) {
	r = &RefHotspotService{}
	r.Options = opts
	r.Geo = NewRefHotspotGeoService(opts...)
	r.Info = NewRefHotspotInfoService(opts...)
	return
}

// Hotspots in a region
func (r *RefHotspotService) List(ctx context.Context, regionCode string, query RefHotspotListParams, opts ...option.RequestOption) (res *[]RefHotspotListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("ref/hotspot/%s", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RefHotspotListResponse struct {
	CountryCode       string                     `json:"countryCode"`
	Lat               float64                    `json:"lat"`
	LatestObsDt       string                     `json:"latestObsDt"`
	Lng               float64                    `json:"lng"`
	LocID             string                     `json:"locId"`
	LocName           string                     `json:"locName"`
	NumSpeciesAllTime int64                      `json:"numSpeciesAllTime"`
	Subnational1Code  string                     `json:"subnational1Code"`
	Subnational2Code  string                     `json:"subnational2Code"`
	JSON              refHotspotListResponseJSON `json:"-"`
}

// refHotspotListResponseJSON contains the JSON metadata for the struct
// [RefHotspotListResponse]
type refHotspotListResponseJSON struct {
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

func (r *RefHotspotListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refHotspotListResponseJSON) RawJSON() string {
	return r.raw
}

type RefHotspotListParams struct {
	// The number of days back to fetch hotspots.
	Back param.Field[int64] `query:"back"`
	// Fetch the records in CSV or JSON format.
	Fmt param.Field[RefHotspotListParamsFmt] `query:"fmt"`
}

// URLQuery serializes [RefHotspotListParams]'s query parameters as `url.Values`.
func (r RefHotspotListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Fetch the records in CSV or JSON format.
type RefHotspotListParamsFmt string

const (
	RefHotspotListParamsFmtCsv  RefHotspotListParamsFmt = "csv"
	RefHotspotListParamsFmtJson RefHotspotListParamsFmt = "json"
)

func (r RefHotspotListParamsFmt) IsKnown() bool {
	switch r {
	case RefHotspotListParamsFmtCsv, RefHotspotListParamsFmtJson:
		return true
	}
	return false
}
