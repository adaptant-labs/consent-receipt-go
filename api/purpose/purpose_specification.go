// Lack of explicitly namespaced enums in Go mean we simply hide this in another package, instead
package purpose

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/camelcase"
	"log"
	"strconv"
	"strings"
)

type PurposeSpecification int

const (
	purposeUndefined PurposeSpecification = iota
	CoreFunction
	ContractedService
	Delivery
	ContactRequested
	PersonalizedExperience
	Marketing
	MarketingThirdParties
	SharingForDelivery
	SharingForMarketing
	ThirdPartySharingForCoreFunction
	ThirdPartySharingForOthers
	LegallyRequiredDataRetention
	RequiredByLawEnforcementOrGovernment
	ProtectingYourHealth
	ProtectingOurInterests
	ImprovePerformance
	maxPurposes
)

func (p PurposeSpecification) MarshalJSON() ([]byte, error) {
	fmt.Printf("Marshalling JSON: %+v\n", p)
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(p.PurposeWithPrefix())
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (p *PurposeSpecification) UnmarshalJSON(b []byte) error {
	var purposeStr string

	err := json.Unmarshal(b, &purposeStr)
	if err != nil {
		return err
	}

	*p = PurposeSpecificationFromString(purposeStr)
	return nil
}

func (p PurposeSpecification) String() string {
	return p.Name()
}

func (p PurposeSpecification) Number() int {
	return int(p)
}

func (p PurposeSpecification) Name() string {
	switch p {
	case CoreFunction:
		return "CoreFunction"
	case ContractedService:
		return "ContractedService"
	case Delivery:
		return "Delivery"
	case ContactRequested:
		return "ContactRequested"
	case PersonalizedExperience:
		return "PersonalizedExperience"
	case Marketing:
		return "Marketing"
	case MarketingThirdParties:
		return "MarketingThirdParties"
	case SharingForDelivery:
		return "SharingForDelivery"
	case SharingForMarketing:
		return "SharingForMarketing"
	case ThirdPartySharingForCoreFunction:
		return "ThirdPartySharingForCoreFunction"
	case ThirdPartySharingForOthers:
		return "ThirdPartySharingForOthers"
	case LegallyRequiredDataRetention:
		return "LegallyRequiredDataRetention"
	case RequiredByLawEnforcementOrGovernment:
		return "RequiredByLawEnforcementOrGovernment"
	case ProtectingYourHealth:
		return "ProtectingYourHealth"
	case ProtectingOurInterests:
		return "ProtectingOurInterests"
	case ImprovePerformance:
		return "ImprovePerformance"
	}

	return ""
}

func (p PurposeSpecification) Description() string {
	switch p {
	case CoreFunction:
		return "Enabling us to carry out the core functions of our site/app/services";
	case ContractedService:
		return "Providing contracted or requested services to you.";
	case Delivery:
		return "Delivering physical goods to you.";
	case ContactRequested:
		return "Communicating with you about information or services you specifically request.";
	case PersonalizedExperience:
		return "Providing you with a personalized experience of our site/app/service.";
	case Marketing:
		return "Communicating with you about our other services you may be interested in.";
	case MarketingThirdParties:
		return "Communicating with you about the services of third parties you may be interested in.";
	case SharingForDelivery:
		return "Providing the information to third parties to deliver our services on our behalf.";
	case SharingForMarketing:
		return "Providing the information to third parties to enable them to communicate with you about their own services you may be interested in.";
	case ThirdPartySharingForCoreFunction:
		return "Providing the information to third parties to enable them to deliver or improve their own services to you.";
	case ThirdPartySharingForOthers:
		return "Providing the information to third parties to enable them to deliver or improve their own services to others.";
	case LegallyRequiredDataRetention:
		return "Complying with our legal obligations for record keeping.";
	case RequiredByLawEnforcementOrGovernment:
		return "Complying with our legal obligations to provide the information to law enforcement or other regulatory/government bodies.";
	case ProtectingYourHealth:
		return "Protecting your vital and health interests.";
	case ProtectingOurInterests:
		return "Protecting our legitimate interests, yours or those of a third party.";
	case ImprovePerformance:
		return "Measure or improve our performance or the delivery of our services.";
	}

	return ""
}

func (p PurposeSpecification) Purpose() string {
	switch p {
	case SharingForDelivery:
		return "Sharing for Delivery";
	case SharingForMarketing:
		return "Sharing for Marketing";
	case ThirdPartySharingForCoreFunction:
		return "3rd Party Sharing for Core Function";
	case ThirdPartySharingForOthers:
		return "3rd Party Sharing for";
	case RequiredByLawEnforcementOrGovernment:
		return "Required by Law Enforcement or Government";
	default:
		// CamelCase -> Camel Case conversion
		purposeString := camelcase.Split(p.Name())
		return strings.Join(purposeString, " ")
	}
}

func (p PurposeSpecification) PurposeWithPrefix() string {
	if p == purposeUndefined {
		return ""
	}

	return fmt.Sprintf("%d - %s", p.Number(), p.Purpose())
}

func PurposeSpecificationFromString(purposeStr string) PurposeSpecification {
	numStr := strings.Split(purposeStr, " - ")
	num, err := strconv.Atoi(numStr[0])
	if err != nil {
		log.Println("Failed to decode purpose specification")
		return purposeUndefined
	}

	// Check if it's in range
	if num > int(purposeUndefined) && num < int(maxPurposes) {
		// Ok, valid, now convert the type.
		return PurposeSpecification(num)
	}

	log.Printf("Invalid purpose specification (%d)\n", num)
	return purposeUndefined
}