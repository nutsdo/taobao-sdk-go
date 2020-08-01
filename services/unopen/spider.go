package unopen

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TaoPwd struct {
	Api  string    `json:"api"`
	Data PwdDetail `json:"data"`
	Ret  []string  `json:"ret"`
	V    string    `json:"v"`
}

type PwdDetail struct {
	CreateAppkey    string `json:"createAppkey"`
	ValidDate       string `json:"validDate"`
	TaopwdOwnerId   string `json:"taopwdOwnerId"`
	Title           string `json:"title"`
	TemplateId      string `json:"templateId"`
	MyTaopwdToast   string `json:"myTaopwdToast"`
	Content         string `json:"content"`
	Url             string `json:"url"`
	PicUrl          string `json:"picUrl"`
	RankPic         string `json:"rankPic"`
	Password        string `json:"password"`
	OwnerName       string `json:"ownerName"`
	Price           string `json:"price"`
	RankNum         string `json:"rankNum"`
	BizId           string `json:"bizId"`
	LeftButtonText  string `json:"leftButtonText"`
	RightButtonText string `json:"rightButtonText"`
}

//解析淘口令接口
func ParseTpwd(code string) (pwd *TaoPwd, err error) {

	if code == "" {
		return nil, errors.New("淘口令不能为空！")
	}
	//替换为标准的淘口令
	//TODO
	codeSlice := code[1 : len(code)-1]
	var codeBuf strings.Builder
	codeBuf.WriteString("￥")
	codeBuf.WriteString(codeSlice)
	codeBuf.WriteString("￥")

	reqUrl := "https://api.m.taobao.com/h5/com.taobao.redbull.getpassworddetail/1.0/?"

	params := make(map[string]string)

	params["appKey"] = "12574478"
	params["t"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	params["api"] = "com.taobao.redbull.getpassworddetail"
	params["v"] = "1.0"

	client := http.DefaultClient

	req, _ := http.NewRequest("GET", reqUrl+queryBuilder(params), nil)
	req.Header.Add("User-Agent", "Mozilla/8.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")

	resp, err := client.Do(req)

	cookieMap := make(map[string]*http.Cookie)

	for _, v := range resp.Cookies() {
		cookieMap[v.Name] = v
	}

	//获取_m_h5_tk
	m_h5_tk := cookieMap["_m_h5_tk"].Value
	m_h5_tk_slice := strings.Split(m_h5_tk, "_")
	fmt.Println(m_h5_tk[0])
	//获取_m_h5_tk_enc
	m_h5_tk_enc := cookieMap["_m_h5_tk_enc"].Value
	//获取cookies过期时间
	fmt.Println("cookie 过期时间:")
	fmt.Println(cookieMap["t"].RawExpires)
	//sign
	var signBuf strings.Builder
	signBuf.WriteString(m_h5_tk_slice[0])
	signBuf.WriteByte('&')
	signBuf.WriteString(params["t"])
	signBuf.WriteByte('&')
	signBuf.WriteString("12574478")
	signBuf.WriteByte('&')

	passwordData := PasswordData{Password: codeBuf.String()}
	data, _ := json.Marshal(passwordData)

	signBuf.WriteString(string(data))
	params["sign"] = fmt.Sprintf("%x", md5.Sum([]byte(signBuf.String())))

	//重新发起请求
	params["data"] = string(data)

	req2, _ := http.NewRequest("GET", reqUrl+queryBuilder(params), nil)
	var cookie strings.Builder

	cs1 := &http.Cookie{Name: "_m_h5_tk", Value: m_h5_tk}
	cs2 := &http.Cookie{Name: "_m_h5_tk_enc", Value: m_h5_tk_enc}

	cookie.WriteString(cs1.String())
	cookie.WriteByte(';')
	cookie.WriteByte(' ')
	cookie.WriteString(cs2.String())
	cookie.WriteByte(';')

	req2.Header.Add("Cookie", cookie.String())

	fmt.Println(reqUrl + queryBuilder(params))
	resp2, err := client.Do(req2)

	defer resp2.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
	}
	bodyBytes, _ := ioutil.ReadAll(resp2.Body)

	pwd = &TaoPwd{}
	err = json.Unmarshal(bodyBytes, pwd)
	if err != nil {
		return nil, err
	}

	return pwd, nil
}

func queryBuilder(params map[string]string) string {
	var query strings.Builder
	for k, v := range params {
		query.WriteString(k)
		query.WriteByte('=')
		query.WriteString(v)
		query.WriteByte('&')
	}
	//return query.String()[:len(query.String())-1]
	return strings.TrimRight(query.String(), "&")
}

type PasswordData struct {
	Password string `json:"password"`
}
