package endpoints

import "io"

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

}
func (r *GetAccessTokenRequest) GetURL() string {

}
func (r *GetAccessTokenRequest) GetType() string {

}
func (r *GetAccessTokenRequest) SetAccessToken(token string) {

}
