package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// NewGetStockRequest
func NewGetStockRequest(productID string) *GetStockRequest {
	return &GetStockRequest{ProductID: productID}
}

// GetStockRequest
type GetStockRequest struct {
	AccessToken string
	ProductID   string
	Lang        string
}

func (r *GetStockRequest) GetBody() io.Reader {
	return bytes.NewReader([]byte{})
}

func (r *GetStockRequest) GetURL() string {
	return fmt.Sprintf(
		"product/getStocks?access_token=%s&product_id=%s&lang=%s",
		r.AccessToken,
		r.ProductID,
		r.Lang,
	)
}

func (r *GetStockRequest) GetType() string {
	return "GET"
}

func (r *GetStockRequest) SetAccessToken(token string) {
	r.AccessToken = token
}

type GetStockResponse struct {
	Stocks []Stocks `json:"stocks"`
	Code   int      `json:"code"`
	Lang   string   `json:"lang"`
}

type Stocks struct {
	Warehouse string      `json:"warehouse"`
	StockList []StockList `json:"stock_list"`
}

type StockList struct {
	PoaID     int    `json:"poa_id"`
	Poa       string `json:"pod"`
	Stock     string `json:"stock"`
	StocksMsg string `json:"stocks_msg"`
}

func (r GetStockResponse) ParseResponse(rsp io.Reader) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rsp)
	json.Unmarshal(buf.Bytes(), &r)
}
