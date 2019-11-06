package category

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/camelcase"
	"log"
	"strconv"
	"strings"
)

type DataCategory int

const (
	categoryUndefined DataCategory = iota
	Biographical
	Contact
	Biometric
	SocialContact
	NetworkService
	Health
	Financial
	OfficialID
	SocialBenefitData
	JudicialData
	AssetData
	HRData
	MentalHealth
	Membership
	Behavioral
	Profiling
)

var toString = map[DataCategory]string{
	categoryUndefined: "",
	Biographical: "Biographical",
	Contact: "Contact",
	Biometric: "Biometric",
	SocialContact: "SocialContact",
	NetworkService: "NetworkService",
	Health: "Health",
	Financial: "Financial",
	OfficialID: "OfficialID",
	SocialBenefitData: "SocialBenefitData",
	JudicialData: "JudicialData",
	AssetData: "AssetData",
	HRData: "HRData",
	MentalHealth: "MentalHealth",
	Membership: "Membership",
	Behavioral: "Behavioral",
	Profiling: "Profiling",
}

func (c DataCategory) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(c.CategoryWithPrefix())
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (c *DataCategory) UnmarshalJSON(b []byte) error {
	var categoryStr string

	err := json.Unmarshal(b, &categoryStr)
	if err != nil {
		return err
	}

	*c = DataCategoryFromString(categoryStr)
	return nil
}

func (c DataCategory) Number() int {
	return int(c)
}

func (c DataCategory) Name() string {
	if name, ok := toString[c]; ok {
		return name
	}

	log.Printf("Unknown data category specified (%d)\n", c)
	return ""
}

func (c DataCategory) Category() string {
	name := c.Name()
	if name == "" {
		log.Printf("Unknown data category specified (%d)\n", c)
		return ""
	}

	switch c {
	case NetworkService:
		return "Network/Service"
	default:
		// CamelCase -> Camel Case conversion
		categoryString := camelcase.Split(name)
		return strings.Join(categoryString, " ")
	}
}

func (c DataCategory) CategoryWithPrefix() string {
	if _, ok := toString[c]; !ok {
		log.Printf("Unknown data category specified (%d)\n", c)
		return ""
	}

	return fmt.Sprintf("%d - %s", c.Number(), c.Category())
}

func (c DataCategory) String() string {
	return c.Name()
}

func DataCategoryFromString(categoryStr string) DataCategory {
	numStr := strings.Split(categoryStr, " - ")
	num, err := strconv.Atoi(numStr[0])
	if err != nil {
		log.Println("Failed to decode data category")
		return categoryUndefined
	}

	// Check if it's in range
	if _, ok := toString[DataCategory(num)]; ok {
		// Ok, valid, now convert the type.
		return DataCategory(num)
	}

	log.Printf("Invalid category specification (%d)\n", num)
	return categoryUndefined
}