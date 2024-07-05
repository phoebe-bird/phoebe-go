// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebebird

import (
	"context"
	"net/http"

	"github.com/phoebe-bird/phoebe-go/internal/apijson"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefTaxonomyVersionService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefTaxonomyVersionService] method instead.
type RefTaxonomyVersionService struct {
	Options []option.RequestOption
}

// NewRefTaxonomyVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRefTaxonomyVersionService(opts ...option.RequestOption) (r *RefTaxonomyVersionService) {
	r = &RefTaxonomyVersionService{}
	r.Options = opts
	return
}

// Returns a list of all versions of the taxonomy, with a flag indicating which is
// the latest.
func (r *RefTaxonomyVersionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]RefTaxonomyVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ref/taxonomy/versions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RefTaxonomyVersionListResponse struct {
	AuthorityVer float64                            `json:"authorityVer"`
	Latest       bool                               `json:"latest"`
	JSON         refTaxonomyVersionListResponseJSON `json:"-"`
}

// refTaxonomyVersionListResponseJSON contains the JSON metadata for the struct
// [RefTaxonomyVersionListResponse]
type refTaxonomyVersionListResponseJSON struct {
	AuthorityVer apijson.Field
	Latest       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RefTaxonomyVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refTaxonomyVersionListResponseJSON) RawJSON() string {
	return r.raw
}
