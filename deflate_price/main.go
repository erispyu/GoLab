package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

const DbPriceToActualPriceRatio int64 = 100000

// DeflatePrice has no decimalPlaces parameter as python, you can use DeflatePriceWithDecimalPlaces if needed
func DeflatePrice(price int64) decimal.Decimal {
	return decimal.NewFromInt(price).Div(decimal.NewFromInt(DbPriceToActualPriceRatio))
}

func main() {
	var amount int64 = -3500000
	fmt.Println(DeflatePrice(amount).IntPart())
	fmt.Println(amount / DbPriceToActualPriceRatio)
}
