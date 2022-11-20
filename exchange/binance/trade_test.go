package binance

import (
	"testing"

	"github.com/dqner/fym/exchange"
)

func TestTradeSubscribe(t *testing.T) {
	return
	client, _ := New("binance", "s", "s", "s")
	//t.Fatalf("client start %v", "a")

	client.SubscribeTrade("btcusdt", "ss", func(trade exchange.TradeDetail) {
		t.Logf("tradeis is %v", trade.Id)
	})
	t.Logf("trade end is %v", "bb")
}
