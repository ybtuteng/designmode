package main

import (
    "awesomeProject/designmodel/decorate"
    "log"
)

func main() {
    var order decorate.Drink
    order = decorate.NewLatte()
    order = decorate.NewMilk(order)
    order = decorate.NewSugar(order)
    log.Println(order.Cost())
}
