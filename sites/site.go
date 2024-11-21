package site

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "CustomSite"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type CustomSite struct {
	internal.MetadataInfo
	XMLName xml.Name `xml:"CustomSite"`
	Xmlns   string   `xml:"xmlns,attr"`
	Active  struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
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
	SiteType struct {
		Text string `xml:",chardata"`
	} `xml:"siteType"`
	Subdomain struct {
		Text string `xml:",chardata"`
	} `xml:"subdomain"`
}

func (c *CustomSite) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomSite) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*CustomSite, error) {
	p := &CustomSite{}
	return p, internal.ParseMetadataXml(p, path)
}
