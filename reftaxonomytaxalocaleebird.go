// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/phoebe-go/internal/apijson"
	"github.com/stainless-sdks/phoebe-go/internal/param"
	"github.com/stainless-sdks/phoebe-go/internal/requestconfig"
	"github.com/stainless-sdks/phoebe-go/option"
)

// RefTaxonomyTaxaLocaleEbirdService contains methods and other services that help
// with interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefTaxonomyTaxaLocaleEbirdService] method instead.
type RefTaxonomyTaxaLocaleEbirdService struct {
	Options []option.RequestOption
}

// NewRefTaxonomyTaxaLocaleEbirdService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRefTaxonomyTaxaLocaleEbirdService(opts ...option.RequestOption) (r *RefTaxonomyTaxaLocaleEbirdService) {
	r = &RefTaxonomyTaxaLocaleEbirdService{}
	r.Options = opts
	return
}

// Returns the list of supported locale codes and names for species common names,
// with the last time they were updated. Use the accept-language header to get
// translated language names when available.
//
// NOTE: The locale codes and names are stable but the other fields in this result
// are not yet finalized and should be used with caution.
func (r *RefTaxonomyTaxaLocaleEbirdService) List(ctx context.Context, query RefTaxonomyTaxaLocaleEbirdListParams, opts ...option.RequestOption) (res *[]RefTaxonomyTaxaLocaleEbirdListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ref/taxa-locales/ebird"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RefTaxonomyTaxaLocaleEbirdListResponse struct {
	Code        string                                     `json:"code"`
	LastUpdated string                                     `json:"lastUpdated"`
	Name        string                                     `json:"name"`
	JSON        refTaxonomyTaxaLocaleEbirdListResponseJSON `json:"-"`
}

// refTaxonomyTaxaLocaleEbirdListResponseJSON contains the JSON metadata for the
// struct [RefTaxonomyTaxaLocaleEbirdListResponse]
type refTaxonomyTaxaLocaleEbirdListResponseJSON struct {
	Code        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefTaxonomyTaxaLocaleEbirdListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refTaxonomyTaxaLocaleEbirdListResponseJSON) RawJSON() string {
	return r.raw
}

type RefTaxonomyTaxaLocaleEbirdListParams struct {
	AcceptLanguage param.Field[string] `header:"Accept-Language"`
}
