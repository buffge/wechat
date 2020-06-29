package payment

import (
	"net/http"
)

type (
	App struct {
		Conf
		Request *http.Request
	}
	Conf struct {
		AppID     string `json:"appID"`
		MchID     string `json:"mchID"`
		Key       string `json:"key"`
		CertPath  string `json:"certPath"`
		KeyPath   string `json:"keyPath"`
		NotifyURL string `json:"notifyURL"`
	}
)

func (app *App) Req(p map[string]string) *http.Response {
	return nil
}
