package payment

import (
	"log"
	"reflect"
	"testing"
)

func TestOrder_UnifiedOrder(t *testing.T) {
	type fields struct {
		App *App
	}
	type args struct {
		d *UnifiedOrderReqData
	}
	conf := &Conf{
		AppID:     "wx1c573801a12a61d1",
		MchID:     "1578535931",
		Key:       "xlkj8888888888888888888888888888",
		NotifyURL: "https://api.stream-capital.com/notify/wxPay",
	}
	app := NewPayment(conf, nil)
	d := &UnifiedOrderReqData{
		Body:       "buffge 测试商品购买",
		TotalFee:   12,
		OutTradeNo: "123asdasd",
		TradeType:  "APP",
		// OpenID:     "12312313asdasd",
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantBody []byte
		wantErr  bool
	}{
		{name: testing.CoverMode(), fields: fields{App: app}, args: args{d: d}, wantBody: []byte{},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := &Order{
				App: tt.fields.App,
			}
			resp, err := order.UnifiedOrder(tt.args.d)
			log.Printf("%+v", resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnifiedOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp != nil && !reflect.DeepEqual(resp.ReturnCode, tt.wantBody) {
				t.Errorf("UnifiedOrder() gotBody = %v, want %v", resp.ReturnCode, tt.wantBody)
			} else {
				t.Errorf("UnifiedOrder() resp is %v", resp)
			}
		})
	}
}
