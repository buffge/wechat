package officialaccount

import (
	"net/url"

	"github.com/buffge/wechat/pkg/base"
	"github.com/buffge/wechat/pkg/utils"
)

type (
	App struct {
		Conf
	}
	Conf struct {
		AppID          string `json:"appID"`
		AppSecret      string `json:"appSecret"`
		Token          string `json:"token"`
		EncodingAESKey string `json:"encodingAESKey"`
	}
)

func NewApp(conf *Conf) *App {
	return &App{
		Conf: *conf,
	}
}
func (app *App) GetAccessToken() {
	query := &url.Values{
		"grant_type": []string{"client_credential"},
		"appid":      []string{app.AppID},
		"secret":     []string{app.AppSecret},
	}
	utils.HTTPGet(app.BuildURL(base.AccessTokenEndpoint), query)
}
func (app *App) BuildURL(endpoint base.URLEndpoint) string {
	fmtEndpoint := string(endpoint)
	return string(base.BaseURL) + fmtEndpoint
}
