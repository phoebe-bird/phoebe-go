// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefTaxonomyService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefTaxonomyService] method instead.
type RefTaxonomyService struct {
	Options       []option.RequestOption
	Ebird         *RefTaxonomyEbirdService
	Forms         *RefTaxonomyFormService
	Locales       *RefTaxonomyLocaleService
	Versions      *RefTaxonomyVersionService
	SpeciesGroups *RefTaxonomySpeciesGroupService
}

// NewRefTaxonomyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRefTaxonomyService(opts ...option.RequestOption) (r *RefTaxonomyService) {
	r = &RefTaxonomyService{}
	r.Options = opts
	r.Ebird = NewRefTaxonomyEbirdService(opts...)
	r.Forms = NewRefTaxonomyFormService(opts...)
	r.Locales = NewRefTaxonomyLocaleService(opts...)
	r.Versions = NewRefTaxonomyVersionService(opts...)
	r.SpeciesGroups = NewRefTaxonomySpeciesGroupService(opts...)
	return
}
