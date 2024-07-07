// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/phoebe-bird/phoebe-go/internal/apijson"
	"github.com/phoebe-bird/phoebe-go/internal/apiquery"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// ProductListHistoricalService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductListHistoricalService] method instead.
type ProductListHistoricalService struct {
	Options []option.RequestOption
}

// NewProductListHistoricalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProductListHistoricalService(opts ...option.RequestOption) (r *ProductListHistoricalService) {
	r = &ProductListHistoricalService{}
	r.Options = opts
	return
}

// Get information on the checklists submitted on a given date for a country or
// region.
func (r *ProductListHistoricalService) Get(ctx context.Context, regionCode string, y int64, m int64, d int64, query ProductListHistoricalGetParams, opts ...option.RequestOption) (res *[]ProductListHistoricalGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("product/lists/%s/%v/%v/%v", regionCode, y, m, d)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ProductListHistoricalGetResponse struct {
	AllObsReported       bool                                 `json:"allObsReported"`
	ChecklistID          string                               `json:"checklistId"`
	CreationDt           string                               `json:"creationDt"`
	DurationHrs          float64                              `json:"durationHrs"`
	ISOObsDate           string                               `json:"isoObsDate"`
	LastEditedDt         string                               `json:"lastEditedDt"`
	Loc                  ProductListHistoricalGetResponseLoc  `json:"loc"`
	LocID                string                               `json:"locId"`
	NumObservers         int64                                `json:"numObservers"`
	NumSpecies           int64                                `json:"numSpecies"`
	Obs                  []ProductListHistoricalGetResponseOb `json:"obs"`
	ObsDt                string                               `json:"obsDt"`
	ObsTime              string                               `json:"obsTime"`
	ObsTimeValid         bool                                 `json:"obsTimeValid"`
	ProjID               string                               `json:"projId"`
	ProtocolID           string                               `json:"protocolId"`
	SubID                string                               `json:"subId"`
	SubmissionMethodCode string                               `json:"submissionMethodCode"`
	Subnational1Code     string                               `json:"subnational1Code"`
	UserDisplayName      string                               `json:"userDisplayName"`
	JSON                 productListHistoricalGetResponseJSON `json:"-"`
}

// productListHistoricalGetResponseJSON contains the JSON metadata for the struct
// [ProductListHistoricalGetResponse]
type productListHistoricalGetResponseJSON struct {
	AllObsReported       apijson.Field
	ChecklistID          apijson.Field
	CreationDt           apijson.Field
	DurationHrs          apijson.Field
	ISOObsDate           apijson.Field
	LastEditedDt         apijson.Field
	Loc                  apijson.Field
	LocID                apijson.Field
	NumObservers         apijson.Field
	NumSpecies           apijson.Field
	Obs                  apijson.Field
	ObsDt                apijson.Field
	ObsTime              apijson.Field
	ObsTimeValid         apijson.Field
	ProjID               apijson.Field
	ProtocolID           apijson.Field
	SubID                apijson.Field
	SubmissionMethodCode apijson.Field
	Subnational1Code     apijson.Field
	UserDisplayName      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProductListHistoricalGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListHistoricalGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProductListHistoricalGetResponseLoc struct {
	CountryCode      string                                  `json:"countryCode"`
	CountryName      string                                  `json:"countryName"`
	HierarchicalName string                                  `json:"hierarchicalName"`
	IsHotspot        bool                                    `json:"isHotspot"`
	Lat              float64                                 `json:"lat"`
	Latitude         float64                                 `json:"latitude"`
	Lng              float64                                 `json:"lng"`
	LocID            string                                  `json:"locId"`
	LocName          string                                  `json:"locName"`
	Longitude        float64                                 `json:"longitude"`
	Name             string                                  `json:"name"`
	Subnational1Code string                                  `json:"subnational1Code"`
	Subnational1Name string                                  `json:"subnational1Name"`
	JSON             productListHistoricalGetResponseLocJSON `json:"-"`
}

// productListHistoricalGetResponseLocJSON contains the JSON metadata for the
// struct [ProductListHistoricalGetResponseLoc]
type productListHistoricalGetResponseLocJSON struct {
	CountryCode      apijson.Field
	CountryName      apijson.Field
	HierarchicalName apijson.Field
	IsHotspot        apijson.Field
	Lat              apijson.Field
	Latitude         apijson.Field
	Lng              apijson.Field
	LocID            apijson.Field
	LocName          apijson.Field
	Longitude        apijson.Field
	Name             apijson.Field
	Subnational1Code apijson.Field
	Subnational1Name apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProductListHistoricalGetResponseLoc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListHistoricalGetResponseLocJSON) RawJSON() string {
	return r.raw
}

type ProductListHistoricalGetResponseOb struct {
	ObsAux      []ProductListHistoricalGetResponseObsObsAux `json:"obsAux"`
	ObsDt       string                                      `json:"obsDt"`
	ObsID       string                                      `json:"obsId"`
	SpeciesCode string                                      `json:"speciesCode"`
	JSON        productListHistoricalGetResponseObJSON      `json:"-"`
}

// productListHistoricalGetResponseObJSON contains the JSON metadata for the struct
// [ProductListHistoricalGetResponseOb]
type productListHistoricalGetResponseObJSON struct {
	ObsAux      apijson.Field
	ObsDt       apijson.Field
	ObsID       apijson.Field
	SpeciesCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListHistoricalGetResponseOb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListHistoricalGetResponseObJSON) RawJSON() string {
	return r.raw
}

type ProductListHistoricalGetResponseObsObsAux struct {
	AuxCode         string                                        `json:"auxCode"`
	EntryMethodCode string                                        `json:"entryMethodCode"`
	FieldName       string                                        `json:"fieldName"`
	ObsID           string                                        `json:"obsId"`
	SpeciesCode     string                                        `json:"speciesCode"`
	SubID           string                                        `json:"subId"`
	Value           string                                        `json:"value"`
	JSON            productListHistoricalGetResponseObsObsAuxJSON `json:"-"`
}

// productListHistoricalGetResponseObsObsAuxJSON contains the JSON metadata for the
// struct [ProductListHistoricalGetResponseObsObsAux]
type productListHistoricalGetResponseObsObsAuxJSON struct {
	AuxCode         apijson.Field
	EntryMethodCode apijson.Field
	FieldName       apijson.Field
	ObsID           apijson.Field
	SpeciesCode     apijson.Field
	SubID           apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProductListHistoricalGetResponseObsObsAux) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListHistoricalGetResponseObsObsAuxJSON) RawJSON() string {
	return r.raw
}

type ProductListHistoricalGetParams struct {
	// Only fetch this number of checklists.
	MaxResults param.Field[int64] `query:"maxResults"`
	// Order the results by the date of the checklist or by the date it was submitted.
	SortKey param.Field[ProductListHistoricalGetParamsSortKey] `query:"sortKey"`
}

// URLQuery serializes [ProductListHistoricalGetParams]'s query parameters as
// `url.Values`.
func (r ProductListHistoricalGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Order the results by the date of the checklist or by the date it was submitted.
type ProductListHistoricalGetParamsSortKey string

const (
	ProductListHistoricalGetParamsSortKeyObsDt      ProductListHistoricalGetParamsSortKey = "obs_dt"
	ProductListHistoricalGetParamsSortKeyCreationDt ProductListHistoricalGetParamsSortKey = "creation_dt"
)

func (r ProductListHistoricalGetParamsSortKey) IsKnown() bool {
	switch r {
	case ProductListHistoricalGetParamsSortKeyObsDt, ProductListHistoricalGetParamsSortKeyCreationDt:
		return true
	}
	return false
}
