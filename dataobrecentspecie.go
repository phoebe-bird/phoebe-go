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

// DataObRecentSpecieService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObRecentSpecieService] method instead.
type DataObRecentSpecieService struct {
	Options []option.RequestOption
}

// NewDataObRecentSpecieService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDataObRecentSpecieService(opts ...option.RequestOption) (r *DataObRecentSpecieService) {
	r = &DataObRecentSpecieService{}
	r.Options = opts
	return
}

// Get the recent observations, up to 30 days ago, of a particular species in a
// country, region or location. Results include only the most recent observation
// from each location in the region specified.
//
// #### Notes
//
// The species code is typically a 6-letter code, e.g. cangoo for Canada Goose. You
// can get complete set of species code from the GET eBird Taxonomy end-point.
//
// When using the _r_ query parameter set the _regionCode_ URL parameter to an
// empty string.
func (r *DataObRecentSpecieService) Get(ctx context.Context, regionCode string, speciesCode string, query DataObRecentSpecieGetParams, opts ...option.RequestOption) (res *[]DataObRecentSpecieGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	if speciesCode == "" {
		err = errors.New("missing required speciesCode parameter")
		return
	}
	path := fmt.Sprintf("data/obs/%s/recent/%s", regionCode, speciesCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObRecentSpecieGetResponse struct {
	ID              int64                             `json:"id"`
	ComName         string                            `json:"comName"`
	Firstname       string                            `json:"firstname"`
	HowMany         int64                             `json:"howMany"`
	Lastname        string                            `json:"lastname"`
	Lat             float64                           `json:"lat"`
	Lng             float64                           `json:"lng"`
	LocationPrivate bool                              `json:"locationPrivate"`
	LocID           string                            `json:"locId"`
	LocName         string                            `json:"locName"`
	ObsDt           string                            `json:"obsDt"`
	ObsReviewed     bool                              `json:"obsReviewed"`
	ObsValid        bool                              `json:"obsValid"`
	SciName         string                            `json:"sciName"`
	SpeciesCode     string                            `json:"speciesCode"`
	SubID           string                            `json:"subId"`
	JSON            dataObRecentSpecieGetResponseJSON `json:"-"`
}

// dataObRecentSpecieGetResponseJSON contains the JSON metadata for the struct
// [DataObRecentSpecieGetResponse]
type dataObRecentSpecieGetResponseJSON struct {
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

func (r *DataObRecentSpecieGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataObRecentSpecieGetResponseJSON) RawJSON() string {
	return r.raw
}

type DataObRecentSpecieGetParams struct {
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Include observations which have not yet been reviewed.
	IncludeProvisional param.Field[bool] `query:"includeProvisional"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Fetch observations from up to 10 locations
	R param.Field[[]string] `query:"r"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObRecentSpecieGetParams]'s query parameters as
// `url.Values`.
func (r DataObRecentSpecieGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
