package stocks

import (
	"testing"
)

func TestGetStock(t *testing.T) {
	stock, _ := GetQuote("aapl")
	if stock.GetSymbol() != "AAPL" {
		t.Error("The stock symbol should be equal to aapl but equals to", stock.GetSymbol())
	}
}
