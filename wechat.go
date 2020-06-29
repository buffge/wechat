package wechat

import (
	"net/http"

	"github.com/buffge/wechat/pkg/payment"
)

type (
	OfficialAccountApp struct {
		OfficialAccountConf
	}
	OfficialAccountConf struct {
		AppID          string `json:"appID"`
		AppSecret      string `json:"appSecret"`
		Token          string `json:"token"`
		EncodingAESKey string `json:"encodingAESKey"`
	}
)

func NewOfficialAccount(conf *OfficialAccountConf) *OfficialAccountApp {
	return &OfficialAccountApp{
		OfficialAccountConf: *conf,
	}
}
func NewPayment(conf *payment.Conf, req *http.Request) *payment.App {
	return &payment.App{
		Conf:    *conf,
		Request: req,
	}
}
