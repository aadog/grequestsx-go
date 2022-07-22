package grequestsx

import (
	. "github.com/levigross/grequests"
)

func UseSessionOrCreate(session *Session, options *RequestOptions) *Session {
	if session != nil {
		return session
	}
	return NewSession(options)
}
