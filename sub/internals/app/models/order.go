package models

import "time"

//модель данных входящего json

type Order struct {
	DateCreated       time.Time `json:"date_created"`
	Delivery          `json:"delivery"`
	Locale            string `json:"locale"`
	Entry             string `json:"entry"`
	ID                string `json:"order_uid"`
	InternalSignature string `json:"internal_signature"`
	CustomerID        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	Shardkey          string `json:"shardkey"`
	TrackNumber       string `json:"track_number"`
	OofShard          string `json:"oof_shard"`
	Items             []Item `json:"items"`
	Payment           `json:"payment"`
	SmID              int `json:"sm_id"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string  `json:"transaction"`
	RequestID    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Bank         string  `json:"bank"`
	Amount       float64 `json:"amount"`
	PaymentDt    int     `json:"payment_dt"`
	DeliveryCost float64 `json:"delivery_cost"`
	GoodsTotal   float64 `json:"goods_total"`
	CustomFee    float64 `json:"custom_fee"`
}

type Item struct {
	TrackNumber string  `json:"track_number"`
	Rid         string  `json:"rid"`
	Name        string  `json:"name"`
	Size        string  `json:"size"`
	Brand       string  `json:"brand"`
	ChrtID      int     `json:"chrt_id"`
	Price       float64 `json:"price"`
	Sale        int     `json:"sale"`
	TotalPrice  float64 `json:"total_price"`
	NmID        int     `json:"nm_id"`
	Status      int     `json:"status"`
}
