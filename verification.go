package braintree

type Verification struct {
	XMLName                      string `xml:"verification"`
	AvsErrorResponseCode         string `xml:"avs-error-response-code"`
	AvsPostalCodeResponseCode    string `xml:"avs-postal-code-response-code"`
	AvsStreetAddressResponseCode string `xml:"avs-street-address-response-code"`
	CurrencyIsoCode              string `xml:"currency-iso-code"`
	CVVResponseCode              string `xml:"cvv-response-code"`
	ProcessorResponseCode        string `xml:"processor-response-code"`
	ProcessorResponseText        string `xml:"processor-response-text"`
	NetworkResponseCode          string `xml:"network-response-code"`
	NetworkResponseText          string `xml:"network-response-text"`
	NetworkTransactionID         string `xml:"network-transaction-id"`
	MerchantAccountID            string `xml:"merchant-account-id"`
	GraphqlID                    string `xml:"graphql-id"`
	ID                           string `xml:"id"`
}

type Verifications struct {
	Verification []*Verification `xml:"verification"`
}
