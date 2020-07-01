package officialaccount

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/buffge/wechat/pkg/base"
	"github.com/buffge/wechat/pkg/utils"
)

type (
	Oauth struct {
		app *App
	}
	// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
	OauthUser struct {
		OpenID     *string   `json:"openid,omitempty"`
		Nickname   *string   `json:"nickname,omitempty"`
		Sex        *base.Sex `json:"sex,omitempty"`
		Province   *string   `json:"province,omitempty"`
		City       *string   `json:"city,omitempty"`
		Country    *string   `json:"country,omitempty"`
		HeadImgURL *string   `json:"headimgurl,omitempty"`
		Privilege  *[]string `json:"privilege,omitempty"`
		UnionID    *string   `json:"unionid,omitempty"`

		ErrCode *string `json:"errcode,omitempty"`
		ErrMsg  *string `json:"errmsg,omitempty"`
	}
	// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html#1
	OauthAccessToken struct {
		AccessToken  *string `json:"access_token,omitempty"`
		ExpiresIn    *int    `json:"expires_in,omitempty"`
		RefreshToken *string `json:"refresh_token,omitempty"`
		OpenID       *string `json:"openid,omitempty"`
		Scope        *string `json:"scope,omitempty"`
		ErrCode      *string `json:"errcode,omitempty"`
		ErrMsg       *string `json:"errmsg,omitempty"`
	}
	// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
	User struct {
		OauthUser
		Subscribe      int     `json:"subscribe"`
		SubscribeTime  *int64  `json:"subscribe_time,omitempty"`
		Remark         *int    `json:"remark,omitempty"`
		GroupID        *int    `json:"groupid,omitempty"`
		TagIDList      *[]int  `json:"tagid_list,omitempty"`
		SubscribeScene *string `json:"subscribe_scene,omitempty"`
		QrScene        *string `json:"qr_scene,omitempty"`
		QrSceneStr     *string `json:"qr_scene_str,omitempty"`
	}
)

func NewOauth(app *App) *Oauth {
	return &Oauth{
		app,
	}
}

func (oauth *Oauth) GetAccessToken(code string) (*OauthAccessToken, error) {
	q := &url.Values{
		"appid":      []string{oauth.app.AppID},
		"secret":     []string{oauth.app.AppSecret},
		"code":       []string{code},
		"grant_type": []string{"authorization_code"},
	}
	resp, err := utils.HTTPGet(oauth.app.BuildURL(base.Oauth2AccessTokenEndpoint), q)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	akData := &OauthAccessToken{}
	if err = json.Unmarshal(bytes, akData); err != nil {
		return nil, fmt.Errorf("resp is not valid, err: %v\n", err)
	}
	if akData.ErrMsg != nil {
		return nil, errors.New(*akData.ErrMsg)
	}
	if akData.AccessToken == nil || akData.ExpiresIn == nil {
		return nil, errors.New("ret val is not valid")
	}
	return akData, nil
}
func (oauth *Oauth) GetUserByToken(ak *OauthAccessToken) (*OauthUser, error) {
	q := &url.Values{
		"access_token": []string{*ak.AccessToken},
		"openid":       []string{*ak.OpenID},
		"lang":         []string{base.DefaultLang},
	}
	resp, err := utils.HTTPGet(oauth.app.BuildURL(base.Oauth2GetUserInfoEndpoint), q)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	user := &OauthUser{}
	if err = json.Unmarshal(bytes, user); err != nil {
		return nil, fmt.Errorf("resp is not valid, err: %v\n", err)
	}
	if user.ErrMsg != nil {
		return nil, errors.New(*user.ErrMsg)
	}
	if user.OpenID == nil {
		return nil, errors.New("ret val is not valid")
	}
	return user, nil
}
func (oauth *Oauth) GetUserByCode(code string) (*OauthUser, error) {
	ak, err := oauth.GetAccessToken(code)
	if err != nil {
		return nil, err
	}
	return oauth.GetUserByToken(ak)
}
func (oauth *Oauth) GetUserByOpenID(openID string) (*User, error) {
	ak, err := oauth.app.GetAccessToken()
	if err != nil {
		return nil, err
	}
	q := &url.Values{
		"access_token": []string{ak},
		"openid":       []string{openID},
		"lang":         []string{base.DefaultLang},
	}
	resp, err := utils.HTTPGet(oauth.app.BuildURL(base.GetUserInfoEndpoint), q)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	user := &User{}
	if err = json.Unmarshal(bytes, user); err != nil {
		return nil, fmt.Errorf("resp is not valid, err: %v\n", err)
	}
	if user.ErrMsg != nil {
		return nil, errors.New(*user.ErrMsg)
	}
	if user.OpenID == nil {
		return nil, errors.New("ret val is not valid")
	}
	return user, nil
}
func (oauth *Oauth) redirect(callbackURL string, state string, endpoint base.URLEndpoint) string {
	q := url.Values{
		"appid":         []string{oauth.app.AppID},
		"redirect_uri":  []string{callbackURL},
		"response_type": []string{"code"},
		"scope":         []string{"snsapi_userinfo"},
		"state":         []string{state},
	}
	return string(base.OpenBaseURL) + string(endpoint) + "?" + q.Encode() + "#wechat_redirect"
}
func (oauth *Oauth) OauthRedirect(callbackURL string, state string) string {
	return oauth.redirect(callbackURL, state, base.Oauth2Endpoint)
}
func (oauth *Oauth) QrCodeAuthRedirect(callbackURL string, state string) string {
	return oauth.redirect(callbackURL, state, base.QrCodeAuthEndpoint)
}
