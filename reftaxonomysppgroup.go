// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/phoebe-go/internal/apijson"
	"github.com/stainless-sdks/phoebe-go/internal/apiquery"
	"github.com/stainless-sdks/phoebe-go/internal/param"
	"github.com/stainless-sdks/phoebe-go/internal/requestconfig"
	"github.com/stainless-sdks/phoebe-go/option"
)

// RefTaxonomySppgroupService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefTaxonomySppgroupService] method instead.
type RefTaxonomySppgroupService struct {
	Options []option.RequestOption
}

// NewRefTaxonomySppgroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRefTaxonomySppgroupService(opts ...option.RequestOption) (r *RefTaxonomySppgroupService) {
	r = &RefTaxonomySppgroupService{}
	r.Options = opts
	return
}

// Get the list of species groups, e.g. terns, finches, etc. #### Notes Merlin puts
// like birds together, with Falcons next to Hawks, whereas eBird follows taxonomic
// order.
func (r *RefTaxonomySppgroupService) Get(ctx context.Context, speciesGrouping RefTaxonomySppgroupGetParamsSpeciesGrouping, query RefTaxonomySppgroupGetParams, opts ...option.RequestOption) (res *[]RefTaxonomySppgroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ref/sppgroup/%v", speciesGrouping)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RefTaxonomySppgroupGetResponse struct {
	GroupName        string                             `json:"groupName"`
	GroupOrder       int64                              `json:"groupOrder"`
	TaxonOrderBounds [][]float64                        `json:"taxonOrderBounds"`
	JSON             refTaxonomySppgroupGetResponseJSON `json:"-"`
}

// refTaxonomySppgroupGetResponseJSON contains the JSON metadata for the struct
// [RefTaxonomySppgroupGetResponse]
type refTaxonomySppgroupGetResponseJSON struct {
	GroupName        apijson.Field
	GroupOrder       apijson.Field
	TaxonOrderBounds apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RefTaxonomySppgroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refTaxonomySppgroupGetResponseJSON) RawJSON() string {
	return r.raw
}

type RefTaxonomySppgroupGetParams struct {
	// Locale for species group names. English names are returned for any non-listed
	// locale or any non-translated group name.
	GroupNameLocale param.Field[string] `query:"groupNameLocale"`
}

// URLQuery serializes [RefTaxonomySppgroupGetParams]'s query parameters as
// `url.Values`.
func (r RefTaxonomySppgroupGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The order in which groups are returned.
type RefTaxonomySppgroupGetParamsSpeciesGrouping string

const (
	RefTaxonomySppgroupGetParamsSpeciesGroupingMerlin RefTaxonomySppgroupGetParamsSpeciesGrouping = "merlin"
	RefTaxonomySppgroupGetParamsSpeciesGroupingEbird  RefTaxonomySppgroupGetParamsSpeciesGrouping = "ebird"
)

func (r RefTaxonomySppgroupGetParamsSpeciesGrouping) IsKnown() bool {
	switch r {
	case RefTaxonomySppgroupGetParamsSpeciesGroupingMerlin, RefTaxonomySppgroupGetParamsSpeciesGroupingEbird:
		return true
	}
	return false
}
