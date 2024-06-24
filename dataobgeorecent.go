// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/phoebe-go/internal/apijson"
	"github.com/stainless-sdks/phoebe-go/internal/apiquery"
	"github.com/stainless-sdks/phoebe-go/internal/param"
	"github.com/stainless-sdks/phoebe-go/internal/requestconfig"
	"github.com/stainless-sdks/phoebe-go/option"
)

// DataObGeoRecentService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObGeoRecentService] method instead.
type DataObGeoRecentService struct {
	Options []option.RequestOption
	Species *DataObGeoRecentSpecieService
	Notable *DataObGeoRecentNotableService
}

// NewDataObGeoRecentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataObGeoRecentService(opts ...option.RequestOption) (r *DataObGeoRecentService) {
	r = &DataObGeoRecentService{}
	r.Options = opts
	r.Species = NewDataObGeoRecentSpecieService(opts...)
	r.Notable = NewDataObGeoRecentNotableService(opts...)
	return
}

// Get the list of recent observations (up to 30 days ago) of birds seen at
// locations within a radius of up to 50 kilometers, from a given set of
// coordinates. Results include only the most recent observation for each species
// in the region specified.
func (r *DataObGeoRecentService) List(ctx context.Context, query DataObGeoRecentListParams, opts ...option.RequestOption) (res *[]DataObGeoRecentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "data/obs/geo/recent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObGeoRecentListResponse struct {
	ID              int64                           `json:"id"`
	ComName         string                          `json:"comName"`
	Firstname       string                          `json:"firstname"`
	HowMany         int64                           `json:"howMany"`
	Lastname        string                          `json:"lastname"`
	Lat             float64                         `json:"lat"`
	Lng             float64                         `json:"lng"`
	LocationPrivate bool                            `json:"locationPrivate"`
	LocID           string                          `json:"locId"`
	LocName         string                          `json:"locName"`
	ObsDt           string                          `json:"obsDt"`
	ObsReviewed     bool                            `json:"obsReviewed"`
	ObsValid        bool                            `json:"obsValid"`
	SciName         string                          `json:"sciName"`
	SpeciesCode     string                          `json:"speciesCode"`
	SubID           string                          `json:"subId"`
	JSON            dataObGeoRecentListResponseJSON `json:"-"`
}

// dataObGeoRecentListResponseJSON contains the JSON metadata for the struct
// [DataObGeoRecentListResponse]
type dataObGeoRecentListResponseJSON struct {
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

func (r *DataObGeoRecentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataObGeoRecentListResponseJSON) RawJSON() string {
	return r.raw
}

type DataObGeoRecentListParams struct {
	Lat param.Field[float64] `query:"lat,required"`
	Lng param.Field[float64] `query:"lng,required"`
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Only fetch observations from these taxonomic categories
	Cat param.Field[DataObGeoRecentListParamsCat] `query:"cat"`
	// The search radius from the given position, in kilometers.
	Dist param.Field[int64] `query:"dist"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Include observations which have not yet been reviewed.
	IncludeProvisional param.Field[bool] `query:"includeProvisional"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Sort observations by taxonomy or by date, most recent first.
	Sort param.Field[DataObGeoRecentListParamsSort] `query:"sort"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObGeoRecentListParams]'s query parameters as
// `url.Values`.
func (r DataObGeoRecentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only fetch observations from these taxonomic categories
type DataObGeoRecentListParamsCat string

const (
	DataObGeoRecentListParamsCatSpecies    DataObGeoRecentListParamsCat = "species"
	DataObGeoRecentListParamsCatSlash      DataObGeoRecentListParamsCat = "slash"
	DataObGeoRecentListParamsCatIssf       DataObGeoRecentListParamsCat = "issf"
	DataObGeoRecentListParamsCatSpuh       DataObGeoRecentListParamsCat = "spuh"
	DataObGeoRecentListParamsCatHybrid     DataObGeoRecentListParamsCat = "hybrid"
	DataObGeoRecentListParamsCatDomestic   DataObGeoRecentListParamsCat = "domestic"
	DataObGeoRecentListParamsCatForm       DataObGeoRecentListParamsCat = "form"
	DataObGeoRecentListParamsCatIntergrade DataObGeoRecentListParamsCat = "intergrade"
)

func (r DataObGeoRecentListParamsCat) IsKnown() bool {
	switch r {
	case DataObGeoRecentListParamsCatSpecies, DataObGeoRecentListParamsCatSlash, DataObGeoRecentListParamsCatIssf, DataObGeoRecentListParamsCatSpuh, DataObGeoRecentListParamsCatHybrid, DataObGeoRecentListParamsCatDomestic, DataObGeoRecentListParamsCatForm, DataObGeoRecentListParamsCatIntergrade:
		return true
	}
	return false
}

// Sort observations by taxonomy or by date, most recent first.
type DataObGeoRecentListParamsSort string

const (
	DataObGeoRecentListParamsSortDate    DataObGeoRecentListParamsSort = "date"
	DataObGeoRecentListParamsSortSpecies DataObGeoRecentListParamsSort = "species"
)

func (r DataObGeoRecentListParamsSort) IsKnown() bool {
	switch r {
	case DataObGeoRecentListParamsSortDate, DataObGeoRecentListParamsSortSpecies:
		return true
	}
	return false
}
