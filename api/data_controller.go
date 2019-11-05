package api

type PostalAddress struct {
	Country				string `json:"addressCountry" mapstructure:"country"`
	Locality			string `json:"addressLocality" mapstructure:"city"`
	Region				string `json:"addressRegion,omitempty" mapstructure:"region"`
	PostOfficeBoxNumber	string `json:"postOfficeBoxNumber,omitempty" mapstructure:"pobox"`
	PostalCode			string `json:"postalCode" mapstructure:"postalcode"`
	StreetAddress		string `json:"streetAddress" mapstructure:"address"`
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
	ControllerName		string `json:"piiController" mapstructure:"name"`
	OnBehalf			bool `json:"on_behalf,omitempty" mapstructure:"onbehalf"`
	Contact				string `json:"contact" mapstructure:"contact"`
	Address				PostalAddress `json:"address" mapstructure:"address"`
	Email				string `json:"email" mapstructure:"email"`
	Phone				string `json:"phone" mapstructure:"phone"`
	ControllerUrl		string `json:"piiControllerUrl,omitempty" mapstructure:"url"`
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

func (dc *DataController) NewConsentReceipt() *ConsentReceipt {
	cr := NewConsentReceipt()

	cr.AddDataController(dc)
	cr.Jurisdiction = dc.Address.Country

	return cr
}