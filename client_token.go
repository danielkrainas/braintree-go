package braintree

type ClientTokenRequest struct {
	XMLName    string `xml:"client-token"`
	CustomerID string `xml:"customerId,omitempty"`
	Version    int    `xml:"version"`
}
