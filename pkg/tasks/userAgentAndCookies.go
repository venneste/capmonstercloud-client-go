package tasks

type UserAgentAndCookies struct {
	UserAgent string `json:"userAgent,omitempty"`
	Cookies   string `json:"cookies,omitempty"`
}

func (t *UserAgentAndCookies) WithCookies(cookies string) {
	t.Cookies = cookies
}

func (t *UserAgentAndCookies) WithUserAgent(userAgent string) {
	t.UserAgent = userAgent
}
