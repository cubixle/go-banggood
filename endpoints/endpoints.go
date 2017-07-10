package endpoints

import (
	"bytes"
	"io"
)

// Request defines the interface for all bang good endpoints.
type Request interface {
	GetBody() io.Reader
	GetURL() string
	GetType() string
	SetAccessToken(token string)
}

type Response interface {
	GetBody() io.Reader
	GetError() error

	SetBody(body []byte)
	SetError(err error)
}

type ResponseHandler struct {
	err  error
	body []byte
}

func (r ResponseHandler) GetBody() io.Reader {
	return bytes.NewReader(r.body)
}

func (r ResponseHandler) GetError() error {
	return r.err
}

func (r ResponseHandler) SetBody(body []byte) {
	r.body = body
}

func (r ResponseHandler) SetError(err error) {
	r.err = err
}
