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

// DataObRecentNotableService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObRecentNotableService] method instead.
type DataObRecentNotableService struct {
	Options []option.RequestOption
}

// NewDataObRecentNotableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDataObRecentNotableService(opts ...option.RequestOption) (r *DataObRecentNotableService) {
	r = &DataObRecentNotableService{}
	r.Options = opts
	return
}

// Get the list of recent, notable observations (up to 30 days ago) of birds seen
// in a country, region or location. Notable observations can be for locally or
// nationally rare species or are otherwise unusual, e.g. over-wintering birds in a
// species which is normally only a summer visitor.
func (r *DataObRecentNotableService) List(ctx context.Context, regionCode string, query DataObRecentNotableListParams, opts ...option.RequestOption) (res *[]DataObRecentNotableListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("data/obs/%s/recent/notable", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObRecentNotableListResponse struct {
	ID              int64                               `json:"id"`
	ComName         string                              `json:"comName"`
	Firstname       string                              `json:"firstname"`
	HowMany         int64                               `json:"howMany"`
	Lastname        string                              `json:"lastname"`
	Lat             float64                             `json:"lat"`
	Lng             float64                             `json:"lng"`
	LocationPrivate bool                                `json:"locationPrivate"`
	LocID           string                              `json:"locId"`
	LocName         string                              `json:"locName"`
	ObsDt           string                              `json:"obsDt"`
	ObsReviewed     bool                                `json:"obsReviewed"`
	ObsValid        bool                                `json:"obsValid"`
	SciName         string                              `json:"sciName"`
	SpeciesCode     string                              `json:"speciesCode"`
	SubID           string                              `json:"subId"`
	JSON            dataObRecentNotableListResponseJSON `json:"-"`
}

// dataObRecentNotableListResponseJSON contains the JSON metadata for the struct
// [DataObRecentNotableListResponse]
type dataObRecentNotableListResponseJSON struct {
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

func (r *DataObRecentNotableListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataObRecentNotableListResponseJSON) RawJSON() string {
	return r.raw
}

type DataObRecentNotableListParams struct {
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Include a subset (simple), or all (full), of the fields available.
	Detail param.Field[DataObRecentNotableListParamsDetail] `query:"detail"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Fetch observations from up to 10 locations
	R param.Field[[]string] `query:"r"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObRecentNotableListParams]'s query parameters as
// `url.Values`.
func (r DataObRecentNotableListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Include a subset (simple), or all (full), of the fields available.
type DataObRecentNotableListParamsDetail string

const (
	DataObRecentNotableListParamsDetailSimple DataObRecentNotableListParamsDetail = "simple"
	DataObRecentNotableListParamsDetailFull   DataObRecentNotableListParamsDetail = "full"
)

func (r DataObRecentNotableListParamsDetail) IsKnown() bool {
	switch r {
	case DataObRecentNotableListParamsDetailSimple, DataObRecentNotableListParamsDetailFull:
		return true
	}
	return false
}
