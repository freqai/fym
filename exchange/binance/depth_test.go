package binance

import (
	"testing"

	"github.com/dqner/fym/exchange"
)

func TestDepthSubscribe(t *testing.T) {
	client, _ := New("binance", "s", "s", "s")
	//t.Fatalf("client start %v", "a")

	client.SubscribeDepth("btcusdt", "ss", func(trade exchange.OrderBook) {
		t.Logf("tradeis is %v", trade.Id)
	})
	t.Logf("trade end is %v", "bb")
}
