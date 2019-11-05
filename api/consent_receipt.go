package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type ConsentReceipt struct {
	Jurisdiction		string `json:"jurisdiction"`
	ConsentReceiptID	string `json:"consentReceiptID"`
	DataControllers		[]*DataController `json:"piiControllers"`
}

func NewConsentReceipt() *ConsentReceipt {
	return &ConsentReceipt{
		ConsentReceiptID: uuid.New().String(),
	}
}

func (cr *ConsentReceipt) AddDataController(controller *DataController) {
	cr.DataControllers = append(cr.DataControllers, controller)
}

type ConsentReceiptClaims struct {
	ConsentReceipt *ConsentReceipt `json:"consentReceipt"`
	jwt.StandardClaims
}

func (cr ConsentReceipt) GenerateClaims() ConsentReceiptClaims {
	now := time.Now().Unix()

	claims := ConsentReceiptClaims{
		ConsentReceipt: &cr,
		StandardClaims: jwt.StandardClaims{
			Issuer:    cr.DataControllers[0].ControllerName,
			IssuedAt:  now,
		},
	}

	return claims
}

