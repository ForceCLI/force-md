package connectedapps

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ConnectedApp"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ConnectedApp struct {
	metadata.MetadataInfo
	XMLName      xml.Name     `xml:"ConnectedApp"`
	Xmlns        string       `xml:"xmlns,attr"`
	ContactEmail TextLiteral  `xml:"contactEmail"`
	ContactPhone *TextLiteral `xml:"contactPhone"`
	Description  *TextLiteral `xml:"description"`
	Label        TextLiteral  `xml:"label"`
	OauthConfig  *struct {
		CallbackUrl                      *TextLiteral  `xml:"callbackUrl"`
		Certificate                      *TextLiteral  `xml:"certificate"`
		ConsumerKey                      *TextLiteral  `xml:"consumerKey"`
		IsAdminApproved                  *TextLiteral  `xml:"isAdminApproved"`
		IsClientCredentialEnabled        *TextLiteral  `xml:"isClientCredentialEnabled"`
		IsCodeCredentialEnabled          *TextLiteral  `xml:"isCodeCredentialEnabled"`
		IsCodeCredentialPostOnly         *TextLiteral  `xml:"isCodeCredentialPostOnly"`
		IsConsumerSecretOptional         *TextLiteral  `xml:"isConsumerSecretOptional"`
		IsIntrospectAllTokens            *TextLiteral  `xml:"isIntrospectAllTokens"`
		IsNamedUserJwtEnabled            *TextLiteral  `xml:"isNamedUserJwtEnabled"`
		IsPkceRequired                   *TextLiteral  `xml:"isPkceRequired"`
		IsRefreshTokenRotationEnabled    *TextLiteral  `xml:"isRefreshTokenRotationEnabled"`
		IsSecretRequiredForRefreshToken  *TextLiteral  `xml:"isSecretRequiredForRefreshToken"`
		IsSecretRequiredForTokenExchange *TextLiteral  `xml:"isSecretRequiredForTokenExchange"`
		IsTokenExchangeEnabled           *TextLiteral  `xml:"isTokenExchangeEnabled"`
		Scopes                           []TextLiteral `xml:"scopes"`
	} `xml:"oauthConfig"`
	OauthPolicy *struct {
		IpRelaxation               *TextLiteral `xml:"ipRelaxation"`
		IsTokenExchangeFlowEnabled *TextLiteral `xml:"isTokenExchangeFlowEnabled"`
		RefreshTokenPolicy         *TextLiteral `xml:"refreshTokenPolicy"`
	} `xml:"oauthPolicy"`
	PermissionSetName []TextLiteral `xml:"permissionSetName"`
	ProfileName       []TextLiteral `xml:"profileName"`
}

func (c *ConnectedApp) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*ConnectedApp, error) {
	p := &ConnectedApp{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (c *ConnectedApp) Type() metadata.MetadataType {
	return NAME
}
