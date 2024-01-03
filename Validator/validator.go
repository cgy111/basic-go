package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type Order struct {
	OrderNo string  `validate:"required,min=6,max=20"`
	Amount  float64 `validate:"required,gt=0"`
	//CreatedAt    time.Time `validate:"required,time"`
	//Status       string    `validate:"required oneof=待支付 已支付 已发货"`
	ShippingDate time.Time `validate:"required"`
}

func main() {
	validate := validator.New()

	order := Order{
		OrderNo: "",
		Amount:  100.88,
		//CreatedAt:    time.Now(),
		//Status:       "待支付",
		ShippingDate: time.Now(),
	}

	err := validate.Struct(order)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			//fmt.Println(err.Field(), err.Tag())
			fmt.Println(err)
		}
	} else {
		fmt.Println("订单合法")
	}
}
