package network

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Network"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Network struct {
	metadata.MetadataInfo
	XMLName                xml.Name `xml:"Network"`
	Xmlns                  string   `xml:"xmlns,attr"`
	AllowInternalUserLogin struct {
		Text string `xml:",chardata"`
	} `xml:"allowInternalUserLogin"`
	AllowMembersToFlag struct {
		Text string `xml:",chardata"`
	} `xml:"allowMembersToFlag"`
	ChangePasswordTemplate struct {
		Text string `xml:",chardata"`
	} `xml:"changePasswordTemplate"`
	CommunityRoles struct {
		CustomerUserRole struct {
			Text string `xml:",chardata"`
		} `xml:"customerUserRole"`
		EmployeeUserRole struct {
			Text string `xml:",chardata"`
		} `xml:"employeeUserRole"`
		PartnerUserRole struct {
			Text string `xml:",chardata"`
		} `xml:"partnerUserRole"`
	} `xml:"communityRoles"`
	DisableReputationRecordConversations struct {
		Text string `xml:",chardata"`
	} `xml:"disableReputationRecordConversations"`
	EmailSenderAddress struct {
		Text string `xml:",chardata"`
	} `xml:"emailSenderAddress"`
	EmailSenderName struct {
		Text string `xml:",chardata"`
	} `xml:"emailSenderName"`
	EnableCustomVFErrorPageOverrides struct {
		Text string `xml:",chardata"`
	} `xml:"enableCustomVFErrorPageOverrides"`
	EnableDirectMessages struct {
		Text string `xml:",chardata"`
	} `xml:"enableDirectMessages"`
	EnableExperienceBundleBasedSnaOverrideEnabled struct {
		Text string `xml:",chardata"`
	} `xml:"enableExperienceBundleBasedSnaOverrideEnabled"`
	EnableGuestChatter struct {
		Text string `xml:",chardata"`
	} `xml:"enableGuestChatter"`
	EnableGuestFileAccess struct {
		Text string `xml:",chardata"`
	} `xml:"enableGuestFileAccess"`
	EnableGuestMemberVisibility struct {
		Text string `xml:",chardata"`
	} `xml:"enableGuestMemberVisibility"`
	EnableInvitation struct {
		Text string `xml:",chardata"`
	} `xml:"enableInvitation"`
	EnableKnowledgeable struct {
		Text string `xml:",chardata"`
	} `xml:"enableKnowledgeable"`
	EnableMemberVisibility struct {
		Text string `xml:",chardata"`
	} `xml:"enableMemberVisibility"`
	EnableNicknameDisplay struct {
		Text string `xml:",chardata"`
	} `xml:"enableNicknameDisplay"`
	EnablePrivateMessages struct {
		Text string `xml:",chardata"`
	} `xml:"enablePrivateMessages"`
	EnableReputation struct {
		Text string `xml:",chardata"`
	} `xml:"enableReputation"`
	EnableShowAllNetworkSettings struct {
		Text string `xml:",chardata"`
	} `xml:"enableShowAllNetworkSettings"`
	EnableSiteAsContainer struct {
		Text string `xml:",chardata"`
	} `xml:"enableSiteAsContainer"`
	EnableTalkingAboutStats struct {
		Text string `xml:",chardata"`
	} `xml:"enableTalkingAboutStats"`
	EnableTopicAssignmentRules struct {
		Text string `xml:",chardata"`
	} `xml:"enableTopicAssignmentRules"`
	EnableTopicSuggestions struct {
		Text string `xml:",chardata"`
	} `xml:"enableTopicSuggestions"`
	EnableUpDownVote struct {
		Text string `xml:",chardata"`
	} `xml:"enableUpDownVote"`
	ForgotPasswordTemplate struct {
		Text string `xml:",chardata"`
	} `xml:"forgotPasswordTemplate"`
	GatherCustomerSentimentData struct {
		Text string `xml:",chardata"`
	} `xml:"gatherCustomerSentimentData"`
	NetworkMemberGroups struct {
		Profile []struct {
			Text string `xml:",chardata"`
		} `xml:"profile"`
	} `xml:"networkMemberGroups"`
	NetworkPageOverrides struct {
		ChangePasswordPageOverrideSetting struct {
			Text string `xml:",chardata"`
		} `xml:"changePasswordPageOverrideSetting"`
		ForgotPasswordPageOverrideSetting struct {
			Text string `xml:",chardata"`
		} `xml:"forgotPasswordPageOverrideSetting"`
		HomePageOverrideSetting struct {
			Text string `xml:",chardata"`
		} `xml:"homePageOverrideSetting"`
		LoginPageOverrideSetting struct {
			Text string `xml:",chardata"`
		} `xml:"loginPageOverrideSetting"`
		SelfRegProfilePageOverrideSetting struct {
			Text string `xml:",chardata"`
		} `xml:"selfRegProfilePageOverrideSetting"`
	} `xml:"networkPageOverrides"`
	PicassoSite struct {
		Text string `xml:",chardata"`
	} `xml:"picassoSite"`
	SelfRegistration struct {
		Text string `xml:",chardata"`
	} `xml:"selfRegistration"`
	SendWelcomeEmail struct {
		Text string `xml:",chardata"`
	} `xml:"sendWelcomeEmail"`
	Site struct {
		Text string `xml:",chardata"`
	} `xml:"site"`
	Status struct {
		Text string `xml:",chardata"`
	} `xml:"status"`
	Tabs struct {
		CustomTab []struct {
			Text string `xml:",chardata"`
		} `xml:"customTab"`
		DefaultTab struct {
			Text string `xml:",chardata"`
		} `xml:"defaultTab"`
		StandardTab []struct {
			Text string `xml:",chardata"`
		} `xml:"standardTab"`
	} `xml:"tabs"`
	UrlPathPrefix struct {
		Text string `xml:",chardata"`
	} `xml:"urlPathPrefix"`
	WelcomeTemplate struct {
		Text string `xml:",chardata"`
	} `xml:"welcomeTemplate"`
}

func (c *Network) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Network) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Network, error) {
	p := &Network{}
	return p, metadata.ParseMetadataXml(p, path)
}
