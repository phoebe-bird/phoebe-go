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

// DataObRecentService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObRecentService] method instead.
type DataObRecentService struct {
	Options []option.RequestOption
	Notable *DataObRecentNotableService
	Species *DataObRecentSpecieService
}

// NewDataObRecentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataObRecentService(opts ...option.RequestOption) (r *DataObRecentService) {
	r = &DataObRecentService{}
	r.Options = opts
	r.Notable = NewDataObRecentNotableService(opts...)
	r.Species = NewDataObRecentSpecieService(opts...)
	return
}

// Get the list of recent observations (up to 30 days ago) of birds seen in a
// country, state, county, or location. Results include only the most recent
// observation for each species in the region specified.
func (r *DataObRecentService) List(ctx context.Context, regionCode string, query DataObRecentListParams, opts ...option.RequestOption) (res *[]DataObRecentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("data/obs/%s/recent", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObRecentListResponse struct {
	ID              int64                        `json:"id"`
	ComName         string                       `json:"comName"`
	Firstname       string                       `json:"firstname"`
	HowMany         int64                        `json:"howMany"`
	Lastname        string                       `json:"lastname"`
	Lat             float64                      `json:"lat"`
	Lng             float64                      `json:"lng"`
	LocationPrivate bool                         `json:"locationPrivate"`
	LocID           string                       `json:"locId"`
	LocName         string                       `json:"locName"`
	ObsDt           string                       `json:"obsDt"`
	ObsReviewed     bool                         `json:"obsReviewed"`
	ObsValid        bool                         `json:"obsValid"`
	SciName         string                       `json:"sciName"`
	SpeciesCode     string                       `json:"speciesCode"`
	SubID           string                       `json:"subId"`
	JSON            dataObRecentListResponseJSON `json:"-"`
}

// dataObRecentListResponseJSON contains the JSON metadata for the struct
// [DataObRecentListResponse]
type dataObRecentListResponseJSON struct {
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

func (r *DataObRecentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataObRecentListResponseJSON) RawJSON() string {
	return r.raw
}

type DataObRecentListParams struct {
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Only fetch observations from these taxonomic categories
	Cat param.Field[DataObRecentListParamsCat] `query:"cat"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Include observations which have not yet been reviewed
	IncludeProvisional param.Field[bool] `query:"includeProvisional"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Fetch observations from up to 10 locations
	R param.Field[[]string] `query:"r"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObRecentListParams]'s query parameters as `url.Values`.
func (r DataObRecentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only fetch observations from these taxonomic categories
type DataObRecentListParamsCat string

const (
	DataObRecentListParamsCatSpecies    DataObRecentListParamsCat = "species"
	DataObRecentListParamsCatSlash      DataObRecentListParamsCat = "slash"
	DataObRecentListParamsCatIssf       DataObRecentListParamsCat = "issf"
	DataObRecentListParamsCatSpuh       DataObRecentListParamsCat = "spuh"
	DataObRecentListParamsCatHybrid     DataObRecentListParamsCat = "hybrid"
	DataObRecentListParamsCatDomestic   DataObRecentListParamsCat = "domestic"
	DataObRecentListParamsCatForm       DataObRecentListParamsCat = "form"
	DataObRecentListParamsCatIntergrade DataObRecentListParamsCat = "intergrade"
)

func (r DataObRecentListParamsCat) IsKnown() bool {
	switch r {
	case DataObRecentListParamsCatSpecies, DataObRecentListParamsCatSlash, DataObRecentListParamsCatIssf, DataObRecentListParamsCatSpuh, DataObRecentListParamsCatHybrid, DataObRecentListParamsCatDomestic, DataObRecentListParamsCatForm, DataObRecentListParamsCatIntergrade:
		return true
	}
	return false
}
