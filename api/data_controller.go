package api

type PostalAddress struct {
	Country				string `json:"addressCountry"`
	Locality			string `json:"addressLocality"`
	Region				string `json:"addressRegion,omitempty"`
	PostOfficeBoxNumber	string `json:"postOfficeBoxNumber,omitempty"`
	PostalCode			string `json:"postalCode"`
	StreetAddress		string `json:"streetAddress"`
}

func NewPostalAddress(country string, locality string, postalcode string, address string) PostalAddress {
	return PostalAddress{
		Country:		country,
		Locality:		locality,
		PostalCode:		postalcode,
		StreetAddress:	address,
	}
}

type DataController struct {
	ControllerName		string `json:"piiController"`
	OnBehalf			bool `json:"on_behalf,omitempty"`
	Contact				string `json:"contact"`
	Address				PostalAddress `json:"address"`
	Email				string `json:"email"`
	Phone				string `json:"phone"`
	ControllerUrl		string `json:"piiControllerUrl,omitempty"`
}

func NewDataController(controllerName string, contactName string, email string, phone string, address PostalAddress) *DataController {
	return &DataController{
		ControllerName:	controllerName,
		OnBehalf:		false,
		Contact:		contactName,
		Address:		address,
		Email:			email,
		Phone:			phone,
	}
}
