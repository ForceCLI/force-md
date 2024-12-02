package settings

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

func init() {
	open := func(path string) (metadata.RegisterableMetadata, error) { return Open(path) }
	internal.TypeRegistry.Register("AccountIntelligenceSettings", open)
	internal.TypeRegistry.Register("AccountSettings", open)
	internal.TypeRegistry.Register("ActionsSettings", open)
	internal.TypeRegistry.Register("ActivitiesSettings", open)
	internal.TypeRegistry.Register("AddressSettings", open)
	internal.TypeRegistry.Register("Ai4mSettings", open)
	internal.TypeRegistry.Register("AnalyticsSettings", open)
	internal.TypeRegistry.Register("ApexSettings", open)
	internal.TypeRegistry.Register("AppAnalyticsSettings", open)
	internal.TypeRegistry.Register("AppExperienceSettings", open)
	internal.TypeRegistry.Register("AutomatedContactsSettings", open)
	internal.TypeRegistry.Register("BlockchainSettings", open)
	internal.TypeRegistry.Register("BotSettings", open)
	internal.TypeRegistry.Register("BusinessHoursSettings", open)
	internal.TypeRegistry.Register("CampaignSettings", open)
	internal.TypeRegistry.Register("CaseSettings", open)
	internal.TypeRegistry.Register("ChatterAnswersSettings", open)
	internal.TypeRegistry.Register("ChatterEmailsMDSettings", open)
	internal.TypeRegistry.Register("ChatterSettings", open)
	internal.TypeRegistry.Register("CodeBuilderSettings", open)
	internal.TypeRegistry.Register("CommerceSettings", open)
	internal.TypeRegistry.Register("CommunitiesSettings", open)
	internal.TypeRegistry.Register("CompanySettings", open)
	internal.TypeRegistry.Register("ConnectedAppSettings", open)
	internal.TypeRegistry.Register("ContentSettings", open)
	internal.TypeRegistry.Register("ContractSettings", open)
	internal.TypeRegistry.Register("ConversationalIntelligenceSettings", open)
	internal.TypeRegistry.Register("CurrencySettings", open)
	internal.TypeRegistry.Register("CustomAddressFieldSettings", open)
	internal.TypeRegistry.Register("CustomerDataPlatformSettings", open)
	internal.TypeRegistry.Register("CustomizablePropensityScoringSettings", open)
	internal.TypeRegistry.Register("DeploymentSettings", open)
	internal.TypeRegistry.Register("DevHubSettings", open)
	internal.TypeRegistry.Register("DiscoverySettings", open)
	internal.TypeRegistry.Register("DocumentChecklistSettings", open)
	internal.TypeRegistry.Register("DynamicFormsSettings", open)
	internal.TypeRegistry.Register("EACSettings", open)
	internal.TypeRegistry.Register("EinsteinAgentSettings", open)
	internal.TypeRegistry.Register("EinsteinDocumentCaptureSettings", open)
	internal.TypeRegistry.Register("EmailAdministrationSettings", open)
	internal.TypeRegistry.Register("EmailIntegrationSettings", open)
	internal.TypeRegistry.Register("EmailTemplateSettings", open)
	internal.TypeRegistry.Register("EmployeeFieldAccessSettings", open)
	internal.TypeRegistry.Register("EmployeeUserSettings", open)
	internal.TypeRegistry.Register("EncryptionKeySettings", open)
	internal.TypeRegistry.Register("EnhancedNotesSettings", open)
	internal.TypeRegistry.Register("EntitlementSettings", open)
	internal.TypeRegistry.Register("EssentialsSettings", open)
	internal.TypeRegistry.Register("EventSettings", open)
	internal.TypeRegistry.Register("ExperienceBundleSettings", open)
	internal.TypeRegistry.Register("ExternalClientAppSettings", open)
	internal.TypeRegistry.Register("FilesConnectSettings", open)
	internal.TypeRegistry.Register("FileUploadAndDownloadSecuritySettings", open)
	internal.TypeRegistry.Register("FlowSettings", open)
	internal.TypeRegistry.Register("ForecastingSettings", open)
	internal.TypeRegistry.Register("FormulaSettings", open)
	internal.TypeRegistry.Register("GoogleAppsSettings", open)
	internal.TypeRegistry.Register("HighVelocitySalesSettings", open)
	internal.TypeRegistry.Register("IdeasSettings", open)
	internal.TypeRegistry.Register("IdentityProviderSettings", open)
	internal.TypeRegistry.Register("IncidentMgmtSettings", open)
	internal.TypeRegistry.Register("IndustriesEinsteinFeatureSettings", open)
	internal.TypeRegistry.Register("IndustriesManufacturingSettings", open)
	internal.TypeRegistry.Register("IndustriesSettings", open)
	internal.TypeRegistry.Register("InvocableActionSettings", open)
	internal.TypeRegistry.Register("KnowledgeGenerationSettings", open)
	internal.TypeRegistry.Register("KnowledgeSettings", open)
	internal.TypeRegistry.Register("LanguageSettings", open)
	internal.TypeRegistry.Register("LeadConfigSettings", open)
	internal.TypeRegistry.Register("LightningExperienceSettings", open)
	internal.TypeRegistry.Register("LiveAgentSettings", open)
	internal.TypeRegistry.Register("MacroSettings", open)
	internal.TypeRegistry.Register("MailMergeSettings", open)
	internal.TypeRegistry.Register("MapsAndLocationSettings", open)
	internal.TypeRegistry.Register("MeetingsSettings", open)
	internal.TypeRegistry.Register("MobileSettings", open)
	internal.TypeRegistry.Register("MyDomainSettings", open)
	internal.TypeRegistry.Register("NameSettings", open)
	internal.TypeRegistry.Register("NotificationsSettings", open)
	internal.TypeRegistry.Register("OauthOidcSettings", open)
	internal.TypeRegistry.Register("ObjectLinkingSettings", open)
	internal.TypeRegistry.Register("OmniChannelPricingSettings", open)
	internal.TypeRegistry.Register("OmniChannelSettings", open)
	internal.TypeRegistry.Register("OnlineSalesSettings", open)
	internal.TypeRegistry.Register("OpportunityScoreSettings", open)
	internal.TypeRegistry.Register("OpportunitySettings", open)
	internal.TypeRegistry.Register("OrderSettings", open)
	internal.TypeRegistry.Register("OrgSettings", open)
	internal.TypeRegistry.Register("PardotEinsteinSettings", open)
	internal.TypeRegistry.Register("PardotSettings", open)
	internal.TypeRegistry.Register("PartyDataModelSettings", open)
	internal.TypeRegistry.Register("PathAssistantSettings", open)
	internal.TypeRegistry.Register("PaymentsManagementEnabledSettings", open)
	internal.TypeRegistry.Register("PicklistSettings", open)
	internal.TypeRegistry.Register("PlatformEncryptionSettings", open)
	internal.TypeRegistry.Register("PlatformEventSettings", open)
	internal.TypeRegistry.Register("PortalsSettings", open)
	internal.TypeRegistry.Register("PredictionBuilderSettings", open)
	internal.TypeRegistry.Register("PrivacySettings", open)
	internal.TypeRegistry.Register("ProductSettings", open)
	internal.TypeRegistry.Register("QuickTextSettings", open)
	internal.TypeRegistry.Register("QuoteSettings", open)
	internal.TypeRegistry.Register("RealTimeEventSettings", open)
	internal.TypeRegistry.Register("RecommendationBuilderSettings", open)
	internal.TypeRegistry.Register("RecordPageSettings", open)
	internal.TypeRegistry.Register("RetailExecutionSettings", open)
	internal.TypeRegistry.Register("SandboxSettings", open)
	internal.TypeRegistry.Register("SceGlobalModelOptOutSettings", open)
	internal.TypeRegistry.Register("SchemaSettings", open)
	internal.TypeRegistry.Register("SearchSettings", open)
	internal.TypeRegistry.Register("SecuritySettings", open)
	internal.TypeRegistry.Register("ServiceCloudVoiceSettings", open)
	internal.TypeRegistry.Register("ServiceSetupAssistantSettings", open)
	internal.TypeRegistry.Register("SharingSettings", open)
	internal.TypeRegistry.Register("SiteSettings", open)
	internal.TypeRegistry.Register("SocialCustomerServiceSettings", open)
	internal.TypeRegistry.Register("SourceTrackingSettings", open)
	internal.TypeRegistry.Register("SurveySettings", open)
	internal.TypeRegistry.Register("SystemNotificationSettings", open)
	internal.TypeRegistry.Register("Territory2Settings", open)
	internal.TypeRegistry.Register("TrailheadSettings", open)
	internal.TypeRegistry.Register("TrialOrgSettings", open)
	internal.TypeRegistry.Register("UserEngagementSettings", open)
	internal.TypeRegistry.Register("UserInterfaceSettings", open)
	internal.TypeRegistry.Register("UserManagementSettings", open)
	internal.TypeRegistry.Register("VoiceSettings", open)
	internal.TypeRegistry.Register("WebToXSettings", open)
	internal.TypeRegistry.Register("WorkDotComSettings", open)
}

type GenericNode struct {
	XMLName xml.Name
	Attrs   []xml.Attr    `xml:",any,attr"`
	Content []GenericNode `xml:",any"`
	Text    string        `xml:",chardata"`
}

type Settings struct {
	metadata.MetadataInfo
	XMLName xml.Name
	XMLNS   string        `xml:"xmlns,attr"`
	Content []GenericNode `xml:",any"`
	Text    string        `xml:",chardata"`
}

func (c *Settings) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Settings) Type() metadata.MetadataType {
	return c.XMLName.Local
}

func Open(path string) (*Settings, error) {
	p := &Settings{}
	return p, metadata.ParseMetadataXml(p, path)
}
