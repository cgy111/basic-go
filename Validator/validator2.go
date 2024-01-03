package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type Order2 struct {
	OrderNo      string    `validate:"required,min=6,max=20"`
	Amount       float64   `validate:"required,gt=0"`
	CreatedAt    time.Time `validate:"required"`
	Status       string    `validate:"required,oneof=待支付 已支付 已发货" utf8string:"true"`
	ShippingDate time.Time `validate:"required"`
}

func main() {
	validate := validator.New()

	order := Order2{
		OrderNo:      "123456",
		Amount:       100.88,
		CreatedAt:    parseTime("2022-01-01 02:00:00"),
		Status:       "待支付",
		ShippingDate: time.Now(),
	}

	err := validate.Struct(order)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("订单合法")
	}
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	return t
}
