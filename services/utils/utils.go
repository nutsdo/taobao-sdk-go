package utils

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/client"
	"github.com/nutsdo/taobao-sdk-go/config"
	"github.com/nutsdo/taobao-sdk-go/services"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

type OAuthClient struct {
	*tbk.TbkClient
}

type AccessToken struct {
	W2ExpiresIn    int64  `json:"w2_expires_in"`
	TaobaoUserId   string `json:"taobao_user_id"`
	TaobaoUserNick string `json:"taobao_user_nick"`
	W1ExpiresIn    int64  `json:"w1_expires_in"`
	ReExpiresIn    int64  `json:"re_expires_in"`
	R2ExpiresIn    int64  `json:"r2_expires_in"`
	ExpiresIn      int64  `json:"expires_in"`
	TokenType      string `json:"token_type"`
	RefreshToken   string `json:"refresh_token"`
	AccessToken    string `json:"access_token"`
	R1ExpiresIn    string `json:"r1_expires_in"`
}

func New(credentials *auth.Credentials, session string) *OAuthClient {

	return &OAuthClient{&tbk.TbkClient{credentials,
		client.DefaultClient, config.DefaultConfig, session}}
}

type AccessTokenRequest struct {
	*services.BaseRequest
}

type AccessTokenResponse struct {
	TopAuthTokenCreateResponse TopAuthTokenCreateResponse `json:"top_auth_token_create_response"`
}

type TopAuthTokenCreateResponse struct {
	TokenResult string `json:"token_result"`
}

func (r *AccessTokenRequest) Method() string {
	return "taobao.top.auth.token.create"
}

func (r *AccessTokenRequest) SetCode(code string) {
	r.SetValue("code", code)
}

func (r *AccessTokenRequest) SetUuid(uuid string) {
	r.SetValue("uuid", uuid)
}

func (c *OAuthClient) AccessToken(request services.TaoBaoRequest) (response *AccessTokenResponse, errresp *services.BaseErrorResponse) {

	response = NewAccessTokenResponse()
	resp, err := c.DoRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	response = NewAccessTokenResponse()

	err = services.Byte2Response(data,response)

	return response, nil
}

func NewAccessTokenResponse() *AccessTokenResponse {
	return &AccessTokenResponse{TopAuthTokenCreateResponse{}}
}
