package response

type TokenResponse struct {
	WorkResponse
	AccessToken           string  `json:"access_token,omitempty"`
	AuthorizerAccessToken string  `json:"authorizer_access_token,omitempty"`
	ExpiresIn             float64 `json:"expires_in"`
}
