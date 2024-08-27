package tasks

import (
	"math"
	"net/url"
)

type FunCaptchaTaskProxyless struct {
	Type                     string  `json:"type"`
	WebsiteURL               string  `json:"websiteURL"`
	FuncaptchaApiJSSubdomain *string `json:"funcaptchaApiJSSubdomain"`
	WebsitePublicKey         string  `json:"websitePublicKey"`
	Data                     *string `json:"data"`
}

func NewFunCaptchaTaskProxyless(websiteURL, websitePublicKey string) FunCaptchaTaskProxyless {
	return FunCaptchaTaskProxyless{
		Type:             "FunCaptchaTaskProxyless",
		WebsiteURL:       websiteURL,
		WebsitePublicKey: websitePublicKey,
	}
}

func (t FunCaptchaTaskProxyless) WithFuncaptchaApiJSSubdomain(funcaptchaApiJSSubdomain string) FunCaptchaTaskProxyless {
	t.FuncaptchaApiJSSubdomain = &funcaptchaApiJSSubdomain
	return t
}

func (t FunCaptchaTaskProxyless) WithData(data string) FunCaptchaTaskProxyless {
	t.Data = &data
	return t
}

func (t FunCaptchaTaskProxyless) Validate() error {
	if _, err := url.ParseRequestURI(t.WebsiteURL); err != nil {
		return ErrInvalidWebsiteUrl
	}
	if len(t.WebsitePublicKey) < 1 || len(t.WebsitePublicKey) > math.MaxInt {
		return ErrInvalidWebSiteKey
	}
	return nil
}

type FunCaptchaTask struct {
	FunCaptchaTaskProxyless
	taskProxy
	UserAgentAndCookies
}

func NewFunCaptchaTask(websiteURL, websitePublicKey, proxyType, proxyAddress, userAgent string, proxyPort int) FunCaptchaTask {
	return FunCaptchaTask{
		FunCaptchaTaskProxyless: FunCaptchaTaskProxyless{
			Type:             "FunCaptchaTask",
			WebsiteURL:       websiteURL,
			WebsitePublicKey: websitePublicKey,
		},
		taskProxy: taskProxy{
			ProxyType:    proxyType,
			ProxyAddress: proxyAddress,
			ProxyPort:    proxyPort,
		},
		UserAgentAndCookies: UserAgentAndCookies{
			UserAgent: &userAgent,
		},
	}
}

func (t FunCaptchaTask) Validate() error {
	if err := t.FunCaptchaTaskProxyless.Validate(); err != nil {
		return err
	}
	if err := t.taskProxy.validate(); err != nil {
		return err
	}
	return nil
}

type FunCaptchaTaskSolution struct {
	Token   string            `json:"Token"`
	Cookies map[string]string `json:"cookies"`
}
