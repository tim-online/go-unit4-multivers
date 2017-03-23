package multivers

type AddressInfo struct {
	AddressGUID             string `json:"addressGuid"`
	AddressID               int    `json:"addressId"`
	AddressTypeID           string `json:"addressTypeId"`
	City                    string `json:"city"`
	Country                 string `json:"country"`
	FullAddress             string `json:"fullAddress"`
	GoogleMapsDirectionsUrl string `json:"googleMapsDirectionsUrl"`
	GoogleMapsUrl           string `json:"googleMapsUrl"`
	Name                    string `json:"name"`
	Street1                 string `json:"street1"`
	Street2                 string `json:"street2"`
	ZipCode                 string `json:"zipCode"`
}
