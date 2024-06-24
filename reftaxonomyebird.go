// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"net/http"
	"net/url"

	"github.com/phoebe-bird/phoebe-go/internal/apiquery"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefTaxonomyEbirdService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefTaxonomyEbirdService] method instead.
type RefTaxonomyEbirdService struct {
	Options []option.RequestOption
}

// NewRefTaxonomyEbirdService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRefTaxonomyEbirdService(opts ...option.RequestOption) (r *RefTaxonomyEbirdService) {
	r = &RefTaxonomyEbirdService{}
	r.Options = opts
	return
}

// Get the taxonomy used by eBird. #### Notes Each entry in the taxonomy contains a
// species code for example, barswa for Barn Swallow. You can download the taxonomy
// for selected species using the _species_ query parameter with a comma separating
// each code. Otherwise the full taxonomy is downloaded.
func (r *RefTaxonomyEbirdService) Get(ctx context.Context, query RefTaxonomyEbirdGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "ref/taxonomy/ebird"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RefTaxonomyEbirdGetParams struct {
	// Only fetch records from these taxonomic categories.
	Cat param.Field[string] `query:"cat"`
	// Fetch the records in CSV or JSON format.
	Fmt param.Field[RefTaxonomyEbirdGetParamsFmt] `query:"fmt"`
	// Use this language for common names.
	Locale param.Field[string] `query:"locale"`
	// Only fetch records for these species.
	Species param.Field[string] `query:"species"`
	// Fetch a specific version of the taxonomy.
	Version param.Field[string] `query:"version"`
}

// URLQuery serializes [RefTaxonomyEbirdGetParams]'s query parameters as
// `url.Values`.
func (r RefTaxonomyEbirdGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Fetch the records in CSV or JSON format.
type RefTaxonomyEbirdGetParamsFmt string

const (
	RefTaxonomyEbirdGetParamsFmtCsv  RefTaxonomyEbirdGetParamsFmt = "csv"
	RefTaxonomyEbirdGetParamsFmtJson RefTaxonomyEbirdGetParamsFmt = "json"
)

func (r RefTaxonomyEbirdGetParamsFmt) IsKnown() bool {
	switch r {
	case RefTaxonomyEbirdGetParamsFmtCsv, RefTaxonomyEbirdGetParamsFmtJson:
		return true
	}
	return false
}
