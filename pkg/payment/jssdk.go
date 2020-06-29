package payment

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/buffge/wechat"
	"github.com/buffge/wechat/pkg/base"
	"github.com/buffge/wechat/pkg/utils"
)

type (
	JSSdk struct {
		App *wechat.PaymentApp
	}
	JSPayConfig struct {
		Timestamp int64         `json:"timestamp"`
		NonceStr  string        `json:"nonceStr"`
		WxPackage string        `json:"wxPackage"`
		SignType  base.SignType `json:"signType"`
		PaySign   string        `json:"paySign"`
	}
)

func (js *JSSdk) PayConfig(prepayID string) *JSPayConfig {
	conf := &JSPayConfig{
		Timestamp: time.Now().Unix(),
		NonceStr:  hex.EncodeToString(utils.RandomBytes(16)),
		WxPackage: "prepay_id=" + prepayID,
		SignType:  base.SignTypeMD5,
	}
	conf.PaySign = js.GenerateSign(conf)
	return conf
}
func (js *JSSdk) GenerateSign(conf *JSPayConfig) string {
	params := map[string]string{
		"appId":     js.App.AppID,
		"timeStamp": strconv.FormatInt(conf.Timestamp, 10),
		"nonceStr":  conf.NonceStr,
		"package":   conf.WxPackage,
		"signType":  string(conf.SignType),
	}
	return base.GenerateSign(params, js.App.Key)
}
