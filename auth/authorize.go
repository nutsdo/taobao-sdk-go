package auth

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	WEB_HOST = "https://oauth.taobao.com/"
	H5_HOST  = "https://oauth.m.taobao.com/"
)

func (c *Credentials) AuthorizeUrlBuilder(redirect, view string) string {
	query := url.Values{}

	query.Set("client_id", c.AppKey)
	query.Set("response_type", "code")
	query.Set("redirect_uri", redirect)
	query.Set("state", "123123")
	query.Set("view", "web")

	var str strings.Builder
	if view == "web" {
		str.WriteString(WEB_HOST)

	} else if view == "wap" {
		str.WriteString(H5_HOST)
	}

	str.WriteString("authorize")
	str.WriteByte('?')
	str.WriteString(query.Encode())

	return str.String()
}

func (c *Credentials) Authorize(redirect, view string) {

	http.RedirectHandler(c.AuthorizeUrlBuilder(redirect, view), 302)
}

//get code
func (c *Credentials) AccessToken(code, redirect, view string) (data []byte, err error) {
	query := url.Values{}

	query.Set("client_id", c.AppKey)
	query.Set("client_secret", string(c.AppSecret))
	query.Set("grant_type", "authorization_code")
	query.Set("code", code)
	query.Set("state", "123123")
	query.Set("redirect_uri", redirect)
	query.Set("view", "web")

	var str strings.Builder
	if view == "web" {
		str.WriteString(WEB_HOST)

	} else if view == "wap" {
		str.WriteString(H5_HOST)
	}
	str.WriteString("token")

	str.WriteByte('?')
	str.WriteString(query.Encode())

	resp, err := http.Post(str.String(), "application/json", nil)

	if err != nil {
		return nil, err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes, nil
}

func (c *Credentials) RefreshToken(refreshToken, view string) (data []byte, err error) {
	query := url.Values{}

	query.Set("client_id", c.AppKey)
	query.Set("client_secret", string(c.AppSecret))
	query.Set("grant_type", "refresh_token")
	query.Set("refresh_token", refreshToken)
	query.Set("state", "123123")
	query.Set("view", "web")

	var str strings.Builder
	if view == "web" {
		str.WriteString(WEB_HOST)

	} else if view == "wap" {
		str.WriteString(H5_HOST)
	}
	str.WriteString("token")

	str.WriteByte('?')
	str.WriteString(query.Encode())

	resp, err := http.Post(str.String(), "application/json", nil)

	if err != nil {
		return nil, err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes, nil
}
