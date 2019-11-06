package api

import "github.com/adaptant-labs/consent-receipt-go/api/purpose"

type Purpose struct {
	Purpose              string                       `json:"purpose"`
	PurposeCategory      purpose.PurposeSpecification `json:"purposeCategory"`
	ConsentType          string                       `json:"consentType"`
	DataCategory         DataCategory                 `json:"piiCategory"`
	PrimaryPurpose       bool                         `json:"primaryPurpose"`
	Termination          string                       `json:"termination"`
	ThirdPartyDisclosure bool                         `json:"thirdPartyDisclosure"`
	ThirdPartyName       string                       `json:"thirdPartyName,omitempty"`
}

func NewPurpose(spec purpose.PurposeSpecification, primaryPurpose bool, termination string) *Purpose {
	return NewPurposeDescription(spec.Description(), spec, primaryPurpose, termination)
}

func NewPurposeDescription(description string, spec purpose.PurposeSpecification, primaryPurpose bool, termination string) *Purpose {
	return &Purpose{
		Purpose:              description,
		PurposeCategory:      spec,
		ConsentType:          "EXPLICIT",
		PrimaryPurpose:       primaryPurpose,
		ThirdPartyDisclosure: false,
		Termination:          termination,
	}
}

func NewPurposeShared(description string, spec purpose.PurposeSpecification, primaryPurpose bool, termination string, thirdPartyName string) *Purpose {
	p := NewPurposeDescription(description, spec, primaryPurpose, termination)

	p.ThirdPartyDisclosure = true
	p.ThirdPartyName = thirdPartyName

	return p
}

func DefaultPurpose() *Purpose {
	return NewPurpose(purpose.CoreFunction, true, "Subscription end date + 1 year")
}
