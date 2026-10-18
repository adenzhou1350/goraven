package util

import (
	"net"
	"net/netip"
	"net/url"
	"strings"
)

// EncodeQuery URL 编码查询参数
func EncodeQuery(s string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(s, " ", "%20"),
			"/", "%2F",
		),
		"?", "%3F",
	)
}

// EncodePath URL 编码路径段
func EncodePath(s string) string {
	return strings.ReplaceAll(s, "/", "%2F")
}

// IsLocalOrIPHost 判断 Host（含可选端口）是否为 IP 地址或 localhost，
// 这类地址不适合作为社交卡片上的对外 URL。
func IsLocalOrIPHost(host string) bool {
	h := strings.TrimSpace(host)
	if hostname, _, err := net.SplitHostPort(h); err == nil {
		h = hostname
	} else if strings.HasPrefix(h, "[") && strings.HasSuffix(h, "]") {
		h = h[1 : len(h)-1]
	}
	if h == "localhost" || strings.HasSuffix(h, ".localhost") {
		return true
	}
	_, err := netip.ParseAddr(h)
	return err == nil
}

// IsLocalOrIPURL 判断形如 "scheme://host[:port]" 的配置域名是否为本地/IP 地址。
// 解析失败或不含 host 时返回 false。
func IsLocalOrIPURL(u string) bool {
	parsed, err := url.Parse(u)
	if err != nil {
		return false
	}
	if parsed.Host == "" {
		return false
	}
	return IsLocalOrIPHost(parsed.Host)
}
