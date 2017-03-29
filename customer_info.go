package multivers

type CustomerInfo struct {
	AccountManagerID    string               `json:"accountManagerId"`
	ApplyOrderSurcharge bool                 `json:"applyOrderSurcharge"`
	BusinessNumber      string               `json:"businessNumber"`
	ChargeVatTypeID     int                  `json:"chargeVatTypeId"`
	City                string               `json:"city"`
	CocCity             string               `json:"cocCity"`
	CocDate             *TimeWithoutTimeZone `json:"cocDate"`
	CocRegistration     string               `json:"cocRegistration"`
	ContactPerson       string               `json:"contactPerson"`
	CountryID           string               `json:"countryId"`
	CreditSqueezeID     string               `json:"creditSqueezeId"`
	CurrencyID          string               `json:"currencyId"`
	CustomerID          string               `json:"customerId"`
	CustomerStateID     string               `json:"customerStateId"`
	DateChanged         *TimeWithoutTimeZone `json:"dateChanged"`
	DateCreated         *TimeWithoutTimeZone `json:"dateCreated"`
	Email               string               `json:"email"`
	Fax                 string               `json:"fax"`
	Homepage            string               `json:"homepage"`
	IsDunForPayment     bool                 `json:"isDunForPayment"`
	LanguageID          string               `json:"languageId"`
	MobilePhone         string               `json:"mobilePhone"`
	Name                string               `json:"name"`
	OrganizationID      int                  `json:"organizationId"`
	PaymentConditionID  string               `json:"paymentConditionId"`
	RevenueAccountID    string               `json:"revenueAccountId"`
	ShortName           string               `json:"shortName"`
	Street1             string               `json:"street1"`
	Street2             string               `json:"street2"`
	Telephone           string               `json:"telephone"`
	VatNumber           string               `json:"vatNumber"`
	VatScenarioID       int                  `json:"vatScenarioId"`
	VatVerificationDate *TimeWithoutTimeZone `json:"vatVerificationDate"`
	ZipCode             string               `json:"zipCode"`
}
