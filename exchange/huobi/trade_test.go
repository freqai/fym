package huobi

import (
	"log"
	"testing"
	"time"

	"github.com/dqner/fym/exchange"
)

func TestClient_SubscribeTrade(t *testing.T) {
	handler := func(trade exchange.TradeDetail) {

		t.Logf("[%d] %v", 1, trade)

	}
	symbol := "btcusdt"
	clientId := "tradetest"
	c.SubscribeTrade(symbol, clientId, handler)
	log.Println("subscribed")
	defer func() {
		c.UnsubscribeTrade(symbol, clientId)
		log.Println("unsubscribed")
	}()

	time.Sleep(time.Minute * 1)
}
