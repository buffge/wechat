package officialaccount

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/buffge/wechat/pkg/base/cache"

	"github.com/buffge/wechat/pkg/base"
	"github.com/buffge/wechat/pkg/utils"
)

type (
	App struct {
		Conf
		Cache base.Cache
		oauth *Oauth
	}
	Conf struct {
		AppID          string `json:"appID"`
		AppSecret      string `json:"appSecret"`
		Token          string `json:"token"`
		EncodingAESKey string `json:"encodingAESKey"`
		Cache          base.Cache
	}
)

func NewApp(conf *Conf) *App {
	app := &App{
		Conf:  *conf,
		Cache: conf.Cache,
	}
	if conf.Cache == nil {
		app.Cache = cache.GetDefaultCache()
	}
	return app
}
func (app *App) GetAccessToken() (accessToken string, err error) {
	accessToken, err = app.getAccessTokenFormCache()
	if err != nil || accessToken == "" {
		return app.getAccessTokenFormServer()
	}
	return
}
func (app *App) getAccessTokenFormCache() (accessToken string, err error) {
	var data interface{}
	if data, err = app.Cache.Get(base.AccessTokenKey); err != nil {
		return "", err
	}
	var ok bool
	if accessToken, ok = data.(string); !ok || accessToken == "" {
		return "", errors.New("not found")
	}
	return
}
func (app *App) getAccessTokenFormServer() (accessToken string, err error) {
	query := &url.Values{
		"grant_type": []string{"client_credential"},
		"appid":      []string{app.AppID},
		"secret":     []string{app.AppSecret},
	}
	resp, err := utils.HTTPGet(app.BuildURL(base.AccessTokenEndpoint), query)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	akData := &AccessTokenRespData{}
	if err = json.Unmarshal(bytes, akData); err != nil {
		return "", fmt.Errorf("resp is not valid, err: %v\n", err)
	}
	if akData.ErrMsg != nil {
		return "", errors.New(*akData.ErrMsg)
	}
	if akData.AccessToken == nil || akData.ExpiresIn == nil {
		return "", errors.New("ret val is not valid")
	}
	accessToken = *akData.AccessToken
	app.Cache.Set(base.AccessTokenKey, accessToken, time.Duration(*akData.ExpiresIn)*time.Second)
	return
}
func (app *App) BuildURL(endpoint base.URLEndpoint) string {
	fmtEndpoint := string(endpoint)
	return string(base.BaseURL) + fmtEndpoint
}
func (app *App) GetOauth() *Oauth {
	if app.oauth == nil {
		app.oauth = NewOauth(app)
	}
	return app.oauth
}
