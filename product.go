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
	Options []option.RequestOption
	// The data/obs end-points are used to fetch observations submitted to eBird in
	// checklists. There are two categories of end-point: 1. Fetch observations for a
	// specific country, region or location. 2. Fetch observations for nearby
	// locations - up to a distance of 50km. Each end-point supports optional query
	// parameters which allow you to filter the list of observations returned.
	Lists *ProductListService
	// The product end-points make it easy to get the information shown in various
	// pages on the eBird web site: 1. The Top 100 contributors on a given date. 2. The
	// checklists submitted on a given date. 3. The most recent checklists
	// submitted. 4. A summary of the checklists submitted on a given date. 5. The
	// details and all the observations of a checklist.
	Top100 *ProductTop100Service
	// The product end-points make it easy to get the information shown in various
	// pages on the eBird web site: 1. The Top 100 contributors on a given date. 2. The
	// checklists submitted on a given date. 3. The most recent checklists
	// submitted. 4. A summary of the checklists submitted on a given date. 5. The
	// details and all the observations of a checklist.
	Stats *ProductStatService
	// The product end-points make it easy to get the information shown in various
	// pages on the eBird web site: 1. The Top 100 contributors on a given date. 2. The
	// checklists submitted on a given date. 3. The most recent checklists
	// submitted. 4. A summary of the checklists submitted on a given date. 5. The
	// details and all the observations of a checklist.
	SpeciesList *ProductSpeciesListService
	// The product end-points make it easy to get the information shown in various
	// pages on the eBird web site: 1. The Top 100 contributors on a given date. 2. The
	// checklists submitted on a given date. 3. The most recent checklists
	// submitted. 4. A summary of the checklists submitted on a given date. 5. The
	// details and all the observations of a checklist.
	Checklist *ProductChecklistService
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
