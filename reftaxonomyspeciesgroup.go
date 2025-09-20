// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
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

// RefTaxonomySpeciesGroupService contains methods and other services that help
// with interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefTaxonomySpeciesGroupService] method instead.
type RefTaxonomySpeciesGroupService struct {
	Options []option.RequestOption
}

// NewRefTaxonomySpeciesGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRefTaxonomySpeciesGroupService(opts ...option.RequestOption) (r *RefTaxonomySpeciesGroupService) {
	r = &RefTaxonomySpeciesGroupService{}
	r.Options = opts
	return
}

// Get the list of species groups, e.g. terns, finches, etc. #### Notes Merlin puts
// like birds together, with Falcons next to Hawks, whereas eBird follows taxonomic
// order.
func (r *RefTaxonomySpeciesGroupService) List(ctx context.Context, speciesGrouping RefTaxonomySpeciesGroupListParamsSpeciesGrouping, query RefTaxonomySpeciesGroupListParams, opts ...option.RequestOption) (res *[]RefTaxonomySpeciesGroupListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ref/sppgroup/%v", speciesGrouping)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RefTaxonomySpeciesGroupListResponse struct {
	GroupName        string                                  `json:"groupName"`
	GroupOrder       int64                                   `json:"groupOrder"`
	TaxonOrderBounds [][]float64                             `json:"taxonOrderBounds"`
	JSON             refTaxonomySpeciesGroupListResponseJSON `json:"-"`
}

// refTaxonomySpeciesGroupListResponseJSON contains the JSON metadata for the
// struct [RefTaxonomySpeciesGroupListResponse]
type refTaxonomySpeciesGroupListResponseJSON struct {
	GroupName        apijson.Field
	GroupOrder       apijson.Field
	TaxonOrderBounds apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RefTaxonomySpeciesGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refTaxonomySpeciesGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type RefTaxonomySpeciesGroupListParams struct {
	// Locale for species group names. English names are returned for any non-listed
	// locale or any non-translated group name.
	GroupNameLocale param.Field[string] `query:"groupNameLocale"`
}

// URLQuery serializes [RefTaxonomySpeciesGroupListParams]'s query parameters as
// `url.Values`.
func (r RefTaxonomySpeciesGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The order in which groups are returned.
type RefTaxonomySpeciesGroupListParamsSpeciesGrouping string

const (
	RefTaxonomySpeciesGroupListParamsSpeciesGroupingMerlin RefTaxonomySpeciesGroupListParamsSpeciesGrouping = "merlin"
	RefTaxonomySpeciesGroupListParamsSpeciesGroupingEbird  RefTaxonomySpeciesGroupListParamsSpeciesGrouping = "ebird"
)

func (r RefTaxonomySpeciesGroupListParamsSpeciesGrouping) IsKnown() bool {
	switch r {
	case RefTaxonomySpeciesGroupListParamsSpeciesGroupingMerlin, RefTaxonomySpeciesGroupListParamsSpeciesGroupingEbird:
		return true
	}
	return false
}
