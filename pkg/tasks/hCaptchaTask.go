package tasks

import (
	"math"
	"net/url"
)

type HCaptchaTaskProxyless struct {
	Type        string  `json:"type"`
	WebsiteURL  string  `json:"websiteURL"`
	WebsiteKey  string  `json:"websiteKey"`
	IsInvisible *bool   `json:"isInvisible,omitempty"`
	Data        *string `json:"data,omitempty"`
	UserAgentAndCookies
}

func NewHCaptchaTaskProxyless(websiteURL, websiteKey string) HCaptchaTaskProxyless {
	return HCaptchaTaskProxyless{
		Type:       "HCaptchaTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
	}
}

func (t HCaptchaTaskProxyless) WithIsInvisible(isInvisible bool) HCaptchaTaskProxyless {
	t.IsInvisible = &isInvisible
	return t
}

func (t HCaptchaTaskProxyless) WithData(data string) HCaptchaTaskProxyless {
	t.Data = &data
	return t
}

func (t HCaptchaTaskProxyless) Validate() error {
	if _, err := url.ParseRequestURI(t.WebsiteURL); err != nil {
		return ErrInvalidWebsiteUrl
	}

	if len(t.WebsiteKey) < 1 || len(t.WebsiteKey) > math.MaxInt {
		return ErrInvalidWebSiteKey
	}

	if t.Data != nil && t.UserAgent == nil {
		return ErrUserAgentRequired
	}
	return nil
}

type HCaptchaTask struct {
	HCaptchaTaskProxyless
	taskProxy
}

func NewHCaptchaTask(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int) HCaptchaTask {
	return HCaptchaTask{
		HCaptchaTaskProxyless: HCaptchaTaskProxyless{
			Type:       "HCaptchaTask",
			WebsiteURL: websiteURL,
			WebsiteKey: websiteKey,
		},
		taskProxy: taskProxy{
			ProxyType:    proxyType,
			ProxyAddress: proxyAddress,
			ProxyPort:    proxyPort,
		},
	}
}

func (t HCaptchaTask) Validate() error {
	if err := t.HCaptchaTaskProxyless.Validate(); err != nil {
		return err
	}
	if err := t.taskProxy.validate(); err != nil {
		return err
	}
	return nil
}

type HCaptchaTaskSolution struct {
	GRecaptchaResponse string            `json:"gRecaptchaResponse"`
	RespKey            string            `json:"respKey"`
	UserAgent          string            `json:"userAgent"`
	Cookies            map[string]string `json:"cookies"`
}
