package grequestsx

import (
	"errors"
	"github.com/json-iterator/go"
	. "github.com/levigross/grequests"
)

func UseSessionOrCreate(session *Session, options *RequestOptions) *Session {
	if session != nil {
		return session
	}
	return NewSession(options)
}

func ResponseToJson(res *Response) (jsoniter.Any, error) {
	isjson := jsoniter.Valid(res.Bytes())
	if isjson == false {
		return nil, errors.New("parse json error")
	}
	v := jsoniter.Get(res.Bytes())
	return v, nil
}
