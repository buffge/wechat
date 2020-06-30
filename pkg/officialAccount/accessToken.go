package officialaccount

type (
	AccessTokenRespData struct {
		AccessToken *string `json:"access_token,omitempty"`
		ExpiresIn   *int    `json:"expires_in,omitempty"`
		ErrCode     *int    `json:"errcode,omitempty"`
		ErrMsg      *string `json:"errmsg,omitempty"`
	}
)
