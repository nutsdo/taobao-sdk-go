package auth

type Credentials struct {
	AppKey    string
	AppSecret []byte
}

func New(appKey, appSecret string) *Credentials {
	return &Credentials{appKey, []byte(appSecret)}
}
