package decorate

type Milk struct {
    Topping
}

func NewMilk(drink Drink) Milk {
    l := Milk{}
    l.price = 5
    l.description = "milk"
    l.drink = drink
    return l
}
