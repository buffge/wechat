package base

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
)

type (
	// 签名方式
	SignType string
	// 付款方式
	TradeType string
	// 货币类型
	FeeType string
	// baseURL
	APIBaseURL string
	// 请求路径
	URLEndpoint string
	// 性别
	Sex string
)

const (
	SignTypeMD5    SignType = "MD5"
	SignTypeSHA256 SignType = "HMAC-SHA256"

	// JSAPI支付（或小程序支付）
	TradeTypeJSAPI TradeType = "JSAPI"
	// native 支付
	TradeTypeNative TradeType = "NATIVE"
	// app 支付
	TradeTypeAPP TradeType = "APP"
	// H5 支付
	TradeTypeMWEB TradeType = "MWEB"
	// 付款码支付
	TradeTypeMICROPAY TradeType = "MICROPAY"

	// 人民币
	FeeTypeCNY FeeType = "CNY"

	SexUndefined Sex = "0"
	SexMan       Sex = "1"
	SexWomen     Sex = "2"

	// 统一下单中 LimitPay字段如果设置此字段则用户不可以用信用卡支付
	NoCredit = "no_credit"
	// 统一下单中 Receipt字段 传入Y时，支付成功消息和支付详情页将出现开票入口
	ShowReceipt = "Y"

	// 商户平台api
	MchBaseURL APIBaseURL = "https://api.mch.weixin.qq.com/"
	// 微信基础api
	BaseURL APIBaseURL = "https://api.weixin.qq.com/"
	// 开放平台基础api
	OpenBaseURL APIBaseURL = "https://open.weixin.qq.com/"

	// 统一下单
	UnifiedOrderEndpoint URLEndpoint = "pay/unifiedorder"
	AccessTokenEndpoint  URLEndpoint = "cgi-bin/token"
	// oauth2 跳转登陆
	Oauth2Endpoint URLEndpoint = "connect/oauth2/authorize"
	// oauth2 获取ak
	Oauth2AccessTokenEndpoint URLEndpoint = "sns/oauth2/access_token"
	// oauth2 获取用户信息
	Oauth2GetUserInfoEndpoint URLEndpoint = "sns/userinfo"
	// 微信公众号获取用户信息
	GetUserInfoEndpoint URLEndpoint = "cgi-bin/user/info"
	// 二维码 扫码登陆
	QrCodeAuthEndpoint URLEndpoint = "connect/qrconnect"

	// 微信调试模式
	SandboxPrefix = "sandboxnew/"

	RetCodeSuccess = "SUCCESS"
	ResCodeSuccess = "SUCCESS"
	ResCodeFailed  = "FAIL"

	DefaultLang = "zh-CN"

	// 用户是否订阅了公众号
	HasSubscribe = 1
	NotSubscribe = 0
)

func GenerateSign(params map[string]string, key string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		if params[k] == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(params[k])
	}
	buf.WriteString("&key")
	buf.WriteByte('=')
	buf.WriteString(key)
	signType := params["signType"]
	var sign []byte
	switch signType {
	case string(SignTypeMD5):
		m5 := md5.New()
		_, _ = m5.Write(buf.Bytes())
		sign = m5.Sum(nil)
	default:
		m5 := md5.New()
		_, _ = m5.Write(buf.Bytes())
		sign = m5.Sum(nil)
	}
	return strings.ToUpper(hex.EncodeToString(sign))
}
