package pop

import (
	"context"
	"net/http"
)

type TxContext interface {
	context.Context
	//Response() http.ResponseWriter
	Request() *http.Request
	// Session() *buffalo.Session
	// Cookies() *buffalo.Cookies
	// Params() ParamValues
	Param(string) string
	// Set(string, interface{})
	// LogField(string, interface{})
	// LogFields(map[string]interface{})
	// Logger() Logger
	// Bind(interface{}) error
	// Render(int, render.Renderer) error
	Error(int, error) error
	// Websocket() (*websocket.Conn, error)
	// Redirect(int, string, ...interface{}) error
	Data() map[string]interface{}
	// Flash() *Flash
	// File(string) (binding.File, error)
}
