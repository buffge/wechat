package payment

import (
	"encoding/xml"
	"strconv"

	"github.com/buffge/wechat/pkg/base"
	"github.com/buffge/wechat/pkg/utils"
)

// wx doc https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
type (
	Order struct {
		app *App
	}
	UnifiedOrderReqData struct {
		AppID      string         `json:"appID" xml:"appid" validate:"required,max=32"`
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

		XMLName struct{} `xml:"xml"`
	}
	UnifiedOrderRespData struct {
		ReturnCode string         `xml:"return_code" validate:"required,max=16"`
		ReturnMsg  string         `xml:"return_msg" validate:"required,max=128"`
		AppID      string         `xml:"appid,omitempty" validate:"omitempty,max=32"`
		MchID      string         `xml:"mch_id,omitempty" validate:"omitempty,max=32"`
		DeviceInfo *string        `xml:"device_info,omitempty" validate:"omitempty,max=32"`
		NonceStr   string         `xml:"nonce_str,omitempty" validate:"omitempty,max=32"`
		Sign       string         `xml:"sign,omitempty" validate:"omitempty,max=32"`
		ResultCode string         `xml:"result_code,omitempty" validate:"omitempty,max=16"`
		ErrCode    *string        `xml:"err_code,omitempty" validate:"omitempty,max=32"`
		ErrCodeDes *string        `xml:"err_code_des,omitempty" validate:"omitempty,max=128"`
		TradeType  base.TradeType `xml:"trade_type,omitempty" validate:"omitempty,oneof=JSAPI,NATIVE,APP,MWEB"`
		PrepayID   string         `xml:"prepay_id,omitempty" validate:"omitempty,max=64"`
		// trade_type=NATIVE时有返回，此url用于生成支付二维码，然后提供给用户进行扫码支付
		CodeURL *string `xml:"code_url,omitempty" validate:"omitempty,max=64"`

		XMLName struct{} `xml:"xml"`
	}
)

func NewOrder(app *App) *Order {
	return &Order{
		app,
	}
}
func (d *UnifiedOrderReqData) GetWXParam() map[string]string {
	return map[string]string{
		"appid":            d.AppID,
		"mch_id":           d.MchID,
		"device_info":      d.DeviceInfo,
		"nonce_str":        d.NonceStr,
		"sign_type":        d.SignType,
		"body":             d.Body,
		"detail":           d.Detail,
		"attach":           d.Attach,
		"out_trade_no":     d.OutTradeNo,
		"fee_type":         string(d.FeeType),
		"total_fee":        strconv.Itoa(int(d.TotalFee)),
		"spbill_create_ip": d.SpbillCreateIP,
		"time_start":       d.TimeStart,
		"time_expire":      d.TimeExpire,
		"goods_tag":        d.GoodsTag,
		"notify_url":       d.NotifyURL,
		"trade_type":       string(d.TradeType),
		"product_id":       d.ProductID,
		"limit_pay":        d.LimitPay,
		"openid":           d.OpenID,
		"receipt":          d.Receipt,
		"scene_info":       d.SceneInfo,
	}
}

func (order *Order) UnifiedOrder(d *UnifiedOrderReqData) (*UnifiedOrderRespData, error) {
	app := order.app
	d.AppID = order.app.AppID
	d.MchID = order.app.MchID
	if d.SpbillCreateIP == "" && app.Request != nil {
		d.SpbillCreateIP = utils.GetClientIP(order.app.Request)
	}
	if d.NotifyURL == "" {
		d.NotifyURL = app.NotifyURL
	}
	if d.NonceStr == "" {
		d.NonceStr = utils.RandomStr(32)
	}
	d.Sign = base.GenerateSign(d.GetWXParam(), app.Key)
	body, err := utils.PostXML(order.BuildURL(base.UnifiedOrderEndpoint), d)
	if err != nil {
		return nil, err
	}
	resp := &UnifiedOrderRespData{}
	if err := xml.Unmarshal(body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (order *Order) BuildURL(endpoint base.URLEndpoint) string {
	fmtEndpoint := string(endpoint)
	if order.app.ISSandbox {
		fmtEndpoint = base.SandboxPrefix + fmtEndpoint
	}
	return string(base.MchBaseURL) + fmtEndpoint
}
