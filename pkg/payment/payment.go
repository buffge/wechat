package payment

import (
	"net/http"
)

type (
	App struct {
		Conf
		Request *http.Request
		jssdk   *JSSdk
		order   *Order
	}
	Conf struct {
		AppID     string `json:"appID"`
		MchID     string `json:"mchID"`
		Key       string `json:"key"`
		CertPath  string `json:"certPath"`
		KeyPath   string `json:"keyPath"`
		NotifyURL string `json:"notifyURL"`
		ISSandbox bool   `json:"isSandbox"`
	}
)

func NewPayment(conf *Conf, req *http.Request) *App {
	return &App{
		Conf:    *conf,
		Request: req,
	}
}
func (app *App) GetJsSDK() *JSSdk {
	if app.jssdk == nil {
		app.jssdk = NewJSSdk(app)
	}
	return app.jssdk
}
func (app *App) GetOauth() *Order {
	if app.order == nil {
		app.order = NewOrder(app)
	}
	return app.order
}
