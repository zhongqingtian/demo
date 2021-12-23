package main

import (
	"context"
	"google.golang.org/grpc/metadata"
	"strconv"
	"strings"
)

type Header struct {
	//request_uri str
	//version         str
	//server          str
	//method          str
	//opaque          str
	//username   str
	//password   str
	//password_set bool
	//host            str
	//path            str
	//raw_path         str
	//force_query      str
	//raw_query        str
	//fragment        str
	//remote_addr      str
	//is_form_http bool
	ctx context.Context

	headers map[string]string
	//trailers map[str]str
	cookies map[string]string
	md      metadata.MD
}

//获取登录的用户ID
func GetLoginUserId(ctx context.Context) uint64 {
	headerData := NewHeader(ctx)
	userId := headerData.GetCookie("userid")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		userIdInt = 0
	}
	return uint64(userIdInt)
}

func NewHeader(ctx context.Context) *Header {
	header := &Header{
		ctx:     ctx,
		headers: make(map[string]string),
		cookies: make(map[string]string),
		md:      make(metadata.MD),
	}
	header.parse()
	return header
}

func (h *Header) parse() {
	md, ok := metadata.FromIncomingContext(h.ctx)
	if ok {
		for key, value := range md {
			if len(value) > 0 {
				h.headers[key] = value[0]
			} else {
				h.headers[key] = ""
			}
		}
	}

	cookies, ok := md["cookie"]
	if ok {
		if len(cookies) > 0 {
			c := strings.Split(cookies[0], ";")
			for _, iv := range c {
				iv = strings.Trim(iv, " ")
				it := strings.Split(iv, "=")
				if len(it) >= 2 {
					k := strings.ToLower(strings.Trim(it[0], " "))
					h.cookies[k] = strings.Trim(it[1], " ")
				}
			}
		}
	}

	md, ok = metadata.FromOutgoingContext(h.ctx)
	if ok {
		for key, value := range md {
			if len(value) > 0 {
				h.headers[key] = value[0]
			} else {
				h.headers[key] = ""
			}
		}
	}

}

// 如果返回空字符串，说明key不存在
func (h *Header) GetHeader(key string) string {
	v, ok := h.headers[strings.ToLower(key)]
	if ok {
		return v
	}
	return ""
}

// 所有的key都是小写
func (h *Header) GetHeaders() map[string]string {
	return h.headers
}

// 如果返回空字符串，说明cookie不存在
func (h *Header) GetCookie(key string) string {
	v, ok := h.cookies[strings.ToLower(key)]
	if ok {
		return v
	}
	return ""
}

// 所有的key都是小写
func (h *Header) GetCookies() map[string]string {
	return h.cookies
}

func (h *Header) GetCookieInt64(key string) int64 {
	v, ok := h.cookies[strings.ToLower(key)]
	if ok {
		d, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0
		}
		return d
	}
	return 0
}

func (h *Header) GetHeaderInt64(key string) int64 {
	v, ok := h.headers[strings.ToLower(key)]
	if ok {
		d, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0
		}
		return d
	}
	return 0
}
