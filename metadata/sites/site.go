package site

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "CustomSite"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type CustomSite struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"CustomSite"`
	Xmlns   string   `xml:"xmlns,attr"`
	Active  struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
	AllowGuestPaymentsApi struct {
		Text string `xml:",chardata"`
	} `xml:"allowGuestPaymentsApi"`
	AllowHomePage struct {
		Text string `xml:",chardata"`
	} `xml:"allowHomePage"`
	AllowStandardAnswersPages struct {
		Text string `xml:",chardata"`
	} `xml:"allowStandardAnswersPages"`
	AllowStandardIdeasPages struct {
		Text string `xml:",chardata"`
	} `xml:"allowStandardIdeasPages"`
	AllowStandardLookups struct {
		Text string `xml:",chardata"`
	} `xml:"allowStandardLookups"`
	AllowStandardPortalPages struct {
		Text string `xml:",chardata"`
	} `xml:"allowStandardPortalPages"`
	AllowStandardSearch struct {
		Text string `xml:",chardata"`
	} `xml:"allowStandardSearch"`
	AuthorizationRequiredPage struct {
		Text string `xml:",chardata"`
	} `xml:"authorizationRequiredPage"`
	BandwidthExceededPage struct {
		Text string `xml:",chardata"`
	} `xml:"bandwidthExceededPage"`
	BrowserXssProtection struct {
		Text string `xml:",chardata"`
	} `xml:"browserXssProtection"`
	CachePublicVisualforcePagesInProxyServers struct {
		Text string `xml:",chardata"`
	} `xml:"cachePublicVisualforcePagesInProxyServers"`
	ClickjackProtectionLevel struct {
		Text string `xml:",chardata"`
	} `xml:"clickjackProtectionLevel"`
	ContentSniffingProtection struct {
		Text string `xml:",chardata"`
	} `xml:"contentSniffingProtection"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	EnableAuraRequests struct {
		Text string `xml:",chardata"`
	} `xml:"enableAuraRequests"`
	FavoriteIcon struct {
		Text string `xml:",chardata"`
	} `xml:"favoriteIcon"`
	FileNotFoundPage struct {
		Text string `xml:",chardata"`
	} `xml:"fileNotFoundPage"`
	GenericErrorPage struct {
		Text string `xml:",chardata"`
	} `xml:"genericErrorPage"`
	InMaintenancePage struct {
		Text string `xml:",chardata"`
	} `xml:"inMaintenancePage"`
	InactiveIndexPage struct {
		Text string `xml:",chardata"`
	} `xml:"inactiveIndexPage"`
	IndexPage struct {
		Text string `xml:",chardata"`
	} `xml:"indexPage"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	RedirectToCustomDomain struct {
		Text string `xml:",chardata"`
	} `xml:"redirectToCustomDomain"`
	ReferrerPolicyOriginWhenCrossOrigin struct {
		Text string `xml:",chardata"`
	} `xml:"referrerPolicyOriginWhenCrossOrigin"`
	RobotsTxtPage struct {
		Text string `xml:",chardata"`
	} `xml:"robotsTxtPage"`
	SelfRegPage struct {
		Text string `xml:",chardata"`
	} `xml:"selfRegPage"`
	ServerIsDown struct {
		Text string `xml:",chardata"`
	} `xml:"serverIsDown"`
	SiteAdmin struct {
		Text string `xml:",chardata"`
	} `xml:"siteAdmin"`
	SiteGuestRecordDefaultOwner struct {
		Text string `xml:",chardata"`
	} `xml:"siteGuestRecordDefaultOwner"`
	SiteRedirectMappings []struct {
		Action struct {
			Text string `xml:",chardata"`
		} `xml:"action"`
		IsActive struct {
			Text string `xml:",chardata"`
		} `xml:"isActive"`
		Source struct {
			Text string `xml:",chardata"`
		} `xml:"source"`
		Target struct {
			Text string `xml:",chardata"`
		} `xml:"target"`
	} `xml:"siteRedirectMappings"`
	SiteTemplate struct {
		Text string `xml:",chardata"`
	} `xml:"siteTemplate"`
	SiteType struct {
		Text string `xml:",chardata"`
	} `xml:"siteType"`
	UrlPathPrefix struct {
		Text string `xml:",chardata"`
	} `xml:"urlPathPrefix"`
	Subdomain struct {
		Text string `xml:",chardata"`
	} `xml:"subdomain"`
}

func (c *CustomSite) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomSite) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CustomSite, error) {
	p := &CustomSite{}
	return p, metadata.ParseMetadataXml(p, path)
}
