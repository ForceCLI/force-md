package accountForecastSettings

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "AccountForecastSettings"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// AccountForecastSettings holds the Manufacturing Cloud settings for the
// generation, display and adjustment of account forecasts. An org has one,
// stored as accountForecastSettings/AccountForecastSettings.accountForecastSetting.
type AccountForecastSettings struct {
	metadata.MetadataInfo
	XMLName                        xml.Name                 `xml:"AccountForecastSettings"`
	Xmlns                          string                   `xml:"xmlns,attr"`
	AccountFilterId                *TextLiteral             `xml:"accountFilterId"`
	AccountForecastFormulas        []AccountForecastFormula `xml:"accountForecastFormulas"`
	AcctPrdctPrdFrcstVolCnt        *IntegerText             `xml:"acctPrdctPrdFrcstVolCnt"`
	CalculationFrequency           *TextLiteral             `xml:"calculationFrequency"`
	DisplayDuration                *IntegerText             `xml:"displayDuration"`
	DisplayedForecastMetrics       *TextLiteral             `xml:"displayedForecastMetrics"`
	DisplayedRevenueMetrics        *TextLiteral             `xml:"displayedRevenueMetrics"`
	EditableAtStartOfPeriod        *BooleanText             `xml:"editableAtStartOfPeriod"`
	EditsAllowedFor                *IntegerText             `xml:"editsAllowedFor"`
	ForecastFrequency              *TextLiteral             `xml:"forecastFrequency"`
	ObjectMapping                  *ObjectMapping           `xml:"objectMapping"`
	OpportunityItemScheduleEnabled *BooleanText             `xml:"opportunityItemScheduleEnabled"`
	OpportunityProbabilityEnabled  *BooleanText             `xml:"opportunityProbabilityEnabled"`
	PrimaryNotifEmailAddress       *TextLiteral             `xml:"primaryNotifEmailAddress"`
	ProductFilterId                *TextLiteral             `xml:"productFilterId"`
	RecalculateAllFrcstCnt         *IntegerText             `xml:"recalculateAllFrcstCnt"`
	RegenerateForecastCnt          *IntegerText             `xml:"regenerateForecastCnt"`
	SalesAgreementFilterId         *TextLiteral             `xml:"salesAgreementFilterId"`
	SecondaryNotifEmailAddress     *TextLiteral             `xml:"secondaryNotifEmailAddress"`
	StartingPeriod                 *IntegerText             `xml:"startingPeriod"`
}

// AccountForecastFormula is the formula that calculates a quantity or
// revenue metric for the periods from StartingPeriod to EndingPeriod.
type AccountForecastFormula struct {
	EndingPeriod   *IntegerText `xml:"endingPeriod"`
	Formula        *TextLiteral `xml:"formula"`
	FormulaType    *TextLiteral `xml:"formulaType"`
	StartingPeriod *IntegerText `xml:"startingPeriod"`
}

// ObjectMapping maps fields of the settings' input object to fields of its
// output object.
type ObjectMapping struct {
	InputObject   *TextLiteral         `xml:"inputObject"`
	MappingFields []ObjectMappingField `xml:"mappingFields"`
	OutputObject  *TextLiteral         `xml:"outputObject"`
}

type ObjectMappingField struct {
	InputField  *TextLiteral `xml:"inputField"`
	OutputField *TextLiteral `xml:"outputField"`
}

func (s *AccountForecastSettings) SetMetadata(m metadata.MetadataInfo) {
	s.MetadataInfo = m
}

func (s *AccountForecastSettings) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*AccountForecastSettings, error) {
	s := &AccountForecastSettings{}
	return s, metadata.ParseMetadataXml(s, path)
}
