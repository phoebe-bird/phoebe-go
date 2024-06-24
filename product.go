// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"github.com/phoebe-bird/phoebe-go/option"
)

// ProductService contains methods and other services that help with interacting
// with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductService] method instead.
type ProductService struct {
	Options     []option.RequestOption
	Lists       *ProductListService
	Top100      *ProductTop100Service
	Stats       *ProductStatService
	SpeciesList *ProductSpeciesListService
	Checklist   *ProductChecklistService
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r *ProductService) {
	r = &ProductService{}
	r.Options = opts
	r.Lists = NewProductListService(opts...)
	r.Top100 = NewProductTop100Service(opts...)
	r.Stats = NewProductStatService(opts...)
	r.SpeciesList = NewProductSpeciesListService(opts...)
	r.Checklist = NewProductChecklistService(opts...)
	return
}
