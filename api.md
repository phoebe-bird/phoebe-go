# Data

## Observations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#Observation">Observation</a>

### Recent

Methods:

- <code title="get /data/obs/{regionCode}/recent">client.Data.Observations.Recent.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationRecentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationRecentListParams">DataObservationRecentListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#Observation">Observation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Notable

Methods:

- <code title="get /data/obs/{regionCode}/recent/notable">client.Data.Observations.Recent.Notable.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationRecentNotableService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationRecentNotableListParams">DataObservationRecentNotableListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#Observation">Observation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Species

Methods:

- <code title="get /data/obs/{regionCode}/recent/{speciesCode}">client.Data.Observations.Recent.Species.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationRecentSpecieService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, speciesCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationRecentSpecieGetParams">DataObservationRecentSpecieGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#Observation">Observation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Historic

Methods:

- <code title="get /data/obs/{regionCode}/historic/{y}/{m}/{d}">client.Data.Observations.Recent.Historic.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationRecentHistoricService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, y <a href="https://pkg.go.dev/builtin#int64">int64</a>, m <a href="https://pkg.go.dev/builtin#int64">int64</a>, d <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationRecentHistoricListParams">DataObservationRecentHistoricListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Geo

#### Recent

Methods:

- <code title="get /data/obs/geo/recent">client.Data.Observations.Geo.Recent.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationGeoRecentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationGeoRecentListParams">DataObservationGeoRecentListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#Observation">Observation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Species

Methods:

- <code title="get /data/obs/geo/recent/{speciesCode}">client.Data.Observations.Geo.Recent.Species.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationGeoRecentSpecieService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, speciesCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationGeoRecentSpecieListParams">DataObservationGeoRecentSpecieListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#Observation">Observation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Notable

Methods:

- <code title="get /data/obs/geo/recent/notable">client.Data.Observations.Geo.Recent.Notable.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationGeoRecentNotableService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationGeoRecentNotableListParams">DataObservationGeoRecentNotableListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Nearest

#### GeoSpecies

Methods:

- <code title="get /data/nearest/geo/recent/{speciesCode}">client.Data.Observations.Nearest.GeoSpecies.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationNearestGeoSpecieService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, speciesCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#DataObservationNearestGeoSpecieListParams">DataObservationNearestGeoSpecieListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Product

## Lists

Methods:

- <code title="get /product/lists/{regionCode}">client.Product.Lists.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductListService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductListGetParams">ProductListGetParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Historical

Methods:

- <code title="get /product/lists/{regionCode}/{y}/{m}/{d}">client.Product.Lists.Historical.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductListHistoricalService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, y <a href="https://pkg.go.dev/builtin#int64">int64</a>, m <a href="https://pkg.go.dev/builtin#int64">int64</a>, d <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductListHistoricalGetParams">ProductListHistoricalGetParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Top100

Methods:

- <code title="get /product/top100/{regionCode}/{y}/{m}/{d}">client.Product.Top100.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductTop100Service.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, y <a href="https://pkg.go.dev/builtin#int64">int64</a>, m <a href="https://pkg.go.dev/builtin#int64">int64</a>, d <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductTop100GetParams">ProductTop100GetParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Stats

Methods:

- <code title="get /product/stats/{regionCode}/{y}/{m}/{d}">client.Product.Stats.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductStatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, y <a href="https://pkg.go.dev/builtin#int64">int64</a>, m <a href="https://pkg.go.dev/builtin#int64">int64</a>, d <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## SpeciesList

Methods:

- <code title="get /product/spplist/{regionCode}">client.Product.SpeciesList.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductSpeciesListService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Checklist

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductChecklistViewResponse">ProductChecklistViewResponse</a>

Methods:

- <code title="get /product/checklist/view/{subId}">client.Product.Checklist.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductChecklistService.View">View</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#ProductChecklistViewResponse">ProductChecklistViewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Ref

## Region

### Adjacent

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionAdjacentListResponse">RefRegionAdjacentListResponse</a>

Methods:

- <code title="get /ref/adjacent/{regionCode}">client.Ref.Region.Adjacent.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionAdjacentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionAdjacentListResponse">RefRegionAdjacentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Info

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionInfoGetResponse">RefRegionInfoGetResponse</a>

Methods:

- <code title="get /ref/region/info/{regionCode}">client.Ref.Region.Info.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionInfoGetParams">RefRegionInfoGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionInfoGetResponse">RefRegionInfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### List

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionListListResponse">RefRegionListListResponse</a>

Methods:

- <code title="get /ref/region/list/{regionType}/{parentRegionCode}">client.Ref.Region.List.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionListService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionType <a href="https://pkg.go.dev/builtin#string">string</a>, parentRegionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionListListParams">RefRegionListListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefRegionListListResponse">RefRegionListListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Hotspot

Methods:

- <code title="get /ref/hotspot/{regionCode}">client.Ref.Hotspot.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, regionCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotListParams">RefHotspotListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Geo

Methods:

- <code title="get /ref/hotspot/geo">client.Ref.Hotspot.Geo.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotGeoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotGeoGetParams">RefHotspotGeoGetParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Info

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotInfoGetResponse">RefHotspotInfoGetResponse</a>

Methods:

- <code title="get /ref/hotspot/info/{locId}">client.Ref.Hotspot.Info.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, locID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefHotspotInfoGetResponse">RefHotspotInfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Taxonomy

### Ebird

Methods:

- <code title="get /ref/taxonomy/ebird">client.Ref.Taxonomy.Ebird.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyEbirdService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyEbirdGetParams">RefTaxonomyEbirdGetParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Forms

Methods:

- <code title="get /ref/taxon/forms/{speciesCode}">client.Ref.Taxonomy.Forms.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyFormService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, speciesCode <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Locales

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyLocaleListResponse">RefTaxonomyLocaleListResponse</a>

Methods:

- <code title="get /ref/taxa-locales/ebird">client.Ref.Taxonomy.Locales.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyLocaleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyLocaleListParams">RefTaxonomyLocaleListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyLocaleListResponse">RefTaxonomyLocaleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyVersionListResponse">RefTaxonomyVersionListResponse</a>

Methods:

- <code title="get /ref/taxonomy/versions">client.Ref.Taxonomy.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomyVersionListResponse">RefTaxonomyVersionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### SpeciesGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySpeciesGroupListResponse">RefTaxonomySpeciesGroupListResponse</a>

Methods:

- <code title="get /ref/sppgroup/{speciesGrouping}">client.Ref.Taxonomy.SpeciesGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySpeciesGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, speciesGrouping <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySpeciesGroupListParamsSpeciesGrouping">RefTaxonomySpeciesGroupListParamsSpeciesGrouping</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySpeciesGroupListParams">RefTaxonomySpeciesGroupListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go">phoebe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/phoebe-go#RefTaxonomySpeciesGroupListResponse">RefTaxonomySpeciesGroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
