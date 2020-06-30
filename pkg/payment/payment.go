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
		ISSandbox bool   `json:"isSandbox"`
	}
)

func NewPayment(conf *Conf, req *http.Request) *App {
	return &App{
		Conf:    *conf,
		Request: req,
	}
}
