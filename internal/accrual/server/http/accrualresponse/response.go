package accrualresponse

import (
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	AppCode        int64  `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

func (e *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequestFormat(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "invalid request format.",
		ErrorText:      err.Error(),
	}
}

func ErrInternalServer(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal server error",
		ErrorText:      err.Error(),
	}
}

var RewardRegistered = &Response{HTTPStatusCode: 200, StatusText: "reward successfully registered."}
var OrderSuccessfully = &Response{HTTPStatusCode: 202, StatusText: "the order has been successfully processed."}
