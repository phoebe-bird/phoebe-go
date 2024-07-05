// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebebird

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

// ProductListService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductListService] method instead.
type ProductListService struct {
	Options    []option.RequestOption
	Historical *ProductListHistoricalService
}

// NewProductListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProductListService(opts ...option.RequestOption) (r *ProductListService) {
	r = &ProductListService{}
	r.Options = opts
	r.Historical = NewProductListHistoricalService(opts...)
	return
}

// Get information on the most recently submitted checklists for a region.
func (r *ProductListService) Get(ctx context.Context, regionCode string, query ProductListGetParams, opts ...option.RequestOption) (res *[]ProductListGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("product/lists/%s", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ProductListGetResponse struct {
	AllObsReported       bool                       `json:"allObsReported"`
	ChecklistID          string                     `json:"checklistId"`
	CreationDt           string                     `json:"creationDt"`
	DurationHrs          float64                    `json:"durationHrs"`
	ISOObsDate           string                     `json:"isoObsDate"`
	LastEditedDt         string                     `json:"lastEditedDt"`
	Loc                  ProductListGetResponseLoc  `json:"loc"`
	LocID                string                     `json:"locId"`
	NumObservers         int64                      `json:"numObservers"`
	NumSpecies           int64                      `json:"numSpecies"`
	Obs                  []ProductListGetResponseOb `json:"obs"`
	ObsDt                string                     `json:"obsDt"`
	ObsTime              string                     `json:"obsTime"`
	ObsTimeValid         bool                       `json:"obsTimeValid"`
	ProjID               string                     `json:"projId"`
	ProtocolID           string                     `json:"protocolId"`
	SubID                string                     `json:"subId"`
	SubID                string                     `json:"subID"`
	SubmissionMethodCode string                     `json:"submissionMethodCode"`
	Subnational1Code     string                     `json:"subnational1Code"`
	UserDisplayName      string                     `json:"userDisplayName"`
	JSON                 productListGetResponseJSON `json:"-"`
}

// productListGetResponseJSON contains the JSON metadata for the struct
// [ProductListGetResponse]
type productListGetResponseJSON struct {
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
	SubID                apijson.Field
	SubmissionMethodCode apijson.Field
	Subnational1Code     apijson.Field
	UserDisplayName      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProductListGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProductListGetResponseLoc struct {
	CountryCode      string                        `json:"countryCode"`
	CountryName      string                        `json:"countryName"`
	HierarchicalName string                        `json:"hierarchicalName"`
	IsHotspot        bool                          `json:"isHotspot"`
	Lat              float64                       `json:"lat"`
	Latitude         float64                       `json:"latitude"`
	Lng              float64                       `json:"lng"`
	LocID            string                        `json:"locId"`
	LocID            string                        `json:"locID"`
	LocName          string                        `json:"locName"`
	Longitude        float64                       `json:"longitude"`
	Name             string                        `json:"name"`
	Subnational1Code string                        `json:"subnational1Code"`
	Subnational1Name string                        `json:"subnational1Name"`
	JSON             productListGetResponseLocJSON `json:"-"`
}

// productListGetResponseLocJSON contains the JSON metadata for the struct
// [ProductListGetResponseLoc]
type productListGetResponseLocJSON struct {
	CountryCode      apijson.Field
	CountryName      apijson.Field
	HierarchicalName apijson.Field
	IsHotspot        apijson.Field
	Lat              apijson.Field
	Latitude         apijson.Field
	Lng              apijson.Field
	LocID            apijson.Field
	LocID            apijson.Field
	LocName          apijson.Field
	Longitude        apijson.Field
	Name             apijson.Field
	Subnational1Code apijson.Field
	Subnational1Name apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProductListGetResponseLoc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListGetResponseLocJSON) RawJSON() string {
	return r.raw
}

type ProductListGetResponseOb struct {
	ObsAux      []ProductListGetResponseObsObsAux `json:"obsAux"`
	ObsDt       string                            `json:"obsDt"`
	ObsID       string                            `json:"obsId"`
	SpeciesCode string                            `json:"speciesCode"`
	JSON        productListGetResponseObJSON      `json:"-"`
}

// productListGetResponseObJSON contains the JSON metadata for the struct
// [ProductListGetResponseOb]
type productListGetResponseObJSON struct {
	ObsAux      apijson.Field
	ObsDt       apijson.Field
	ObsID       apijson.Field
	SpeciesCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListGetResponseOb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListGetResponseObJSON) RawJSON() string {
	return r.raw
}

type ProductListGetResponseObsObsAux struct {
	AuxCode         string                              `json:"auxCode"`
	EntryMethodCode string                              `json:"entryMethodCode"`
	FieldName       string                              `json:"fieldName"`
	ObsID           string                              `json:"obsId"`
	SpeciesCode     string                              `json:"speciesCode"`
	SubID           string                              `json:"subId"`
	Value           string                              `json:"value"`
	JSON            productListGetResponseObsObsAuxJSON `json:"-"`
}

// productListGetResponseObsObsAuxJSON contains the JSON metadata for the struct
// [ProductListGetResponseObsObsAux]
type productListGetResponseObsObsAuxJSON struct {
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

func (r *ProductListGetResponseObsObsAux) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListGetResponseObsObsAuxJSON) RawJSON() string {
	return r.raw
}

type ProductListGetParams struct {
	// Only fetch this number of checklists.
	MaxResults param.Field[int64] `query:"maxResults"`
}

// URLQuery serializes [ProductListGetParams]'s query parameters as `url.Values`.
func (r ProductListGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
