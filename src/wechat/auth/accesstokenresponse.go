package auth

type Accesstokenresponse struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
}
