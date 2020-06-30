package payment

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/buffge/wechat/pkg/base"
	"github.com/buffge/wechat/pkg/utils"
)

type (
	JSSdk struct {
		App *App
	}
	JSPayConfig struct {
		Timestamp int64         `json:"timestamp"`
		NonceStr  string        `json:"nonceStr"`
		WxPackage string        `json:"wxPackage"`
		SignType  base.SignType `json:"signType"`
		PaySign   string        `json:"paySign"`
	}
)

func (conf *JSPayConfig) GetWxParam(appID string) map[string]string {
	return map[string]string{
		"appId":     appID,
		"timeStamp": strconv.FormatInt(conf.Timestamp, 10),
		"nonceStr":  conf.NonceStr,
		"package":   conf.WxPackage,
		"signType":  string(conf.SignType),
	}
}
func (js *JSSdk) PayConfig(prepayID string) *JSPayConfig {
	conf := &JSPayConfig{
		Timestamp: time.Now().Unix(),
		NonceStr:  hex.EncodeToString(utils.RandomBytes(16)),
		WxPackage: "prepay_id=" + prepayID,
		SignType:  base.SignTypeMD5,
	}
	conf.PaySign = base.GenerateSign(conf.GetWxParam(js.App.AppID), js.App.Key)
	return conf
}
