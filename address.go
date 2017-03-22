package multivers

type Address struct {
	AddressesGuID  string   `json:"addressGuid"`
	AddressID      int      `json:"addressId"`
	AddressType    int      `json:"addressType"`
	Messages       []string `json:"messages"`
	CanChange      bool     `json:"canChange"`
	City           string   `json:"city"`
	ContactPerson  string   `json:"contactPerson"`
	CountryId      string   `json:"countryId"`
	Email          string   `json:"email"`
	Fax            string   `json:"fax"`
	FullAddress    string   `json:"fullAddress"`
	LanguageId     string   `json:"languageId"`
	Name           string   `json:"name"`
	OrganizationId int      `json:"organizationId"`
	Street1        string   `json:"street1"`
	Street2        string   `json:"street2"`
	Telephone      string   `json:"telephone"`
	ZipCode        string   `json:"zipCode"`
}