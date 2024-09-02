package model

type IdentityProvider struct {
	ID               int32
	Name             string
	Type             IdpType
	IdentifierFilter string
	Config           string
}

type IdpType string

const (
	IdpUnspecified IdpType = "UNSPECIFIED"
	IdpOAuth2      IdpType = "OAUTH2"
)
