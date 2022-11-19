package company

var requestFirstPart = (`[
		{
			"operationName": "GenericAnalyticsQuery",
			"variables": {},
			"query": "query GenericAnalyticsQuery {\n  context {\n    countryCode\n    __typename\n  }\n  analytics {\n    config {\n      debug\n      version\n      ga_codes\n      ga_plugins\n      ga_params\n      ym_codes\n      city\n      district\n      region_id\n      device_type\n      user_id\n      user_type\n      owner\n      chunk_name\n      __typename\n    }\n    __typename\n  }\n}"
		},
		{
			"operationName": "ListingPanelQuery",
			"variables": {
				"categoryId": 0,
				"controller": "search-listing-page",
				"path": "/search"
			},
			"query": "query ListingPanelQuery($categoryId: Int, $controller: String, $path: String) {\n  banners(\n    type: \"listing-panel\"\n    targeting: {category_id: $categoryId, controller: $controller, path: $path}\n  ) {\n    id\n    imageId\n    imageUrl(width: 2048, height: 2048)\n    text\n    textColor\n    gradientStart\n    gradientStop\n    url\n    openInNewTab\n    __typename\n  }\n}"
		},
		{
			"operationName": "PortableSearchFiltersQuery",
			"variables": {
				"regionId": null,
				"search_term": "`)

var requestSecondPart = (`",
	"params": {
		"binary_filters": []
	},
	"limit": 48,
	"offset": 0
},
"query": "query PortableSearchFiltersQuery($search_term: String!, $params: Any, $company_id: Int, $sort: String, $regionId: Int = null) {\n  listing: searchListing(\n    search_term: $search_term\n    params: $params\n    company_id: $company_id\n    sort: $sort\n    region: {id: $regionId}\n  ) {\n    filters {\n      ...FiltersFragment\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment FiltersFragment on ListingFilters {\n  total\n  priceChartFilter {\n    ...PriceFilterFragment\n    __typename\n  }\n  binaryFilters {\n    ...PromoFilterFragment\n    __typename\n  }\n  attributeFilters {\n    ...AttributeFilterFragment\n    __typename\n  }\n  categoryFilter {\n    ...AttributeFilterFragment\n    __typename\n  }\n  companyFilter {\n    ...AttributeFilterFragment\n    __typename\n  }\n  productGroupFilter {\n    ...AttributeFilterFragment\n    __typename\n  }\n  deliveryFilter {\n    ...DeliveryFilterFragment\n    __typename\n  }\n  paymentFilter {\n    ...PaymentFilterFragment\n    __typename\n  }\n  colorFilter {\n    ...AttributeFilterFragment\n    __typename\n  }\n  __typename\n}\n\nfragment PriceFilterFragment on PriceChartFilter {\n  measureUnit\n  values\n  __typename\n}\n\nfragment PromoFilterFragment on Filter {\n  name\n  values {\n    selected\n    value\n    count\n    title\n    __typename\n  }\n  __typename\n}\n\nfragment AttributeFilterFragment on AttributeFilter {\n  name\n  title\n  type\n  min\n  max\n  measureUnit\n  sorting\n  previewSorting\n  values {\n    selected\n    value\n    count\n    title\n    position\n    parent\n    used_count\n    __typename\n  }\n  __typename\n}\n\nfragment DeliveryFilterFragment on Filter {\n  name\n  values {\n    selected\n    value\n    count\n    title\n    __typename\n  }\n  __typename\n}\n\nfragment PaymentFilterFragment on Filter {\n  name\n  values {\n    selected\n    value\n    count\n    title\n    __typename\n  }\n  __typename\n}"
},
{
"operationName": "PromoPopupQuery",
"variables": {
	"categoryId": 0,
	"controller": "search-listing-page",
	"path": "/search"
},
"query": "query PromoPopupQuery($categoryId: Int, $controller: String, $path: String) {\n  banners(\n    type: \"promo-popup\"\n    targeting: {category_id: $categoryId, controller: $controller, path: $path}\n  ) {\n    id\n    imageId\n    imageUrl(width: 2048, height: 2048)\n    text\n    textColor\n    gradientStart\n    gradientStop\n    url\n    openInNewTab\n    __typename\n  }\n}"
},
{
"operationName": "BesidaDataQuery",
"variables": {},
"query": "query BesidaDataQuery {\n  context {\n    currentUser {\n      id\n      evo_id\n      has_verified_phone\n      company_id\n      company {\n        id\n        premiumServiceId\n        __typename\n      }\n      __typename\n    }\n    besidaUrl\n    isBesidaEnabled\n    socialLinks\n    isSudo\n    defaultCurrencyText\n    __typename\n  }\n}"
},
{
"operationName": "PromoPopupQuery",
"variables": {
	"controller": "search-listing-page",
	"path": "/search"
},
"query": "query PromoPopupQuery($categoryId: Int, $controller: String, $path: String) {\n  banners(\n    type: \"promo-popup\"\n    targeting: {category_id: $categoryId, controller: $controller, path: $path}\n  ) {\n    id\n    imageId\n    imageUrl(width: 2048, height: 2048)\n    text\n    textColor\n    gradientStart\n    gradientStop\n    url\n    openInNewTab\n    __typename\n  }\n}"
}
]`)

type Values struct {
	Value int    `json:"value"`
	Title string `json:"title"`
}

type CompanyFilter struct {
	Values []Values `json:"values"`
}

type Filters struct {
	CompanyFilter CompanyFilter `json:"companyFilter"`
}

type Listing struct {
	Filters Filters `json:"filters"`
}

type Data struct {
	Listing Listing `json:"listing"`
}

type ResponseJson struct {
	Data Data `json:"data"`
}
