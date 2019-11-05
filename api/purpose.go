package api

type Purpose struct {
	Purpose					string `json:"purpose"`
	PurposeCategory			int `json:"purposeCategory"`
	ConsentType				string `json:"consentType"`
	DataCategory			DataCategory `json:"piiCategory"`
	PrimaryPurpose			bool `json:"primaryPurpose"`
	Termination				string `json:"termination"`
	ThirdPartyDisclosure	bool `json:"thirdPartyDisclosure"`
	ThirdPartyName			string `json:"thirdPartyName,omitempty"`
}

func NewPurpose(purposeName string, termination string) *Purpose {
	return &Purpose{
		Purpose:				purposeName,
		ConsentType:			"EXPLICIT",
		ThirdPartyDisclosure:	false,
		Termination:			termination,
	}
}

func NewPurposeShared(purposeName string, termination string, thirdPartyName string) *Purpose {
	purpose := NewPurpose(purposeName, termination)

	purpose.ThirdPartyDisclosure = true
	purpose.ThirdPartyName = thirdPartyName

	return purpose
}
