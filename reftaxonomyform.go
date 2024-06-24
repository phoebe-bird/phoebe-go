// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefTaxonomyFormService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefTaxonomyFormService] method instead.
type RefTaxonomyFormService struct {
	Options []option.RequestOption
}

// NewRefTaxonomyFormService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRefTaxonomyFormService(opts ...option.RequestOption) (r *RefTaxonomyFormService) {
	r = &RefTaxonomyFormService{}
	r.Options = opts
	return
}

// For a species, get the list of subspecies recognised in the taxonomy. The
// results include the species that was passed in.
func (r *RefTaxonomyFormService) List(ctx context.Context, speciesCode string, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	if speciesCode == "" {
		err = errors.New("missing required speciesCode parameter")
		return
	}
	path := fmt.Sprintf("ref/taxon/forms/%s", speciesCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
