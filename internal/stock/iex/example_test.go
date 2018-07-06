package iex_test

import (
	"context"
	"fmt"
	"log"

	"github.com/btmura/ponzi2/internal/stock/iex"
)

func ExampleClient_GetTradingSessionSeries() {
	ctx := context.Background()

	c := iex.NewClient(false)

	req := &iex.GetTradingSessionSeriesRequest{Symbol: "MSFT"}

	resp, err := c.GetTradingSessionSeries(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	for i, t := range resp.TradingSessions {
		fmt.Printf("%d: %s O: %.2f H: %.2f L: %.2f C: %.2f V: %d\n",
			i, t.Date, t.Open, t.High, t.Low, t.Close, t.Volume)
	}
}