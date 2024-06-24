// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/phoebe-go/internal/apijson"
	"github.com/stainless-sdks/phoebe-go/internal/apiquery"
	"github.com/stainless-sdks/phoebe-go/internal/param"
	"github.com/stainless-sdks/phoebe-go/internal/requestconfig"
	"github.com/stainless-sdks/phoebe-go/option"
)

// DataObGeoRecentSpecieService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObGeoRecentSpecieService] method instead.
type DataObGeoRecentSpecieService struct {
	Options []option.RequestOption
}

// NewDataObGeoRecentSpecieService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDataObGeoRecentSpecieService(opts ...option.RequestOption) (r *DataObGeoRecentSpecieService) {
	r = &DataObGeoRecentSpecieService{}
	r.Options = opts
	return
}

// Get all observations of a species, seen up to 30 days ago, at any location
// within a radius of up to 50 kilometers, from a given set of coordinates. Results
// include only the most recent observation from each location in the region
// specified.
//
// #### URL parameters
//
// | Name        | Description             |
// | ----------- | ----------------------- |
// | speciesCode | The eBird species code. |
//
// #### Notes
//
// The species code is typically a 6-letter code, e.g. horlar for Horned Lark. You
// can get complete set of species code from the GET eBird Taxonomy end-point.
func (r *DataObGeoRecentSpecieService) List(ctx context.Context, speciesCode string, query DataObGeoRecentSpecieListParams, opts ...option.RequestOption) (res *[]DataObGeoRecentSpecieListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if speciesCode == "" {
		err = errors.New("missing required speciesCode parameter")
		return
	}
	path := fmt.Sprintf("data/obs/geo/recent/%s", speciesCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObGeoRecentSpecieListResponse struct {
	ID              int64                                 `json:"id"`
	ComName         string                                `json:"comName"`
	Firstname       string                                `json:"firstname"`
	HowMany         int64                                 `json:"howMany"`
	Lastname        string                                `json:"lastname"`
	Lat             float64                               `json:"lat"`
	Lng             float64                               `json:"lng"`
	LocationPrivate bool                                  `json:"locationPrivate"`
	LocID           string                                `json:"locId"`
	LocName         string                                `json:"locName"`
	ObsDt           string                                `json:"obsDt"`
	ObsReviewed     bool                                  `json:"obsReviewed"`
	ObsValid        bool                                  `json:"obsValid"`
	SciName         string                                `json:"sciName"`
	SpeciesCode     string                                `json:"speciesCode"`
	SubID           string                                `json:"subId"`
	JSON            dataObGeoRecentSpecieListResponseJSON `json:"-"`
}

// dataObGeoRecentSpecieListResponseJSON contains the JSON metadata for the struct
// [DataObGeoRecentSpecieListResponse]
type dataObGeoRecentSpecieListResponseJSON struct {
	ID              apijson.Field
	ComName         apijson.Field
	Firstname       apijson.Field
	HowMany         apijson.Field
	Lastname        apijson.Field
	Lat             apijson.Field
	Lng             apijson.Field
	LocationPrivate apijson.Field
	LocID           apijson.Field
	LocName         apijson.Field
	ObsDt           apijson.Field
	ObsReviewed     apijson.Field
	ObsValid        apijson.Field
	SciName         apijson.Field
	SpeciesCode     apijson.Field
	SubID           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DataObGeoRecentSpecieListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataObGeoRecentSpecieListResponseJSON) RawJSON() string {
	return r.raw
}

type DataObGeoRecentSpecieListParams struct {
	Lat param.Field[float64] `query:"lat,required"`
	Lng param.Field[float64] `query:"lng,required"`
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// The search radius from the given position, in kilometers.
	Dist param.Field[int64] `query:"dist"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Include observations which have not yet been reviewed.
	IncludeProvisional param.Field[bool] `query:"includeProvisional"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObGeoRecentSpecieListParams]'s query parameters as
// `url.Values`.
func (r DataObGeoRecentSpecieListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
