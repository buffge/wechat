package officialaccount

import (
	"github.com/buffge/wechat/pkg/base"
	"testing"
)

func TestApp_GetAccessToken(t *testing.T) {
	// os.Setenv("HTTP_PROXY", "")
	// os.Setenv("HTTPS_PROXY", "")
	// os.Setenv("NO_PROXY", "true")
	type fields struct {
		Conf  Conf
		Cache base.Cache
	}
	conf := &Conf{
		AppID:     "wx9b02442e98762de9",
		AppSecret: "e261ca5f3f7a84094c4aa5e10cea490d",
	}
	tests := []struct {
		name            string
		fields          fields
		wantAccessToken string
		wantErr         bool
	}{
		{name: testing.CoverMode(), fields: fields{
			Conf: *conf,
		}, wantAccessToken: "", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := NewApp(conf)
			gotAccessToken, err := app.GetAccessToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAccessToken != tt.wantAccessToken {
				t.Errorf("GetAccessToken() gotAccessToken = %v, want %v", gotAccessToken, tt.wantAccessToken)
			}
		})
	}
}
