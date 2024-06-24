# Data

## Obs

### Recent

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentListResponse">DataObRecentListResponse</a>

Methods:

- <code title="get /data/obs/{regionCode}/recent">client.Data.Obs.Recent.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentListParams">DataObRecentListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentListResponse">DataObRecentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Notable

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentNotableListResponse">DataObRecentNotableListResponse</a>

Methods:

- <code title="get /data/obs/{regionCode}/recent/notable">client.Data.Obs.Recent.Notable.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentNotableService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentNotableListParams">DataObRecentNotableListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentNotableListResponse">DataObRecentNotableListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Species

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentSpecieGetResponse">DataObRecentSpecieGetResponse</a>

Methods:

- <code title="get /data/obs/{regionCode}/recent/{speciesCode}">client.Data.Obs.Recent.Species.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentSpecieService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, speciesCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentSpecieGetParams">DataObRecentSpecieGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObRecentSpecieGetResponse">DataObRecentSpecieGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Geo

#### Recent

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObGeoRecentListResponse">DataObGeoRecentListResponse</a>

Methods:

- <code title="get /data/obs/geo/recent">client.Data.Obs.Geo.Recent.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObGeoRecentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObGeoRecentListParams">DataObGeoRecentListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObGeoRecentListResponse">DataObGeoRecentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Species

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObGeoRecentSpecieListResponse">DataObGeoRecentSpecieListResponse</a>

Methods:

- <code title="get /data/obs/geo/recent/{speciesCode}">client.Data.Obs.Geo.Recent.Species.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObGeoRecentSpecieService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, speciesCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObGeoRecentSpecieListParams">DataObGeoRecentSpecieListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObGeoRecentSpecieListResponse">DataObGeoRecentSpecieListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Notable

Methods:

- <code title="get /data/obs/geo/recent/notable">client.Data.Obs.Geo.Recent.Notable.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObGeoRecentNotableService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObGeoRecentNotableListParams">DataObGeoRecentNotableListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Historic

Methods:

- <code title="get /data/obs/{regionCode}/historic/{y}/{m}/{d}">client.Data.Obs.Historic.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObHistoricService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, y <a href="https://pkg.go.dev/builtin#int64">int64</a>, m <a href="https://pkg.go.dev/builtin#int64">int64</a>, d <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObHistoricListParams">DataObHistoricListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Nearest

### Geo

#### Recent

##### Species

Methods:

- <code title="get /data/nearest/geo/recent/{speciesCode}">client.Data.Nearest.Geo.Recent.Species.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataNearestGeoRecentSpecieService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, speciesCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataNearestGeoRecentSpecieListParams">DataNearestGeoRecentSpecieListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Product

## Lists

### Region

Methods:

- <code title="get /product/lists/{regionCode}">client.Product.Lists.Region.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductListRegionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductListRegionGetParams">ProductListRegionGetParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Top100

Methods:

- <code title="get /product/top100/{regionCode}/{y}/{m}/{d}">client.Product.Top100.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductTop100Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, y <a href="https://pkg.go.dev/builtin#int64">int64</a>, m <a href="https://pkg.go.dev/builtin#int64">int64</a>, d <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductTop100ListParams">ProductTop100ListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Checklists

Methods:

- <code title="get /product/lists/{regionCode}/{y}/{m}/{d}">client.Product.Checklists.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductChecklistService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, y <a href="https://pkg.go.dev/builtin#int64">int64</a>, m <a href="https://pkg.go.dev/builtin#int64">int64</a>, d <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductChecklistGetParams">ProductChecklistGetParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Summary

Methods:

- <code title="get /product/stats/{regionCode}/{y}/{m}/{d}">client.Product.Checklists.Summary.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductChecklistSummaryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, y <a href="https://pkg.go.dev/builtin#int64">int64</a>, m <a href="https://pkg.go.dev/builtin#int64">int64</a>, d <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Species

Methods:

- <code title="get /product/spplist/{regionCode}">client.Product.Species.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductSpecieService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Checklist

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductChecklistGetResponse">ProductChecklistGetResponse</a>

Methods:

- <code title="get /product/checklist/view/{subId}">client.Product.Checklist.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductChecklistService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductChecklistGetResponse">ProductChecklistGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Ref

## Geo

### Adjacent

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefGeoAdjacentGetResponse">RefGeoAdjacentGetResponse</a>

Methods:

- <code title="get /ref/adjacent/{regionCode}">client.Ref.Geo.Adjacent.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefGeoAdjacentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefGeoAdjacentGetResponse">RefGeoAdjacentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Hotspots

Methods:

- <code title="get /ref/hotspot/{regionCode}">client.Ref.Hotspots.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotListParams">RefHotspotListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Geo

Methods:

- <code title="get /ref/hotspot/geo">client.Ref.Hotspots.Geo.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotGeoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotGeoGetParams">RefHotspotGeoGetParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Info

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotInfoGetResponse">RefHotspotInfoGetResponse</a>

Methods:

- <code title="get /ref/hotspot/info/{locId}">client.Ref.Hotspots.Info.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, locID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotInfoGetResponse">RefHotspotInfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Taxonomy

### Ebird

Methods:

- <code title="get /ref/taxonomy/ebird">client.Ref.Taxonomy.Ebird.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyEbirdService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyEbirdGetParams">RefTaxonomyEbirdGetParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Subspecies

Methods:

- <code title="get /ref/taxon/forms/{speciesCode}">client.Ref.Taxonomy.Subspecies.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySubspecieService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, speciesCode <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RefTaxonomy

## TaxaLocales

### Ebird

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyTaxaLocaleEbirdListResponse">RefTaxonomyTaxaLocaleEbirdListResponse</a>

Methods:

- <code title="get /ref/taxa-locales/ebird">client.RefTaxonomy.TaxaLocales.Ebird.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyTaxaLocaleEbirdService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyTaxaLocaleEbirdListParams">RefTaxonomyTaxaLocaleEbirdListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyTaxaLocaleEbirdListResponse">RefTaxonomyTaxaLocaleEbirdListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyVersionListResponse">RefTaxonomyVersionListResponse</a>

Methods:

- <code title="get /ref/taxonomy/versions">client.RefTaxonomy.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyVersionListResponse">RefTaxonomyVersionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sppgroup

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySppgroupGetResponse">RefTaxonomySppgroupGetResponse</a>

Methods:

- <code title="get /ref/sppgroup/{speciesGrouping}">client.RefTaxonomy.Sppgroup.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySppgroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, speciesGrouping <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySppgroupGetParamsSpeciesGrouping">RefTaxonomySppgroupGetParamsSpeciesGrouping</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySppgroupGetParams">RefTaxonomySppgroupGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySppgroupGetResponse">RefTaxonomySppgroupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RefRegion

## Info

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionInfoGetResponse">RefRegionInfoGetResponse</a>

Methods:

- <code title="get /ref/region/info/{regionCode}">client.RefRegion.Info.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionInfoGetParams">RefRegionInfoGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionInfoGetResponse">RefRegionInfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## List

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionListGetResponse">RefRegionListGetResponse</a>

Methods:

- <code title="get /ref/region/list/{regionType}/{parentRegionCode}">client.RefRegion.List.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionListService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionType <a href="https://pkg.go.dev/builtin#string">string</a>, parentRegionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionListGetParams">RefRegionListGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionListGetResponse">RefRegionListGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
