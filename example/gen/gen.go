// Code generated by gowsdl DO NOT EDIT.

package gen

import (
	"context"
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type TradePriceRequest struct {
	XMLName xml.Name `xml:"http://example.com/stockquote.xsd TradePriceRequest"`

	TickerSymbol string `xml:"tickerSymbol,omitempty" json:"tickerSymbol,omitempty"`
}

type TradePrice struct {
	XMLName xml.Name `xml:"http://example.com/stockquote.xsd TradePrice"`

	Price float32 `xml:"price,omitempty" json:"price,omitempty"`
}

type StockQuotePortType interface {
	GetLastTradePrice(request *TradePriceRequest) (*TradePrice, error)

	GetLastTradePriceContext(ctx context.Context, request *TradePriceRequest) (*TradePrice, error)
}

type SOAPCaller interface {
	CallContext(ctx context.Context, soapAction string, request, response interface{}) error
}

type stockQuotePortType struct {
	client SOAPCaller
}

func NewStockQuotePortType(client SOAPCaller) StockQuotePortType {
	return &stockQuotePortType{
		client: client,
	}
}

func (service *stockQuotePortType) GetLastTradePriceContext(ctx context.Context, request *TradePriceRequest) (*TradePrice, error) {
	response := new(TradePrice)
	err := service.client.CallContext(ctx, "http://example.com/GetLastTradePrice", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *stockQuotePortType) GetLastTradePrice(request *TradePriceRequest) (*TradePrice, error) {
	return service.GetLastTradePriceContext(
		context.Background(),
		request,
	)
}
