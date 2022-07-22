package grequestsx

import (
	. "github.com/levigross/grequests"
	"github.com/valyala/fastjson"
)

func UseSessionOrCreate(session *Session, options *RequestOptions) *Session {
	if session != nil {
		return session
	}
	return NewSession(options)
}

func ResponseToJson(res *Response) (*fastjson.Value, error) {
	v, err := fastjson.Parse(res.String())
	if err != nil {
		return nil, err
	}
	return v, nil
}
