package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

func NewGetAccessToken(appID string, appSecret string) *GetAccessTokenRequest {
	return &GetAccessTokenRequest{
		AppID:     appID,
		AppSecret: appSecret,
	}
}

type GetAccessTokenRequest struct {
	AppID     string
	AppSecret string
}

func (r *GetAccessTokenRequest) GetBody() io.Reader {
	return bytes.NewReader([]byte{})
}
func (r *GetAccessTokenRequest) GetURL() string {
	return fmt.Sprintf(
		"product/getAccessToken?api_id=%s&app_secret=%s",
		r.AppID,
		r.AppSecret,
	)
}
func (r *GetAccessTokenRequest) GetType() string {
	return "GET"
}
func (r *GetAccessTokenRequest) SetAccessToken(token string) {
}

type GetAccessTokenReponse struct {
	AccessToken string `json:"access_token"`
	Expiry      int    `json:"expires_in"`
	Code        int    `json:"code"`
}

func (r *GetAccessTokenReponse) ParseResponse(rsp io.Reader) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rsp)
	json.Unmarshal(buf.Bytes(), &r)
}
