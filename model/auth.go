package model

type Signin struct {
	Username string `json:"username" form:"username" query:"username" validate:"required,email"`
	Password string `json:"password" form:"password" query:"password" validate:"required,min=6,max=20"`
}

type OauthErr string

const (
	OauthErrInvalidRequest     OauthErr = "invalid_request"
	OauthErrInvalidClient      OauthErr = "invalid_client"
	OauthErrInvalidGrant       OauthErr = "invalid_grant"
	OauthErrInvalidScope       OauthErr = "invalid_scope"
	OauthErrUnauthorizedClient OauthErr = "unauthorized_client"
	OauthErrUnsupportedGrant   OauthErr = "unsupported_grant_type"
	OauthErrServer             OauthErr = "server_error"
	OauthErrTemporarily        OauthErr = "temporarily_unavailable"
)

type TokenResponse struct {
	AccessToken      string   `json:"access_token"`
	TokenType        string   `json:"token_type"`
	ExpiresIn        int      `json:"expires_in"`
	RefreshToken     string   `json:"refresh_token,omitempty"`
	Scope            string   `json:"scope,omitempty"`
	Error            OauthErr `json:"error,omitempty"`
	ErrorDescription string   `json:"error_description,omitempty"`
	ErrorURI         string   `json:"error_uri,omitempty"`
}
