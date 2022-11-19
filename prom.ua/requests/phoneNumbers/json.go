package phonenumbers

var requestFirstPart = (`[{
	"operationName": "companyContactsQuery",
	"variables": {
		"withGroupManagerPhones": false,
		"withWorkingHoursWarning": false,
		"getProductDetails": false,
		"company_id":
	`)

var requestSecondPart = (`},
	"query": "query companyContactsQuery($company_id: Int!, $groupId: Int = null, $productId: Int = null, $withGroupManagerPhones: Boolean = false, $withWorkingHoursWarning: Boolean = false, $getProductDetails: Boolean = false) {\n  context {\n    context_meta\n    currentRegionId\n    recaptchaToken\n    __typename\n  }\n  company(id: $company_id) {\n    ...CompanyWorkingHoursFragment @include(if: $withWorkingHoursWarning)\n    id\n    name\n    contactPerson\n    contactEmail\n    phones\n    addressText\n    addressId\n    isChatVisible\n    mainLogoUrl(width: 100, height: 50)\n    urlForCompanyProducts\n    isOneClickOrderAllowed\n    isOrderableInCatalog\n    isPackageCPA\n    region {\n      id\n      __typename\n    }\n    geoCoordinates {\n      latitude\n      longtitude\n      __typename\n    }\n    branches {\n      id\n      name\n      phones\n      address {\n        region_id\n        country_id\n        city\n        zipCode\n        street\n        regionText\n        __typename\n      }\n      __typename\n    }\n    shouldShowOnMap\n    webSiteUrl\n    site {\n      isDisabled\n      __typename\n    }\n    __typename\n  }\n  productGroup(id: $groupId) @include(if: $withGroupManagerPhones) {\n    managerPhones\n    __typename\n  }\n  product(id: $productId) @include(if: $getProductDetails) {\n    id\n    name\n    image(width: 60, height: 60)\n    price\n    priceCurrencyLocalized\n    buyButtonDisplayType\n    regions {\n      id\n      name\n      isCity\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CompanyWorkingHoursFragment on Company {\n  id\n  isWorkingNow\n  isOrderableInCatalog\n  scheduleSettings {\n    currentDayCaption\n    __typename\n  }\n  scheduleDays {\n    name\n    dayType\n    hasBreak\n    workTimeRangeStart\n    workTimeRangeEnd\n    breakTimeRangeStart\n    breakTimeRangeEnd\n    __typename\n  }\n  __typename\n}"
}]`)

type CompanyInfo struct {
	Data struct {
		Company struct {
			Name   string
			Phones []struct {
				Description string `json:"description"`
				Number      string `json:"number"`
			} `json:"phones"`
		} `json:"company"`
	} `json:"data"`
}
