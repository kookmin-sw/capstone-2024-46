package api

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/mileusna/useragent"
)

type DeviceType int

const (
	UnknownDeviceUser DeviceType = iota
	MobileUser
	TabletUser
	DesktopUser
	BotUser
)

type ClientInfo struct {
	ReqURL     url.URL
	OriginURL  *url.URL
	IP         string
	DeviceType DeviceType
}

func (c *ClientInfo) Proto() string {
	return c.ReqURL.Scheme
}

func (c *ClientInfo) HostWithoutPort() string {
	return c.ReqURL.Hostname()
}

func GetClientInfo(eCtx echo.Context) *ClientInfo {
	req := eCtx.Request()
	reqURL := *req.URL
	reqURL.Host = req.Host

	scheme := req.Header.Get("X-Forwarded-Proto")
	if scheme == "" {
		// 헤더가 없다면, TLS 필드를 확인합니다.
		if req.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}
	if scheme != "" {
		reqURL.Scheme = scheme
	}

	deviceType := UnknownDeviceUser
	userAgent := req.Header.Get("User-Agent")
	if userAgent != "" {
		ua := useragent.Parse(userAgent)
		switch {
		case ua.Mobile:
			deviceType = MobileUser
		case ua.Tablet:
			deviceType = TabletUser
		case ua.Desktop:
			deviceType = DesktopUser
		case ua.Bot:
			deviceType = BotUser
		}
	}

	var originURL *url.URL
	origin := req.Header.Get("Origin")
	if origin != "" {
		originURL, _ = url.Parse(origin)
	}

	return &ClientInfo{
		ReqURL:     reqURL,
		OriginURL:  originURL,
		IP:         eCtx.RealIP(),
		DeviceType: deviceType,
	}
}
