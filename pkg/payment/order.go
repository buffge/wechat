package payment

import (
	"github.com/buffge/wechat"
	"github.com/buffge/wechat/pkg/base"
	"github.com/buffge/wechat/pkg/utils"
)

// wx doc https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
type (
	Order struct {
		App *wechat.PaymentApp
	}
	UnifiedOrderReqData struct {
		AppID      string         `json:"appID" xml:"app_id" validate:"required,max=32"`
		MchID      string         `json:"mchID" xml:"mch_id" validate:"required,max=32"`
		NonceStr   string         `json:"nonceStr" xml:"nonce_str" validate:"required,max=32"`
		Sign       string         `json:"sign" xml:"sign" validate:"required,max=32"`
		Body       string         `json:"body" xml:"body" validate:"required,max=128"`
		TotalFee   uint           `json:"totalFee" xml:"total_fee" validate:"required,min=1"`
		OutTradeNo string         `json:"outTradeNo" xml:"out_trade_no" validate:"required,max=32"`
		TradeType  base.TradeType `json:"tradeType" xml:"trade_type" validate:"required,oneof=JSAPI,NATIVE,APP,MWEB"`
		NotifyURL  string         `json:"notifyURL" xml:"notify_url" validate:"required,max=256"`
		// no required
		DeviceInfo     string       `json:"deviceInfo" xml:"device_info" validate:"omitempty,max=32"`
		SignType       string       `json:"signType" xml:"sign_type" validate:"omitempty,max=32"`
		Detail         string       `json:"detail" xml:"detail" validate:"omitempty,max=6000"`
		Attach         string       `json:"attach" xml:"attach" validate:"omitempty,max=127"`
		FeeType        base.FeeType `json:"feeType" xml:"fee_type" validate:"omitempty,oneof=CNY"`
		SpbillCreateIP string       `json:"spbillCreateIP" xml:"spbill_create_ip" validate:"omitempty,ip"`
		TimeStart      string       `json:"timeStart" xml:"time_start" validate:"omitempty,len=14"`
		TimeExpire     string       `json:"timeExpire" xml:"time_expire" validate:"omitempty,len=14"`
		GoodsTag       string       `json:"goodsTag" xml:"goods_tag" validate:"omitempty,max=32"`
		ProductID      string       `json:"productID" xml:"product_id" validate:"omitempty,max=32"`
		LimitPay       string       `json:"limitPay" xml:"limit_pay" validate:"omitempty,max=32"`
		OpenID         string       `json:"openid" xml:"openid" validate:"omitempty,max=128"`
		Receipt        string       `json:"receipt" xml:"receipt" validate:"omitempty,max=8"`
		SceneInfo      string       `json:"sceneInfo" xml:"scene_info" validate:"omitempty,max=256"`
	}
)

func (order *Order) UnifiedOrder(d *UnifiedOrderReqData) {

	app := order.App
	if d.SpbillCreateIP == "" && app.Request != nil {
		d.SpbillCreateIP = utils.GetClientIP(order.App.Request)
	}
	if d.NotifyURL == "" {
		d.NotifyURL = app.NotifyURL
	}

}
func (order *Order) Request(p map[string]string) {

}
