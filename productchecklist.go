// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebebird

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/phoebe-bird/phoebe-go/internal/apijson"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// ProductChecklistService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductChecklistService] method instead.
type ProductChecklistService struct {
	Options []option.RequestOption
}

// NewProductChecklistService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProductChecklistService(opts ...option.RequestOption) (r *ProductChecklistService) {
	r = &ProductChecklistService{}
	r.Options = opts
	return
}

// Get the details and observations of a checklist.
//
// #### Notes Do NOT use this to download large amounts of data. You will be banned if you do. In the fields for each observation, the following fields are duplicates or obsolete and will be removed at a future date: _howManyAtleast_, _howManyAtmost_, _hideFlags_, _projId_, _subId_, _subnational1Code_ and _present_.
func (r *ProductChecklistService) View(ctx context.Context, subID string, opts ...option.RequestOption) (res *ProductChecklistViewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if subID == "" {
		err = errors.New("missing required subId parameter")
		return
	}
	path := fmt.Sprintf("product/checklist/view/%s", subID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ProductChecklistViewResponse struct {
	AllObsReported       bool                             `json:"allObsReported"`
	ChecklistID          string                           `json:"checklistId"`
	CreationDt           string                           `json:"creationDt"`
	DurationHrs          float64                          `json:"durationHrs"`
	ISOObsDate           string                           `json:"isoObsDate"`
	LastEditedDt         string                           `json:"lastEditedDt"`
	Loc                  ProductChecklistViewResponseLoc  `json:"loc"`
	LocID                string                           `json:"locId"`
	NumObservers         int64                            `json:"numObservers"`
	NumSpecies           int64                            `json:"numSpecies"`
	Obs                  []ProductChecklistViewResponseOb `json:"obs"`
	ObsDt                string                           `json:"obsDt"`
	ObsTime              string                           `json:"obsTime"`
	ObsTimeValid         bool                             `json:"obsTimeValid"`
	ProjID               string                           `json:"projId"`
	ProtocolID           string                           `json:"protocolId"`
	SubID                string                           `json:"subId"`
	SubID                string                           `json:"subID"`
	SubmissionMethodCode string                           `json:"submissionMethodCode"`
	Subnational1Code     string                           `json:"subnational1Code"`
	UserDisplayName      string                           `json:"userDisplayName"`
	JSON                 productChecklistViewResponseJSON `json:"-"`
}

// productChecklistViewResponseJSON contains the JSON metadata for the struct
// [ProductChecklistViewResponse]
type productChecklistViewResponseJSON struct {
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

func (r *ProductChecklistViewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productChecklistViewResponseJSON) RawJSON() string {
	return r.raw
}

type ProductChecklistViewResponseLoc struct {
	CountryCode      string                              `json:"countryCode"`
	CountryName      string                              `json:"countryName"`
	HierarchicalName string                              `json:"hierarchicalName"`
	IsHotspot        bool                                `json:"isHotspot"`
	Lat              float64                             `json:"lat"`
	Latitude         float64                             `json:"latitude"`
	Lng              float64                             `json:"lng"`
	LocID            string                              `json:"locId"`
	LocID            string                              `json:"locID"`
	LocName          string                              `json:"locName"`
	Longitude        float64                             `json:"longitude"`
	Name             string                              `json:"name"`
	Subnational1Code string                              `json:"subnational1Code"`
	Subnational1Name string                              `json:"subnational1Name"`
	JSON             productChecklistViewResponseLocJSON `json:"-"`
}

// productChecklistViewResponseLocJSON contains the JSON metadata for the struct
// [ProductChecklistViewResponseLoc]
type productChecklistViewResponseLocJSON struct {
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

func (r *ProductChecklistViewResponseLoc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productChecklistViewResponseLocJSON) RawJSON() string {
	return r.raw
}

type ProductChecklistViewResponseOb struct {
	ObsAux      []ProductChecklistViewResponseObsObsAux `json:"obsAux"`
	ObsDt       string                                  `json:"obsDt"`
	ObsID       string                                  `json:"obsId"`
	SpeciesCode string                                  `json:"speciesCode"`
	JSON        productChecklistViewResponseObJSON      `json:"-"`
}

// productChecklistViewResponseObJSON contains the JSON metadata for the struct
// [ProductChecklistViewResponseOb]
type productChecklistViewResponseObJSON struct {
	ObsAux      apijson.Field
	ObsDt       apijson.Field
	ObsID       apijson.Field
	SpeciesCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductChecklistViewResponseOb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productChecklistViewResponseObJSON) RawJSON() string {
	return r.raw
}

type ProductChecklistViewResponseObsObsAux struct {
	AuxCode         string                                    `json:"auxCode"`
	EntryMethodCode string                                    `json:"entryMethodCode"`
	FieldName       string                                    `json:"fieldName"`
	ObsID           string                                    `json:"obsId"`
	SpeciesCode     string                                    `json:"speciesCode"`
	SubID           string                                    `json:"subId"`
	Value           string                                    `json:"value"`
	JSON            productChecklistViewResponseObsObsAuxJSON `json:"-"`
}

// productChecklistViewResponseObsObsAuxJSON contains the JSON metadata for the
// struct [ProductChecklistViewResponseObsObsAux]
type productChecklistViewResponseObsObsAuxJSON struct {
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

func (r *ProductChecklistViewResponseObsObsAux) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productChecklistViewResponseObsObsAuxJSON) RawJSON() string {
	return r.raw
}
