package decorate

type Sugar struct {
    Topping
}

func NewSugar(drink Drink) Milk {
    l := Milk{}
    l.price = 3
    l.description = "sugar"
    l.drink = drink
    return l
}