package tasks

type UserAgentAndCookies struct {
	UserAgent *string `json:"userAgent,omitempty"`
	Cookies   *string `json:"cookies,omitempty"`
}

func (t UserAgentAndCookies) WithCookies(cookies string) UserAgentAndCookies {
	t.Cookies = &cookies
	return t
}

func (t UserAgentAndCookies) WithUserAgent(userAgent string) UserAgentAndCookies {
	t.UserAgent = &userAgent
	return t
}
