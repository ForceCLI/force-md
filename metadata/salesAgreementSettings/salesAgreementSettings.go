package salesAgreementSettings

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/accountForecastSettings"
)

const NAME = "SalesAgreementSettings"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// SalesAgreementSettings holds the Manufacturing Cloud settings for sales
// agreements: the agreement terms metrics, the actuals calculation, renewal
// and approval. An org has one, stored as
// salesAgreementSettings/SalesAgreementSettings.salesAgreementSetting.
type SalesAgreementSettings struct {
	metadata.MetadataInfo
	XMLName                        xml.Name                               `xml:"SalesAgreementSettings"`
	Xmlns                          string                                 `xml:"xmlns,attr"`
	ActualsCalculationMode         *TextLiteral                           `xml:"actualsCalculationMode"`
	ArePredfndStatusValOveride     *BooleanText                           `xml:"arePredfndStatusValOveride"`
	DecimalScale                   *IntegerText                           `xml:"decimalScale"`
	DisplayGroups                  []DisplayGroup                         `xml:"displayGroups"`
	DisplayedAgreementTermsMetrics *TextLiteral                           `xml:"displayedAgreementTermsMetrics"`
	FutureActCalcSchedules         *IntegerText                           `xml:"futureActCalcSchedules"`
	IsOnlyApprovalProcessUsed      *BooleanText                           `xml:"isOnlyApprovalProcessUsed"`
	MeasureDefinitions             []MeasureDefinition                    `xml:"measureDefinitions"`
	ObjectMapping                  *accountForecastSettings.ObjectMapping `xml:"objectMapping"`
	PrimaryNotifEmailAddress       *TextLiteral                           `xml:"primaryNotifEmailAddress"`
	RenewalPeriodDayCount          *IntegerText                           `xml:"renewalPeriodDayCount"`
	SecondaryNotifEmailAddress     *TextLiteral                           `xml:"secondaryNotifEmailAddress"`
	TransactionGraceDay            *IntegerText                           `xml:"transactionGraceDay"`
}

// DisplayGroup is a group of measures shown together in agreement terms.
type DisplayGroup struct {
	AdvAcctFrcstDisplayGroupName *TextLiteral       `xml:"advAcctFrcstDisplayGroupName"`
	DisplayGroupItems            []DisplayGroupItem `xml:"displayGroupItems"`
	DisplayGroupType             *TextLiteral       `xml:"displayGroupType"`
	IsDefault                    *BooleanText       `xml:"isDefault"`
	UserProfileName              *TextLiteral       `xml:"userProfileName"`
}

type DisplayGroupItem struct {
	AdvAcctFrcstDplyGroupItemName *TextLiteral `xml:"advAcctFrcstDplyGroupItemName"`
	DisplayOrder                  *IntegerText `xml:"displayOrder"`
	MeasureReferenceName          *TextLiteral `xml:"measureReferenceName"`
}

// MeasureDefinition defines a quantity or revenue measure of sales
// agreement schedules.
type MeasureDefinition struct {
	AdvAcctForecastMeasureDefName *TextLiteral `xml:"advAcctForecastMeasureDefName"`
	AggregationType               *TextLiteral `xml:"aggregationType"`
	ComputationMethod             *TextLiteral `xml:"computationMethod"`
	ForecastDataMeasureName       *TextLiteral `xml:"forecastDataMeasureName"`
	ForecastMeasureName           *TextLiteral `xml:"forecastMeasureName"`
	ForecastMeasureType           *TextLiteral `xml:"forecastMeasureType"`
	IsAdjustmentTracked           *BooleanText `xml:"isAdjustmentTracked"`
}

func (s *SalesAgreementSettings) SetMetadata(m metadata.MetadataInfo) {
	s.MetadataInfo = m
}

func (s *SalesAgreementSettings) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*SalesAgreementSettings, error) {
	s := &SalesAgreementSettings{}
	return s, metadata.ParseMetadataXml(s, path)
}
