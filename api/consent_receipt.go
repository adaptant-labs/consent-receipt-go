package api

import (
	"github.com/adaptant-labs/consent-receipt-go/api/category"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"strings"
	"time"
)

type ConsentReceipt struct {
	Version             string                   `json:"version"`
	Jurisdiction        string                   `json:"jurisdiction"`
	ConsentTimestamp    int64                    `json:"consentTimestamp"`
	CollectionMethod    string                   `json:"collectionMethod"`
	ConsentReceiptID    string                   `json:"consentReceiptID"`
	PublicKey           string                   `json:"publicKey,omitempty"`
	Language            string                   `json:"language,omitempty"`
	SubjectID           string                   `json:"piiPrincipalId"`
	DataControllers     []*DataController        `json:"piiControllers"`
	PolicyUrl           string                   `json:"policyURL"`
	Services            []*Service               `json:"services,omitempty"`
	Sensitive           bool                     `json:"sensitive"`
	SensitiveCategories []*category.DataCategory `json:"spiCat,omitempty"`
}

func NewConsentReceipt() *ConsentReceipt {
	return &ConsentReceipt{
		Version:          "KI-CR-v1.1.0",
		ConsentTimestamp: time.Now().Unix(),
		ConsentReceiptID: uuid.New().String(),
		Language:         "EN",
		Sensitive:        false,
	}
}

func (cr *ConsentReceipt) AddDataController(controller *DataController) {
	cr.DataControllers = append(cr.DataControllers, controller)
}

func (cr *ConsentReceipt) AddService(service *Service) {
	cr.Services = append(cr.Services, service)
}

func (cr *ConsentReceipt) AddSensitiveCategory(category *category.DataCategory) {
	cr.Sensitive = true
	cr.SensitiveCategories = append(cr.SensitiveCategories, category)
}
func (cr *ConsentReceipt) GenerateJurisdictions() {
	// The simple case - a single controller serving a single country
	if len(cr.DataControllers) == 1 {
		cr.Jurisdiction = cr.DataControllers[0].Address.Country
		return
	}

	var jurisdictions []string

	// In the case of multiple controllers, each controller is a possible jurisdiction
	for _, val := range cr.DataControllers {
		jurisdictions = append(jurisdictions, val.Address.Country)
	}

	// Pull out the unique ones and return this in the expected format
	uniqueJurisdictions := uniqueStrings(jurisdictions)
	cr.Jurisdiction = strings.Join(uniqueJurisdictions, " ")
}

type ConsentReceiptClaims struct {
	ConsentReceipt *ConsentReceipt `json:"consentReceipt"`
	jwt.StandardClaims
}

func (cr ConsentReceipt) GenerateClaims() ConsentReceiptClaims {
	claims := ConsentReceiptClaims{
		ConsentReceipt: &cr,
		StandardClaims: jwt.StandardClaims{
			Issuer:   cr.DataControllers[0].ControllerName,
			IssuedAt: time.Now().Unix(),
		},
	}

	return claims
}
