package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type ConsentReceipt struct {
	Version				string `json:"version"`
	Jurisdiction		string `json:"jurisdiction"`
	ConsentTimestamp	int64 `json:"consentTimestamp"`
	CollectionMethod	string `json:"collectionMethod"`
	ConsentReceiptID	string `json:"consentReceiptID"`
	PublicKey			string `json:"publicKey,omitempty"`
	Language			string `json:"language,omitempty"`
	SubjectID			string `json:"piiPrincipalId"`
	DataControllers		[]*DataController `json:"piiControllers"`
	PolicyUrl			string `json:"policyURL"`
	Services			[]*Service `json:"services,omitempty"`
	Sensitive			bool `json:"sensitive"`
	DataCategories		[]*DataCategory `json:"spiCat,omitempty"`
}

func NewConsentReceipt() *ConsentReceipt {
	return &ConsentReceipt{
		Version:			"KI-CR-v1.1.0",
		ConsentTimestamp:	time.Now().Unix(),
		ConsentReceiptID:	uuid.New().String(),
		Language:			"EN",
		Sensitive:			false,
	}
}

func (cr *ConsentReceipt) AddDataController(controller *DataController) {
	cr.DataControllers = append(cr.DataControllers, controller)
}

func (cr *ConsentReceipt) AddService(service *Service) {
	cr.Services = append(cr.Services, service)
}

type ConsentReceiptClaims struct {
	ConsentReceipt *ConsentReceipt `json:"consentReceipt"`
	jwt.StandardClaims
}

func (cr ConsentReceipt) GenerateClaims() ConsentReceiptClaims {
	claims := ConsentReceiptClaims{
		ConsentReceipt: &cr,
		StandardClaims: jwt.StandardClaims{
			Issuer:    cr.DataControllers[0].ControllerName,
			IssuedAt:  time.Now().Unix(),
		},
	}

	return claims
}

