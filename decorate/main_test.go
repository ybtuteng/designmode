package decorate

import (
    "log"
    "testing"
)

//golang实现的一个装饰者模式，大概思想就是用配料去包装咖啡类，递归包装，计算价格的时候也递归计算
func TestDecorate(t *testing.T) {
    var order Drink
    order = NewLatte()
    order = NewMilk(order)
    order = NewSugar(order)
    log.Println(order.Cost())
}
