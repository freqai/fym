package gateio

import (
	"github.com/dqner/fym/exchange"
	"testing"
)

// test the interface requirement
func TestInterface(t *testing.T) {
	var rest exchange.RestAPIExchange
	rest = New("key", "secret", "host", nil)
	_ = rest
	var ws exchange.WsAPIExchange
	ws = New("key", "secret", "host", nil)
	_ = ws
}
